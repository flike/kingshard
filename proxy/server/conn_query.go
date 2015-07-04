package server

import (
	"fmt"
	"github.com/flike/kingshard/backend"
	. "github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/core/hack"
	. "github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
	"strconv"
	"strings"
	"sync"
)

/*处理query语句*/
func (c *ClientConn) handleQuery(sql string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("execute %s error %v", sql, e)
			return
		}
	}()

	sql = strings.TrimRight(sql, ";") //删除sql语句最后的分号

	hasHandled, err := c.handleUnsupport(sql)
	if err != nil {
		golog.Error("server", "parse", err.Error(), 0, "hasHandled", hasHandled)
		return err
	}
	if hasHandled {
		return nil
	}

	var stmt sqlparser.Statement
	stmt, err = sqlparser.Parse(sql) //解析sql语句,得到的stmt是一个interface
	if err != nil {
		golog.Error("server", "parse", err.Error(), 0, "hasHandled", hasHandled)
		return err
	}

	switch v := stmt.(type) {
	case *sqlparser.Select:
		return c.handleSelect(v, sql, nil)
	case *sqlparser.Insert:
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Update:
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Delete:
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Replace:
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Set:
		return c.handleSet(v)
	case *sqlparser.Begin:
		return c.handleBegin()
	case *sqlparser.Commit:
		return c.handleCommit()
	case *sqlparser.Rollback:
		return c.handleRollback()
	case *sqlparser.SimpleSelect:
		return c.handleSimpleSelect(sql, v)
	case *sqlparser.Show:
		return c.handleShow(sql, v)
	case *sqlparser.Admin:
		return c.handleAdmin(v)
	default:
		return fmt.Errorf("statement %T not support now", stmt)
	}

	return nil
}

/*获取分片的节点列表*/
func (c *ClientConn) getShardList(stmt sqlparser.Statement, bindVars map[string]interface{}) ([]*backend.Node, error) {
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	ns, err := c.schema.rule.GetStmtShardList(stmt, bindVars)
	if err != nil {
		return nil, err
	}

	if len(ns) == 0 {
		return nil, nil
	}

	n := make([]*backend.Node, 0, len(ns))
	for _, name := range ns {
		n = append(n, c.proxy.GetNode(name))
	}
	return n, nil
}

func (c *ClientConn) getBackendConn(n *backend.Node, isSelect bool) (co *backend.BackendConn, err error) {
	if !c.needBeginTx() {
		if isSelect {
			co, err = n.GetSlaveConn()
			if err != nil && err == ErrNoSlaveConn {
				co, err = n.GetMasterConn()
			}
		} else {
			co, err = n.GetMasterConn()
		}
		if err != nil {
			golog.Error("server", "getBackendConn", err.Error(), 0)
			return
		}
	} else {
		var ok bool
		c.Lock()
		co, ok = c.txConns[n]
		c.Unlock()

		if !ok {
			if co, err = n.GetMasterConn(); err != nil {
				return
			}

			if err = co.Begin(); err != nil {
				return
			}

			c.Lock()
			c.txConns[n] = co
			c.Unlock()
		}
	}

	//todo, set conn charset, etc...
	if err = co.UseDB(c.schema.db); err != nil {
		return
	}

	if err = co.SetCharset(c.charset); err != nil {
		return
	}

	return
}

/*获取shard的conn，第一个参数表示是不是select*/
func (c *ClientConn) getShardConns(isSelect bool, stmt sqlparser.Statement, bindVars map[string]interface{}) ([]*backend.BackendConn, error) {
	nodes, err := c.getShardList(stmt, bindVars)
	if err != nil {
		return nil, err
	} else if nodes == nil {
		return nil, nil
	}

	conns := make([]*backend.BackendConn, 0, len(nodes))

	var co *backend.BackendConn
	for _, n := range nodes {
		co, err = c.getBackendConn(n, isSelect)
		if err != nil {
			break
		}

		conns = append(conns, co)
	}

	return conns, err
}

func (c *ClientConn) executeInShard(conns []*backend.BackendConn, sql string, args []interface{}) ([]*Result, error) {
	var wg sync.WaitGroup
	wg.Add(len(conns))

	rs := make([]interface{}, len(conns))

	f := func(rs []interface{}, i int, co *backend.BackendConn) {
		r, err := co.Execute(sql, args...)
		if err != nil {
			rs[i] = err
		} else {
			rs[i] = r
		}

		wg.Done()
	}

	for i, co := range conns {
		go f(rs, i, co)
	}

	wg.Wait()

	var err error
	r := make([]*Result, len(conns))
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*Result)
	}

	return r, err
}

func (c *ClientConn) closeShardConns(conns []*backend.BackendConn, rollback bool) {
	if c.isInTransaction() {
		return
	}

	for _, co := range conns {
		if rollback {
			co.Rollback()
		}

		co.Close()
	}
}

func (c *ClientConn) newEmptyResultset(stmt *sqlparser.Select) *Resultset {
	r := new(Resultset)
	r.Fields = make([]*Field, len(stmt.SelectExprs))

	for i, expr := range stmt.SelectExprs {
		r.Fields[i] = &Field{}
		switch e := expr.(type) {
		case *sqlparser.StarExpr:
			r.Fields[i].Name = []byte("*")
		case *sqlparser.NonStarExpr:
			if e.As != nil {
				r.Fields[i].Name = e.As
				r.Fields[i].OrgName = hack.Slice(nstring(e.Expr))
			} else {
				r.Fields[i].Name = hack.Slice(nstring(e.Expr))
			}
		default:
			r.Fields[i].Name = hack.Slice(nstring(e))
		}
	}

	r.Values = make([][]interface{}, 0)
	r.RowDatas = make([]RowData, 0)

	return r
}

func makeBindVars(args []interface{}) map[string]interface{} {
	bindVars := make(map[string]interface{}, len(args))

	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}

	return bindVars
}

//返回true表示已经处理，false表示未处理
func (c *ClientConn) handleUnsupport(sql string) (bool, error) {
	var rs []*Result
	var TK_FROM string = "from"

	sql = strings.ToLower(sql)
	tokens := strings.Fields(sql)
	tokensLen := len(tokens)
	if 0 < tokensLen {
		//token is in WHITE_TOKEN_MAP
		if 0 < WHITE_TOKEN_MAP[tokens[0]] {
			//select
			if 1 < WHITE_TOKEN_MAP[tokens[0]] {
				for i := 1; i < tokensLen; i++ {
					if tokens[i] == TK_FROM {
						return false, nil
					}
				}
			} else {
				return false, nil
			}
		}
	}

	defaultRule := c.schema.rule.DefaultRule
	if len(defaultRule.Nodes) == 0 {

		return false, ErrNoDefaultNode
	}
	defaultNode := c.proxy.GetNode(defaultRule.Nodes[0])

	//execute in Master DB
	conn, err := c.getBackendConn(defaultNode, false)
	if err != nil {

		return false, err
	}

	conns := []*backend.BackendConn{conn}
	rs, err = c.executeInShard(conns, sql, nil)
	if err != nil {

		return false, err
	}

	c.closeShardConns(conns, false)
	if len(rs) == 0 {
		msg := fmt.Sprintf("result is empty")
		golog.Error("ClientConn", "handleUnsupport", msg, c.connectionId)
		return false, NewError(ER_UNKNOWN_ERROR, msg)
	}

	err = c.writeResultset(c.status, rs[0].Resultset)
	if err != nil {

		return false, err
	}

	return true, nil
}

/*处理select语句*/
func (c *ClientConn) handleSelect(stmt *sqlparser.Select, sql string, args []interface{}) error {
	bindVars := makeBindVars(args) //对于select语句，arg为空，不考虑

	conns, err := c.getShardConns(true, stmt, bindVars)
	if err != nil {
		golog.Error("ClientConn", "handleSelect", err.Error(), c.connectionId)
		return err
	} else if conns == nil {
		r := c.newEmptyResultset(stmt)
		return c.writeResultset(c.status, r)
	}

	var rs []*Result

	rs, err = c.executeInShard(conns, sql, args)

	c.closeShardConns(conns, false)

	if err != nil {
		golog.Error("ClientConn", "handleSelect", err.Error(), c.connectionId)
		return err
	}

	err = c.mergeSelectResult(rs, stmt)
	if err != nil {
		golog.Error("ClientConn", "handleSelect", err.Error(), c.connectionId)
	}

	return err
}

func (c *ClientConn) beginShardConns(conns []*backend.BackendConn) error {
	if c.isInTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Begin(); err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientConn) commitShardConns(conns []*backend.BackendConn) error {
	if c.isInTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientConn) handleExec(stmt sqlparser.Statement, sql string, args []interface{}) error {
	bindVars := makeBindVars(args)

	conns, err := c.getShardConns(false, stmt, bindVars)
	if err != nil {
		return err
	} else if conns == nil {
		return c.writeOK(nil)
	}

	var rs []*Result

	if len(conns) == 1 {
		rs, err = c.executeInShard(conns, sql, args)
	} else {
		//for multi nodes, 2PC simple, begin, exec, commit
		//if commit error, data maybe corrupt
		for {
			if err = c.beginShardConns(conns); err != nil {
				break
			}

			if rs, err = c.executeInShard(conns, sql, args); err != nil {
				break
			}

			err = c.commitShardConns(conns)
			break
		}
	}

	c.closeShardConns(conns, err != nil)

	if err == nil {
		err = c.mergeExecResult(rs)
	}

	return err
}

func (c *ClientConn) mergeExecResult(rs []*Result) error {
	r := new(Result)

	for _, v := range rs {
		r.Status |= v.Status
		r.AffectedRows += v.AffectedRows
		if r.InsertId == 0 {
			r.InsertId = v.InsertId
		} else if r.InsertId > v.InsertId {
			//last insert id is first gen id for multi row inserted
			//see http://dev.mysql.com/doc/refman/5.6/en/information-functions.html#function_last-insert-id
			r.InsertId = v.InsertId
		}
	}

	if r.InsertId > 0 {
		c.lastInsertId = int64(r.InsertId)
	}

	c.affectedRows = int64(r.AffectedRows)

	return c.writeOK(r)
}

func (c *ClientConn) mergeSelectResult(rs []*Result, stmt *sqlparser.Select) error {
	r := rs[0].Resultset

	status := c.status | rs[0].Status

	for i := 1; i < len(rs); i++ {
		status |= rs[i].Status

		//check fields equal

		for j := range rs[i].Values {
			r.Values = append(r.Values, rs[i].Values[j])
			r.RowDatas = append(r.RowDatas, rs[i].RowDatas[j])
		}
	}

	//to do order by, group by, limit offset
	c.sortSelectResult(r, stmt)
	//to do, add log here, sort may error because order by key not exist in resultset fields

	if err := c.limitSelectResult(r, stmt); err != nil {
		return err
	}

	return c.writeResultset(status, r)
}

func (c *ClientConn) sortSelectResult(r *Resultset, stmt *sqlparser.Select) error {
	if stmt.OrderBy == nil {
		return nil
	}

	sk := make([]SortKey, len(stmt.OrderBy))

	for i, o := range stmt.OrderBy {
		sk[i].Name = nstring(o.Expr)
		sk[i].Direction = o.Direction
	}

	return r.Sort(sk)
}

func (c *ClientConn) limitSelectResult(r *Resultset, stmt *sqlparser.Select) error {
	if stmt.Limit == nil {
		return nil
	}

	var offset, count int64
	var err error
	if stmt.Limit.Offset == nil {
		offset = 0
	} else {
		if o, ok := stmt.Limit.Offset.(sqlparser.NumVal); !ok {
			return fmt.Errorf("invalid select limit %s", nstring(stmt.Limit))
		} else {
			if offset, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
				return err
			}
		}
	}

	if o, ok := stmt.Limit.Rowcount.(sqlparser.NumVal); !ok {
		return fmt.Errorf("invalid limit %s", nstring(stmt.Limit))
	} else {
		if count, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
			return err
		} else if count < 0 {
			return fmt.Errorf("invalid limit %s", nstring(stmt.Limit))
		}
	}

	if offset+count > int64(len(r.Values)) {
		count = int64(len(r.Values)) - offset
	}

	r.Values = r.Values[offset : offset+count]
	r.RowDatas = r.RowDatas[offset : offset+count]

	return nil
}
