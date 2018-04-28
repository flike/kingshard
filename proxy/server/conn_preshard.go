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
	"github.com/flike/kingshard/core/hack"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/proxy/router"
	"github.com/flike/kingshard/sqlparser"
)

type ExecuteDB struct {
	ExecNode *backend.Node
	IsSlave  bool
	sql      string
}

func (c *ClientConn) isBlacklistSql(sql string) bool {
	fingerprint := mysql.GetFingerprint(sql)
	md5 := mysql.GetMd5(fingerprint)
	if _, ok := c.proxy.blacklistSqls[c.proxy.blacklistSqlsIndex].sqls[md5]; ok {
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
	if c.proxy.blacklistSqls[c.proxy.blacklistSqlsIndex].sqlsLen != 0 {
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

	tokens := strings.FieldsFunc(sql, hack.IsSqlSep)

	if len(tokens) == 0 {
		return false, errors.ErrCmdUnsupport
	}

	if c.isInTransaction() {
		executeDB, err = c.GetTransExecDB(tokens, sql)
	} else {
		executeDB, err = c.GetExecDB(tokens, sql)
	}

	if err != nil {
		//this SQL doesn't need execute in the backend.
		if err == errors.ErrIgnoreSQL {
			err = c.writeOK(nil)
			if err != nil {
				return false, err
			}
			return true, nil
		}
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
	//execute.sql may be rewritten in getShowExecDB
	rs, err = c.executeInNode(conn, executeDB.sql, nil)
	if err != nil {
		return false, err
	}

	if len(rs) == 0 {
		msg := fmt.Sprintf("result is empty")
		golog.Error("ClientConn", "handleUnsupport", msg, 0, "sql", sql)
		return false, mysql.NewError(mysql.ER_UNKNOWN_ERROR, msg)
	}

	c.lastInsertId = int64(rs[0].InsertId)
	c.affectedRows = int64(rs[0].AffectedRows)

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
	executeDB.sql = sql

	//transaction execute in master db
	executeDB.IsSlave = false

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
				return c.getSelectExecDB(sql, tokens, tokensLen)
			case mysql.TK_ID_DELETE:
				return c.getDeleteExecDB(sql, tokens, tokensLen)
			case mysql.TK_ID_INSERT, mysql.TK_ID_REPLACE:
				return c.getInsertOrReplaceExecDB(sql, tokens, tokensLen)
			case mysql.TK_ID_UPDATE:
				return c.getUpdateExecDB(sql, tokens, tokensLen)
			case mysql.TK_ID_SET:
				return c.getSetExecDB(sql, tokens, tokensLen)
			case mysql.TK_ID_SHOW:
				return c.getShowExecDB(sql, tokens, tokensLen)
			case mysql.TK_ID_TRUNCATE:
				return c.getTruncateExecDB(sql, tokens, tokensLen)
			default:
				return nil, nil
			}
		}
	}
	executeDB := new(ExecuteDB)
	executeDB.sql = sql
	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}
	return executeDB, nil
}

func (c *ClientConn) setExecuteNode(tokens []string, tokensLen int, executeDB *ExecuteDB) error {
	if 2 <= tokensLen {
		//for /*node1*/
		if 1 < len(tokens) && tokens[0][0] == mysql.COMMENT_PREFIX {
			nodeName := strings.Trim(tokens[0], mysql.COMMENT_STRING)
			if c.schema.nodes[nodeName] != nil {
				executeDB.ExecNode = c.schema.nodes[nodeName]
			}
			//for /*node1*/ select
			if strings.ToLower(tokens[1]) == mysql.TK_STR_SELECT {
				executeDB.IsSlave = true
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
func (c *ClientConn) getSelectExecDB(sql string, tokens []string, tokensLen int) (*ExecuteDB, error) {
	var ruleDB string
	executeDB := new(ExecuteDB)
	executeDB.sql = sql
	executeDB.IsSlave = true

	schema := c.schema
	router := schema.rule
	rules := router.Rules

	if len(rules) != 0 {
		for i := 1; i < tokensLen; i++ {
			if strings.ToLower(tokens[i]) == mysql.TK_STR_FROM {
				if i+1 < tokensLen {
					DBName, tableName := sqlparser.GetDBTable(tokens[i+1])
					//if the token[i+1] like this:kingshard.test_shard_hash
					if DBName != "" {
						ruleDB = DBName
					} else {
						ruleDB = c.db
					}
					if router.GetRule(ruleDB, tableName) != router.DefaultRule {
						return nil, nil
					} else {
						//if the table is not shard table,send the sql
						//to default db
						break
					}
				}
			}

			if strings.ToLower(tokens[i]) == mysql.TK_STR_LAST_INSERT_ID {
				return nil, nil
			}
		}
	}

	//if send to master
	if 2 < tokensLen {
		if strings.ToLower(tokens[1]) == mysql.TK_STR_MASTER_HINT {
			executeDB.IsSlave = false
		}
	}
	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}
	return executeDB, nil
}

//get the execute database for delete sql
func (c *ClientConn) getDeleteExecDB(sql string, tokens []string, tokensLen int) (*ExecuteDB, error) {
	var ruleDB string
	executeDB := new(ExecuteDB)
	executeDB.sql = sql
	schema := c.schema
	router := schema.rule
	rules := router.Rules

	if len(rules) != 0 {
		for i := 1; i < tokensLen; i++ {
			if strings.ToLower(tokens[i]) == mysql.TK_STR_FROM {
				if i+1 < tokensLen {
					DBName, tableName := sqlparser.GetDBTable(tokens[i+1])
					//if the token[i+1] like this:kingshard.test_shard_hash
					if DBName != "" {
						ruleDB = DBName
					} else {
						ruleDB = c.db
					}
					if router.GetRule(ruleDB, tableName) != router.DefaultRule {
						return nil, nil
					} else {
						break
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
func (c *ClientConn) getInsertOrReplaceExecDB(sql string, tokens []string, tokensLen int) (*ExecuteDB, error) {
	var ruleDB string
	executeDB := new(ExecuteDB)
	executeDB.sql = sql
	schema := c.schema
	router := schema.rule
	rules := router.Rules

	if len(rules) != 0 {
		for i := 0; i < tokensLen; i++ {
			if strings.ToLower(tokens[i]) == mysql.TK_STR_INTO {
				if i+1 < tokensLen {
					DBName, tableName := sqlparser.GetInsertDBTable(tokens[i+1])
					//if the token[i+1] like this:kingshard.test_shard_hash
					if DBName != "" {
						ruleDB = DBName
					} else {
						ruleDB = c.db
					}
					if router.GetRule(ruleDB, tableName) != router.DefaultRule {
						return nil, nil
					} else {
						break
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
func (c *ClientConn) getUpdateExecDB(sql string, tokens []string, tokensLen int) (*ExecuteDB, error) {
	var ruleDB string
	executeDB := new(ExecuteDB)
	executeDB.sql = sql
	schema := c.schema
	router := schema.rule
	rules := router.Rules

	if len(rules) != 0 {
		for i := 0; i < tokensLen; i++ {
			if strings.ToLower(tokens[i]) == mysql.TK_STR_SET {
				DBName, tableName := sqlparser.GetDBTable(tokens[i-1])
				//if the token[i+1] like this:kingshard.test_shard_hash
				if DBName != "" {
					ruleDB = DBName
				} else {
					ruleDB = c.db
				}
				if router.GetRule(ruleDB, tableName) != router.DefaultRule {
					return nil, nil
				} else {
					break
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
func (c *ClientConn) getSetExecDB(sql string, tokens []string, tokensLen int) (*ExecuteDB, error) {
	executeDB := new(ExecuteDB)
	executeDB.sql = sql

	//handle three styles:
	//set autocommit= 0
	//set autocommit = 0
	//set autocommit=0
	if 2 <= len(tokens) {
		before := strings.Split(sql, "=")
		//uncleanWorld is 'autocommit' or 'autocommit '
		uncleanWord := strings.Split(before[0], " ")
		secondWord := strings.ToLower(uncleanWord[1])
		if _, ok := mysql.SET_KEY_WORDS[secondWord]; ok {
			return nil, nil
		}

		//SET [gobal/session] TRANSACTION ISOLATION LEVEL SERIALIZABLE
		//ignore this sql
		if 3 <= len(uncleanWord) {
			if strings.ToLower(uncleanWord[1]) == mysql.TK_STR_TRANSACTION ||
				strings.ToLower(uncleanWord[2]) == mysql.TK_STR_TRANSACTION {
				return nil, errors.ErrIgnoreSQL
			}
		}
	}

	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	return executeDB, nil
}

//get the execute database for show sql
//choose slave preferentially
//tokens[0] is show
func (c *ClientConn) getShowExecDB(sql string, tokens []string, tokensLen int) (*ExecuteDB, error) {
	executeDB := new(ExecuteDB)
	executeDB.IsSlave = true
	executeDB.sql = sql

	//handle show columns/fields
	err := c.handleShowColumns(sql, tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	err = c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	return executeDB, nil
}

//handle show columns/fields
func (c *ClientConn) handleShowColumns(sql string, tokens []string,
	tokensLen int, executeDB *ExecuteDB) error {
	var ruleDB string
	for i := 0; i < tokensLen; i++ {
		tokens[i] = strings.ToLower(tokens[i])
		//handle SQL:
		//SHOW [FULL] COLUMNS FROM tbl_name [FROM db_name] [like_or_where]
		if (strings.ToLower(tokens[i]) == mysql.TK_STR_FIELDS ||
			strings.ToLower(tokens[i]) == mysql.TK_STR_COLUMNS) &&
			i+2 < tokensLen {
			if strings.ToLower(tokens[i+1]) == mysql.TK_STR_FROM {
				tableName := strings.Trim(tokens[i+2], "`")
				//get the ruleDB
				if i+4 < tokensLen && strings.ToLower(tokens[i+1]) == mysql.TK_STR_FROM {
					ruleDB = strings.Trim(tokens[i+4], "`")
				} else {
					ruleDB = c.db
				}
				showRouter := c.schema.rule
				showRule := showRouter.GetRule(ruleDB, tableName)
				//this SHOW is sharding SQL
				if showRule.Type != router.DefaultRuleType {
					if 0 < len(showRule.SubTableIndexs) {
						tableIndex := showRule.SubTableIndexs[0]
						nodeIndex := showRule.TableToNode[tableIndex]
						nodeName := showRule.Nodes[nodeIndex]
						tokens[i+2] = fmt.Sprintf("%s_%04d", tableName, tableIndex)
						executeDB.sql = strings.Join(tokens, " ")
						executeDB.ExecNode = c.schema.nodes[nodeName]
						return nil
					}
				}
			}
		}
	}
	return nil
}

//get the execute database for truncate sql
//sql: TRUNCATE [TABLE] tbl_name
func (c *ClientConn) getTruncateExecDB(sql string, tokens []string, tokensLen int) (*ExecuteDB, error) {
	var ruleDB string
	executeDB := new(ExecuteDB)
	executeDB.sql = sql
	schema := c.schema
	router := schema.rule
	rules := router.Rules
	if len(rules) != 0 && tokensLen >= 2 {
		DBName, tableName := sqlparser.GetDBTable(tokens[tokensLen-1])
		//if the token[i+1] like this:kingshard.test_shard_hash
		if DBName != "" {
			ruleDB = DBName
		} else {
			ruleDB = c.db
		}
		if router.GetRule(ruleDB, tableName) != router.DefaultRule {
			return nil, nil
		}

	}

	err := c.setExecuteNode(tokens, tokensLen, executeDB)
	if err != nil {
		return nil, err
	}

	return executeDB, nil
}
