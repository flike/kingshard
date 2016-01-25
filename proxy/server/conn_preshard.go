// Copyright 2016 The kingshard Authors. All rights reserved.
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
	"strings"

	"github.com/flike/kingshard/backend"
	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
)

type ExecuteDB struct {
	ExecNode *backend.Node
	IsSlave  bool
}

func (c *ClientConn) isBlacklistSql(sql string) bool {
	fingerprint := mysql.GetFingerprint(sql)
	md5 := mysql.GetMd5(fingerprint)
	if _, ok := c.proxy.blacklistSqls.sqls[md5]; ok {
		return true
	}
	return false
}

//preprocessing sql before parse sql
func (c *ClientConn) preHandleShard(sql string) (bool, error) {
	var rs []*mysql.Result
	var err error
	var executeDB *ExecuteDB

	if len(sql) == 0 {
		return false, errors.ErrCmdUnsupport
	}
	//filter the blacklist sql
	if c.proxy.blacklistSqls.sqlsLen != 0 {
		if c.isBlacklistSql(sql) {
			golog.OutputSql("Forbidden", "%s->%s:%s",
				c.c.RemoteAddr(),
				c.proxy.addr,
				sql,
			)
			err := mysql.NewError(mysql.ER_UNKNOWN_ERROR, "sql in blacklist.")
			return false, err
		}
	}

	tokens := strings.Fields(sql)
	if len(tokens) == 0 {
		return false, errors.ErrCmdUnsupport
	}

	if c.isInTransaction() {
		executeDB, err = c.GetTransExecDB(tokens, sql)
	} else {
		executeDB, err = c.GetExecDB(tokens, sql)
	}

	if err != nil {
		return false, err
	}
	//need shard sql
	if executeDB == nil {
		return false, nil
	}
	//get connection in DB
	conn, err := c.getBackendConn(executeDB.ExecNode, executeDB.IsSlave)
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
		golog.Error("ClientConn", "handleUnsupport", msg, 0, "sql", sql)
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

func (c *ClientConn) GetTransExecDB(tokens []string, sql string) (*ExecuteDB, error) {
	var err error
	tokensLen := len(tokens)
	executeDB := new(ExecuteDB)

	if 2 <= tokensLen {
		if tokens[0][0] == mysql.COMMENT_PREFIX {
			nodeName := strings.Trim(tokens[0], mysql.COMMENT_STRING)
			if c.schema.nodes[nodeName] != nil {
				executeDB.ExecNode = c.schema.nodes[nodeName]
			}
		}
	}

	if executeDB.ExecNode == nil {
		executeDB, err = c.GetExecDB(tokens, sql)
		if err != nil {
			return nil, err
		}
		if executeDB == nil {
			return nil, nil
		}
		//transaction execute in master db
		executeDB.IsSlave = false
		return executeDB, nil
	}
	if len(c.txConns) == 1 && c.txConns[executeDB.ExecNode] == nil {
		return nil, errors.ErrTransInMulti
	}
	return executeDB, nil
}

//if sql need shard return nil, else return the unshard db
func (c *ClientConn) GetExecDB(tokens []string, sql string) (*ExecuteDB, error) {
	tokensLen := len(tokens)
	if 0 < tokensLen {
		tokenId, ok := mysql.PARSE_TOKEN_MAP[strings.ToLower(tokens[0])]
		if ok == true {
			switch tokenId {
			case mysql.TK_ID_SELECT:
				return c.getSelectExecDB(tokens, tokensLen)
			case mysql.TK_ID_DELETE:
				return c.getDeleteExecDB(tokens, tokensLen)
			case mysql.TK_ID_INSERT, mysql.TK_ID_REPLACE:
				return c.getInsertOrReplaceExecDB(tokens, tokensLen)
			case mysql.TK_ID_UPDATE:
				return c.getUpdateExecDB(tokens, tokensLen)
			case mysql.TK_ID_SET:
				return c.getSetExecDB(tokens, tokensLen, sql)
			default:
				return nil, nil
			}
		}
	}
	executeDB := new(ExecuteDB)
	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}
	return executeDB, nil
}

func (c *ClientConn) setExecuteNode(tokens []string, tokensLen int, executeDB *ExecuteDB) error {
	if 2 <= tokensLen {
		if tokens[0][0] == mysql.COMMENT_PREFIX {
			nodeName := strings.Trim(tokens[0], mysql.COMMENT_STRING)
			if c.schema.nodes[nodeName] != nil {
				executeDB.ExecNode = c.schema.nodes[nodeName]
			}
		}
	}

	if executeDB.ExecNode == nil {
		defaultRule := c.schema.rule.DefaultRule
		if len(defaultRule.Nodes) == 0 {
			return errors.ErrNoDefaultNode
		}
		executeDB.ExecNode = c.proxy.GetNode(defaultRule.Nodes[0])
	}

	return nil
}

//get the execute database for select sql
func (c *ClientConn) getSelectExecDB(tokens []string, tokensLen int) (*ExecuteDB, error) {
	executeDB := new(ExecuteDB)
	schema := c.proxy.schema

	rules := schema.rule.Rules
	executeDB.IsSlave = true

	if len(rules) != 0 {
		for i := 1; i < tokensLen; i++ {
			if strings.ToLower(tokens[i]) == mysql.TK_STR_FROM {
				if i+1 < tokensLen {
					tableName := sqlparser.GetTableName(tokens[i+1])
					if _, ok := rules[tableName]; ok {
						return nil, nil
					}
				}
			}
		}
	}

	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	return executeDB, nil
}

//get the execute database for delete sql
func (c *ClientConn) getDeleteExecDB(tokens []string, tokensLen int) (*ExecuteDB, error) {
	executeDB := new(ExecuteDB)
	schema := c.proxy.schema
	rules := schema.rule.Rules

	if len(rules) != 0 {
		for i := 1; i < tokensLen; i++ {
			if strings.ToLower(tokens[i]) == mysql.TK_STR_FROM {
				if i+1 < tokensLen {
					tableName := sqlparser.GetTableName(tokens[i+1])
					if _, ok := rules[tableName]; ok {
						return nil, nil
					}
				}
			}
		}
	}

	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	return executeDB, nil
}

//get the execute database for insert or replace sql
func (c *ClientConn) getInsertOrReplaceExecDB(tokens []string, tokensLen int) (*ExecuteDB, error) {
	executeDB := new(ExecuteDB)
	schema := c.proxy.schema
	rules := schema.rule.Rules

	if len(rules) != 0 {
		for i := 0; i < tokensLen; i++ {
			if strings.ToLower(tokens[i]) == mysql.TK_STR_INTO {
				if i+1 < tokensLen {
					tableName := sqlparser.GetInsertTableName(tokens[i+1])
					if _, ok := rules[tableName]; ok {
						return nil, nil
					}
				}
			}
		}
	}

	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	return executeDB, nil
}

//get the execute database for update sql
func (c *ClientConn) getUpdateExecDB(tokens []string, tokensLen int) (*ExecuteDB, error) {
	executeDB := new(ExecuteDB)
	schema := c.proxy.schema
	rules := schema.rule.Rules

	if len(rules) != 0 {
		for i := 0; i < tokensLen; i++ {
			if strings.ToLower(tokens[i]) == mysql.TK_STR_SET {
				tableName := sqlparser.GetTableName(tokens[i-1])
				if _, ok := rules[tableName]; ok {
					return nil, nil
				}
			}
		}
	}

	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	return executeDB, nil
}

//get the execute database for set sql
func (c *ClientConn) getSetExecDB(tokens []string, tokensLen int, sql string) (*ExecuteDB, error) {
	executeDB := new(ExecuteDB)

	//handle two styles: set autocommit= 0 or set autocommit = 0
	if 2 < len(tokens) {
		before := strings.Split(sql, "=")
		//uncleanWorld is 'autocommit' or 'autocommit '
		uncleanWord := strings.Split(before[0], " ")
		secondWord := strings.ToLower(uncleanWord[1])
		if secondWord == mysql.TK_STR_NAMES ||
			secondWord == mysql.TK_STR_RESULTS ||
			secondWord == mysql.TK_STR_CLIENT ||
			secondWord == mysql.TK_STR_CONNECTION ||
			secondWord == mysql.TK_STR_AUTOCOMMIT {
			return nil, nil
		}
	}

	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	return executeDB, nil
}
