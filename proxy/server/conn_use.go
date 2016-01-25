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

	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
)

func (c *ClientConn) handleUseDB(stmt *sqlparser.UseDB) error {
	if len(stmt.DB) == 0 {
		return fmt.Errorf("must have database, not %s", sqlparser.String(stmt))
	}
	if c.schema == nil {
		return mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
	}

	nodeName := c.schema.rule.DefaultRule.Nodes[0]

	n := c.proxy.GetNode(nodeName)
	co, err := n.GetMasterConn()
	defer c.closeConn(co, false)
	if err != nil {
		return err
	}

	if err = co.UseDB(string(stmt.DB)); err != nil {
		return err
	}
	c.db = string(stmt.DB)
	return c.writeOK(nil)
}
