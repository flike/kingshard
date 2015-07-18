package server

import (
	"fmt"
	"github.com/flike/kingshard/backend"
	. "github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/core/hack"
	. "github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/proxy/router"
	"github.com/flike/kingshard/sqlparser"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const (
	MasterComment = "/*master*/"
)

/*处理query语句*/
func (c *ClientConn) handleQuery(sql string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			golog.OutputSql("Error", "%s", sql)

			if err, ok := e.(error); ok {
				const size = 4096
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]

				golog.Error("ClientConn", "handleQuery",
					err.Error(), 0,
					"stack", string(buf), "sql", sql)
			}
			return
		}
		golog.OutputSql("INFO", "%s", sql)
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
		return c.handleSelect(v, nil)
	case *sqlparser.Insert:
		return c.handleExec(stmt, nil)
	case *sqlparser.Update:
		return c.handleExec(stmt, nil)
	case *sqlparser.Delete:
		return c.handleExec(stmt, nil)
	case *sqlparser.Replace:
		return c.handleExec(stmt, nil)
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

func (c *ClientConn) getBackendConn(n *backend.Node, fromSlave bool) (co *backend.BackendConn, err error) {
	if !c.needBeginTx() {
		if fromSlave {
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
func (c *ClientConn) getShardConns(fromSlave bool, plan *router.Plan) (map[string]*backend.BackendConn, error) {
	var err error
	if plan == nil || len(plan.RouteNodeIndexs) == 0 {
		return nil, ErrNoRouteNode
	}

	nodesCount := len(plan.RouteNodeIndexs)
	nodes := make([]*backend.Node, 0, nodesCount)
	for i := 0; i < nodesCount; i++ {
		nodeIndex := plan.RouteNodeIndexs[i]
		nodes = append(nodes, c.proxy.GetNode(plan.Rule.Nodes[nodeIndex]))
	}

	conns := make(map[string]*backend.BackendConn)
	var co *backend.BackendConn
	for _, n := range nodes {
		co, err = c.getBackendConn(n, fromSlave)
		if err != nil {
			break
		}

		conns[n.Cfg.Name] = co
	}

	return conns, err
}

func (c *ClientConn) executeInNode(conn *backend.BackendConn, sql string, args []interface{}) ([]*Result, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	rs := make([]interface{}, 1)

	f := func(rs []interface{}, i int, co *backend.BackendConn) {
		r, err := co.Execute(sql, args...)
		if err != nil {
			rs[i] = err
		} else {
			rs[i] = r
		}

		wg.Done()
	}
	go f(rs, 0, conn)

	wg.Wait()

	var err error
	r := make([]*Result, 1)
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*Result)
	}

	return r, err
}

func (c *ClientConn) executeInMultiNodes(conns map[string]*backend.BackendConn, sqls map[string][]string, args []interface{}) ([]*Result, error) {
	if len(conns) != len(sqls) {
		golog.Error("ClientConn", "executeInMultiNodes", ErrConnNotEqual.Error(), c.connectionId,
			"conns", conns,
			"sqls", sqls,
		)
		return nil, ErrConnNotEqual
	}

	var wg sync.WaitGroup

	if len(conns) == 0 {
		return nil, ErrNoPlan
	}

	wg.Add(len(conns))

	resultCount := 0
	for _, sqlSlice := range sqls {
		resultCount += len(sqlSlice)
	}

	rs := make([]interface{}, resultCount)

	f := func(rs []interface{}, i int, execSqls []string, co *backend.BackendConn) {
		for _, v := range execSqls {
			r, err := co.Execute(v, args...)
			if err != nil {
				rs[i] = err
			} else {
				rs[i] = r
			}
			i++
		}
		wg.Done()
	}

	offsert := 0
	for nodeName, co := range conns {
		s := sqls[nodeName] //[]string
		go f(rs, offsert, s, co)
		offsert += len(s)
	}

	wg.Wait()

	var err error
	r := make([]*Result, resultCount)
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*Result)
	}

	return r, err
}

func (c *ClientConn) closeConn(conn *backend.BackendConn, rollback bool) {
	if c.isInTransaction() {
		return
	}

	if rollback {
		conn.Rollback()
	}

	conn.Close()
}

func (c *ClientConn) closeShardConns(conns map[string]*backend.BackendConn, rollback bool) {
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

	rs, err = c.executeInNode(conn, sql, nil)
	if err != nil {

		return false, err
	}

	c.closeConn(conn, false)
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
func (c *ClientConn) handleSelect(stmt *sqlparser.Select, args []interface{}) error {
	var fromSlave bool = true
	plan, err := c.schema.rule.BuildPlan(stmt)
	if err != nil {
		return err
	}
	if 0 < len(stmt.Comments) {
		comment := string(stmt.Comments[0])
		if 0 < len(comment) && strings.ToLower(comment) == MasterComment {
			fromSlave = false
		}
	}

	conns, err := c.getShardConns(fromSlave, plan)
	if err != nil {
		golog.Error("ClientConn", "handleSelect", err.Error(), c.connectionId)
		return err
	}
	if conns == nil {
		r := c.newEmptyResultset(stmt)
		return c.writeResultset(c.status, r)
	}

	var rs []*Result
	rs, err = c.executeInMultiNodes(conns, plan.RewrittenSqls, args)
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

func (c *ClientConn) handleExec(stmt sqlparser.Statement, args []interface{}) error {
	plan, err := c.schema.rule.BuildPlan(stmt)
	conns, err := c.getShardConns(false, plan)
	if err != nil {
		golog.Error("ClientConn", "handleExec", err.Error(), c.connectionId)
		return err
	}
	if conns == nil {
		return c.writeOK(nil)
	}

	var rs []*Result
	if 1 < len(conns) {
		return ErrExecInMulti
	}
	if 1 < len(plan.RewrittenSqls) {
		nodeIndex := plan.RouteNodeIndexs[0]
		nodeName := plan.Rule.Nodes[nodeIndex]
		txSqls := []string{"begin;"}
		txSqls = append(txSqls, plan.RewrittenSqls[nodeName]...)
		txSqls = append(txSqls, "commit;")
		plan.RewrittenSqls[nodeName] = txSqls
	}

	rs, err = c.executeInMultiNodes(conns, plan.RewrittenSqls, args)
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
