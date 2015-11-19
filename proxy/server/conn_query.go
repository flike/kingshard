// Copyright 2015 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package server

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/flike/kingshard/backend"
	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/core/hack"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/proxy/router"
	"github.com/flike/kingshard/sqlparser"
)

/*处理query语句*/
func (c *ClientConn) handleQuery(sql string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			golog.OutputSql("Error", "err:%v,sql:%s", e, sql)

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
	}()

	sql = strings.TrimRight(sql, ";") //删除sql语句最后的分号
	hasHandled, err := c.preHandleShard(sql)
	if err != nil {
		golog.Error("server", "preHandleShard", err.Error(), 0, "hasHandled", hasHandled)
		return err
	}
	if hasHandled {
		return nil
	}

	var stmt sqlparser.Statement
	stmt, err = sqlparser.Parse(sql) //解析sql语句,得到的stmt是一个interface
	if err != nil {
		golog.Error("server", "parse", err.Error(), 0, "hasHandled", hasHandled, "sql", sql)
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
		return c.handleSet(v, sql)
	case *sqlparser.Begin:
		return c.handleBegin()
	case *sqlparser.Commit:
		return c.handleCommit()
	case *sqlparser.Rollback:
		return c.handleRollback()
	case *sqlparser.Admin:
		return c.handleAdmin(v)
	case *sqlparser.UseDB:
		return c.handleUseDB(v)
	default:
		return fmt.Errorf("statement %T not support now", stmt)
	}

	return nil
}

func (c *ClientConn) getBackendConn(n *backend.Node, fromSlave bool) (co *backend.BackendConn, err error) {
	if !c.isInTransaction() {
		if fromSlave {
			co, err = n.GetSlaveConn()
			if err != nil {
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
	if err = co.UseDB(c.db); err != nil {
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
		return nil, errors.ErrNoRouteNode
	}

	nodesCount := len(plan.RouteNodeIndexs)
	nodes := make([]*backend.Node, 0, nodesCount)
	for i := 0; i < nodesCount; i++ {
		nodeIndex := plan.RouteNodeIndexs[i]
		nodes = append(nodes, c.proxy.GetNode(plan.Rule.Nodes[nodeIndex]))
	}
	if c.isInTransaction() {
		if 1 < len(nodes) {
			return nil, errors.ErrTransInMulti
		}
		//exec in multi node
		if len(c.txConns) == 1 && c.txConns[nodes[0]] == nil {
			return nil, errors.ErrTransInMulti
		}
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

func (c *ClientConn) executeInNode(conn *backend.BackendConn, sql string, args []interface{}) ([]*mysql.Result, error) {
	var state string
	startTime := time.Now().UnixNano()
	r, err := conn.Execute(sql, args...)
	if err != nil {
		state = "ERROR"
	} else {
		state = "OK"
	}
	execTime := float64(time.Now().UnixNano()-startTime) / float64(time.Millisecond)
	if strings.ToLower(c.proxy.cfg.LogSql) != golog.LogSqlOff &&
		execTime > float64(c.proxy.cfg.SlowLogTime) {
		golog.OutputSql(state, "%.1fms - %s->%s:%s",
			execTime,
			c.c.RemoteAddr(),
			conn.GetAddr(),
			sql,
		)
	}

	if err != nil {
		return nil, err
	}

	return []*mysql.Result{r}, err
}

func (c *ClientConn) executeInMultiNodes(conns map[string]*backend.BackendConn, sqls map[string][]string, args []interface{}) ([]*mysql.Result, error) {
	if len(conns) != len(sqls) {
		golog.Error("ClientConn", "executeInMultiNodes", errors.ErrConnNotEqual.Error(), c.connectionId,
			"conns", conns,
			"sqls", sqls,
		)
		return nil, errors.ErrConnNotEqual
	}

	var wg sync.WaitGroup

	if len(conns) == 0 {
		return nil, errors.ErrNoPlan
	}

	wg.Add(len(conns))

	resultCount := 0
	for _, sqlSlice := range sqls {
		resultCount += len(sqlSlice)
	}

	rs := make([]interface{}, resultCount)

	f := func(rs []interface{}, i int, execSqls []string, co *backend.BackendConn) {
		var state string
		for _, v := range execSqls {
			startTime := time.Now().UnixNano()
			r, err := co.Execute(v, args...)
			if err != nil {
				state = "ERROR"
				rs[i] = err
			} else {
				state = "OK"
				rs[i] = r
			}
			execTime := float64(time.Now().UnixNano()-startTime) / float64(time.Millisecond)
			if c.proxy.cfg.LogSql != golog.LogSqlOff &&
				execTime > float64(c.proxy.cfg.SlowLogTime) {
				golog.OutputSql(state, "%.1fms - %s->%s:%s",
					execTime,
					c.c.RemoteAddr(),
					co.GetAddr(),
					v,
				)
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
	r := make([]*mysql.Result, resultCount)
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*mysql.Result)
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

func (c *ClientConn) newEmptyResultset(stmt *sqlparser.Select) *mysql.Resultset {
	r := new(mysql.Resultset)
	r.Fields = make([]*mysql.Field, len(stmt.SelectExprs))

	for i, expr := range stmt.SelectExprs {
		r.Fields[i] = &mysql.Field{}
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
	r.RowDatas = make([]mysql.RowData, 0)

	return r
}

func (c *ClientConn) GetTransExecNode(tokens []string, sql string) (*backend.Node, error) {
	var execNode *backend.Node
	var err error

	tokensLen := len(tokens)
	if 2 <= tokensLen {
		if tokens[0][0] == mysql.COMMENT_PREFIX {
			nodeName := strings.Trim(tokens[0], mysql.COMMENT_STRING)
			if c.schema.nodes[nodeName] != nil {
				execNode = c.schema.nodes[nodeName]
			}
		}
	}

	if execNode == nil {
		execNode, _, err = c.GetExecNode(tokens, sql)
		if err != nil {
			return nil, err
		}
		return execNode, nil
	}
	if len(c.txConns) == 1 && c.txConns[execNode] == nil {
		return nil, errors.ErrTransInMulti
	}
	return execNode, nil
}

func (c *ClientConn) GetExecNode(tokens []string,
	sql string) (*backend.Node, bool, error) {
	var execNode *backend.Node
	var fromSlave bool

	schema := c.proxy.schema
	rules := schema.rule.Rules

	tokensLen := len(tokens)
	if 0 < tokensLen {
		tokenId, ok := mysql.PARSE_TOKEN_MAP[strings.ToLower(tokens[0])]
		if ok == true {
			switch tokenId {
			case mysql.TK_ID_SELECT, mysql.TK_ID_DELETE:
				if len(rules) == 0 {
					if tokenId == mysql.TK_ID_SELECT {
						fromSlave = true
					}
					break
				}
				for i := 1; i < tokensLen; i++ {
					if strings.ToLower(tokens[i]) == mysql.TK_STR_FROM {
						if i+1 < tokensLen {
							tableName := sqlparser.GetTableName(tokens[i+1])

							if _, ok := rules[tableName]; ok {
								return nil, false, nil
							} else {
								if tokenId == mysql.TK_ID_SELECT {
									fromSlave = true
								}
							}
						}
					}
				}
			case mysql.TK_ID_INSERT, mysql.TK_ID_REPLACE:
				if len(rules) == 0 {
					break
				}
				for i := 0; i < tokensLen; i++ {
					if strings.ToLower(tokens[i]) == mysql.TK_STR_INTO {
						if i+1 < tokensLen {
							tableName := sqlparser.GetInsertTableName(tokens[i+1])
							if _, ok := rules[tableName]; ok {
								return nil, false, nil
							}
						}
					}
				}
			case mysql.TK_ID_UPDATE:
				if len(rules) == 0 {
					break
				}
				for i := 0; i < tokensLen; i++ {
					if strings.ToLower(tokens[i]) == mysql.TK_STR_SET {
						tableName := sqlparser.GetTableName(tokens[i-1])
						if _, ok := rules[tableName]; ok {
							return nil, false, nil
						}
					}
				}
			case mysql.TK_ID_SET:
				if len(tokens) < 2 {
					break
				}
				secondWord := strings.ToLower(tokens[1])
				if secondWord == mysql.TK_STR_NAMES ||
					secondWord == mysql.TK_STR_RESULTS ||
					secondWord == mysql.TK_STR_CLIENT ||
					secondWord == mysql.TK_STR_CONNECTION ||
					secondWord == mysql.TK_STR_AUTOCOMMIT {
					return nil, false, nil
				}
			default:
				return nil, false, nil
			}
		}
	}
	//get node
	if 2 <= tokensLen {
		if tokens[0][0] == mysql.COMMENT_PREFIX {
			nodeName := strings.Trim(tokens[0], mysql.COMMENT_STRING)
			if c.schema.nodes[nodeName] != nil {
				execNode = c.schema.nodes[nodeName]
			}
			//select
			if mysql.PARSE_TOKEN_MAP[tokens[1]] == mysql.TK_ID_SELECT {
				fromSlave = true
			}
		}
	}

	if execNode == nil {
		defaultRule := c.schema.rule.DefaultRule
		if len(defaultRule.Nodes) == 0 {
			return nil, false, errors.ErrNoDefaultNode
		}
		execNode = c.proxy.GetNode(defaultRule.Nodes[0])
	}

	return execNode, fromSlave, nil
}

//返回true表示已经处理，false表示未处理
func (c *ClientConn) preHandleShard(sql string) (bool, error) {
	var rs []*mysql.Result
	var err error

	var execNode *backend.Node
	var fromSlave bool = false

	if len(sql) == 0 {
		return false, errors.ErrCmdUnsupport
	}

	tokens := strings.Fields(sql)
	if len(tokens) == 0 {
		return false, errors.ErrCmdUnsupport
	}

	if c.isInTransaction() {
		execNode, err = c.GetTransExecNode(tokens, sql)
	} else {
		execNode, fromSlave, err = c.GetExecNode(tokens, sql)
	}

	if err != nil {
		return false, err
	}
	//need shard sql
	if execNode == nil {
		return false, nil
	}
	//execute in Master DB
	conn, err := c.getBackendConn(execNode, fromSlave)
	defer c.closeConn(conn, false)
	if err != nil {
		return false, err
	}
	rs, err = c.executeInNode(conn, sql, nil)
	if err != nil {
		return false, err
	}

	if len(rs) == 0 {
		msg := fmt.Sprintf("result is empty")
		golog.Error("ClientConn", "handleUnsupport", msg, c.connectionId)
		return false, mysql.NewError(mysql.ER_UNKNOWN_ERROR, msg)
	}

	if rs[0].Resultset != nil {
		err = c.writeResultset(c.status, rs[0].Resultset)
	} else {
		err = c.writeOK(rs[0])
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *ClientConn) handleExec(stmt sqlparser.Statement, args []interface{}) error {
	plan, err := c.schema.rule.BuildPlan(stmt)
	if err != nil {
		return err
	}
	conns, err := c.getShardConns(false, plan)
	defer c.closeShardConns(conns, err != nil)
	if err != nil {
		golog.Error("ClientConn", "handleExec", err.Error(), c.connectionId)
		return err
	}
	if conns == nil {
		return c.writeOK(nil)
	}

	var rs []*mysql.Result
	if 1 < len(conns) {
		return errors.ErrExecInMulti
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
	if err == nil {
		err = c.mergeExecResult(rs)
	}

	return err
}

func (c *ClientConn) mergeExecResult(rs []*mysql.Result) error {
	r := new(mysql.Result)
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
