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
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
)

var nstring = sqlparser.String

func (c *ClientConn) handleSet(stmt *sqlparser.Set, sql string) error {
	if len(stmt.Exprs) != 1 && len(stmt.Exprs) != 2 {
		return fmt.Errorf("must set one item once, not %s", nstring(stmt))
	}

	k := string(stmt.Exprs[0].Name.Name)
	switch strings.ToUpper(k) {
	case `AUTOCOMMIT`:
		return c.handleSetAutoCommit(stmt.Exprs[0].Expr)
	case `NAMES`, `CHARACTER_SET_RESULTS`, `CHARACTER_SET_CLIENT`, `CHARACTER_SET_CONNECTION`:
		if len(stmt.Exprs) == 2 {
			c.handleSetNames(stmt.Exprs[0].Expr, stmt.Exprs[1].Expr)
		}
		return c.handleSetNames(stmt.Exprs[0].Expr, nil)
	default:
		golog.Error("ClientConn", "handleSet", "command not supported",
			c.connectionId, "sql", sql)
		return c.writeOK(nil)
	}
}

func (c *ClientConn) handleSetAutoCommit(val sqlparser.ValExpr) error {
	value, ok := val.(sqlparser.NumVal)
	if !ok {
		return fmt.Errorf("set autocommit error")
	}
	switch value[0] {
	case '1':
		c.status |= mysql.SERVER_STATUS_AUTOCOMMIT
		if c.status&mysql.SERVER_STATUS_IN_TRANS > 0 {
			c.status &= ^mysql.SERVER_STATUS_IN_TRANS
		}
		for _, co := range c.txConns {
			if e := co.SetAutoCommit(1); e != nil {
				co.Close()
				c.txConns = make(map[*backend.Node]*backend.BackendConn)
				return fmt.Errorf("set autocommit error, %v", e)
			}
			co.Close()
		}
		c.txConns = make(map[*backend.Node]*backend.BackendConn)
	case '0':
		c.status &= ^mysql.SERVER_STATUS_AUTOCOMMIT
	default:
		return fmt.Errorf("invalid autocommit flag %s", value)
	}

	return c.writeOK(nil)
}

func (c *ClientConn) handleSetNames(ch, ci sqlparser.ValExpr) error {
	var cid mysql.CollationId
	var ok bool

	value := sqlparser.String(ch)
	value = strings.Trim(value, "'`\"")

	charset := strings.ToLower(value)
	if charset == "null" {
		return c.writeOK(nil)
	}
	if ci == nil {
		cid, ok = mysql.CharsetIds[charset]
		if !ok {
			return fmt.Errorf("invalid charset %s", charset)
		}
	} else {
		collate := sqlparser.String(ci)
		collate = strings.Trim(value, "'`\"")
		cid, ok = mysql.CollationNames[collate]
		if !ok {
			return fmt.Errorf("invalid charset %s", charset)
		}
	}
	c.charset = charset
	c.collation = cid

	return c.writeOK(nil)
}
