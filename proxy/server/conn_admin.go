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
	"strconv"
	"strings"
	"time"

	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
)

const (
	Master = "master"
	Slave  = "slave"

	ServerRegion = "server"
	NodeRegion   = "node"

	ADMIN_OPT_ADD  = "add"
	ADMIN_OPT_DEL  = "del"
	ADMIN_OPT_UP   = "up"
	ADMIN_OPT_DOWN = "down"
	ADMIN_OPT_SHOW = "show"

	ADMIN_PROXY  = "proxy"
	ADMIN_NODE   = "node"
	ADMIN_SCHEMA = "schema"

	ADMIN_CONFIG = "config"
)

var cmdServerOrder = []string{"opt", "k", "v"}
var cmdNodeOrder = []string{"opt", "node", "k", "v"}

func (c *ClientConn) handleNodeCmd(rows sqlparser.InsertRows) error {
	var err error
	var opt, nodeName, role, addr string

	vals := rows.(sqlparser.Values)
	if len(vals) == 0 {
		return errors.ErrCmdUnsupport
	}

	tuple := vals[0].(sqlparser.ValTuple)
	if len(tuple) != len(cmdNodeOrder) {
		return errors.ErrCmdUnsupport
	}

	opt = sqlparser.String(tuple[0])
	opt = strings.Trim(opt, "'")

	nodeName = sqlparser.String(tuple[1])
	nodeName = strings.Trim(nodeName, "'")

	role = sqlparser.String(tuple[2])
	role = strings.Trim(role, "'")

	addr = sqlparser.String(tuple[3])
	addr = strings.Trim(addr, "'")

	switch strings.ToLower(opt) {
	case ADMIN_OPT_ADD:
		err = c.AddDatabase(
			nodeName,
			role,
			addr,
		)
	case ADMIN_OPT_DEL:
		err = c.DeleteDatabase(
			nodeName,
			role,
			addr,
		)

	case ADMIN_OPT_UP:
		err = c.UpDatabase(
			nodeName,
			role,
			addr,
		)
	case ADMIN_OPT_DOWN:
		err = c.DownDatabase(
			nodeName,
			role,
			addr,
		)
	default:
		err = errors.ErrCmdUnsupport
		golog.Error("ClientConn", "handleNodeCmd", err.Error(),
			c.connectionId, "opt", opt)
	}
	return err
}

func (c *ClientConn) handleServerCmd(rows sqlparser.InsertRows) (*mysql.Resultset, error) {
	var err error
	var result *mysql.Resultset
	var opt, k, v string

	vals := rows.(sqlparser.Values)
	if len(vals) == 0 {
		return nil, errors.ErrCmdUnsupport
	}

	tuple := vals[0].(sqlparser.ValTuple)
	if len(tuple) != len(cmdServerOrder) {
		return nil, errors.ErrCmdUnsupport
	}

	opt = sqlparser.String(tuple[0])
	opt = strings.Trim(opt, "'")

	k = sqlparser.String(tuple[1])
	k = strings.Trim(k, "'")

	v = sqlparser.String(tuple[2])
	v = strings.Trim(v, "'")

	switch strings.ToLower(opt) {
	case ADMIN_OPT_SHOW:
		result, err = c.handleAdminShow(k, v)
	default:
		err = errors.ErrCmdUnsupport
		golog.Error("ClientConn", "handleNodeCmd", err.Error(),
			c.connectionId, "opt", opt)
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ClientConn) AddDatabase(nodeName string, role string, addr string) error {
	//can not add a new master database
	if role != Slave {
		return errors.ErrCmdUnsupport
	}

	return c.proxy.AddSlave(nodeName, addr)
}

func (c *ClientConn) DeleteDatabase(nodeName string, role string, addr string) error {
	//can not delete a master database
	if role != Slave {
		return errors.ErrCmdUnsupport
	}

	return c.proxy.DeleteSlave(nodeName, addr)
}

func (c *ClientConn) UpDatabase(nodeName string, role string, addr string) error {
	if role != Master && role != Slave {
		return errors.ErrCmdUnsupport
	}
	if role == Master {
		return c.proxy.UpMaster(nodeName, addr)
	}

	return c.proxy.UpSlave(nodeName, addr)
}

func (c *ClientConn) DownDatabase(nodeName string, role string, addr string) error {
	if role != Master && role != Slave {
		return errors.ErrCmdUnsupport
	}
	if role == Master {
		return c.proxy.DownMaster(nodeName, addr)
	}

	return c.proxy.DownSlave(nodeName, addr)
}

func (c *ClientConn) checkCmdOrder(region string, columns sqlparser.Columns) error {
	var cmdOrder []string
	node := sqlparser.SelectExprs(columns)

	switch region {
	case NodeRegion:
		cmdOrder = cmdNodeOrder
	case ServerRegion:
		cmdOrder = cmdServerOrder
	default:
		return errors.ErrCmdUnsupport
	}

	for i := 0; i < len(node); i++ {
		val := sqlparser.String(node[i])
		if val != cmdOrder[i] {
			return errors.ErrCmdUnsupport
		}
	}

	return nil
}

func (c *ClientConn) handleAdmin(admin *sqlparser.Admin) error {
	var err error
	var result *mysql.Resultset

	region := sqlparser.String(admin.Region)

	err = c.checkCmdOrder(region, admin.Columns)
	if err != nil {
		return err
	}

	switch strings.ToLower(region) {
	case NodeRegion:
		err = c.handleNodeCmd(admin.Rows)
	case ServerRegion:
		result, err = c.handleServerCmd(admin.Rows)
	default:
		return fmt.Errorf("admin %s not supported now", region)
	}

	if err != nil {
		golog.Error("ClientConn", "handleAdmin", err.Error(),
			c.connectionId, "sql", sqlparser.String(admin))
		return err
	}

	if result != nil {
		return c.writeResultset(c.status, result)
	}

	return c.writeOK(nil)
}

func (c *ClientConn) handleAdminShow(k, v string) (*mysql.Resultset, error) {
	if len(k) == 0 || len(v) == 0 {
		return nil, errors.ErrCmdUnsupport
	}
	if k == ADMIN_PROXY && v == ADMIN_CONFIG {
		return c.handleShowProxyConfig()
	}

	if k == ADMIN_NODE && v == ADMIN_CONFIG {
		return c.handleShowNodeConfig()
	}

	if k == ADMIN_SCHEMA && v == ADMIN_CONFIG {
		return c.handleShowSchemaConfig()
	}

	return nil, errors.ErrCmdUnsupport
}

func (c *ClientConn) handleShowProxyConfig() (*mysql.Resultset, error) {
	var names []string = []string{"Key", "Value"}
	var rows [][]string
	var nodeNames []string

	const (
		Column = 2
	)
	for name := range c.schema.nodes {
		nodeNames = append(nodeNames, name)
	}

	rows = append(rows, []string{"Addr", c.proxy.cfg.Addr})
	rows = append(rows, []string{"User", c.proxy.cfg.User})
	rows = append(rows, []string{"LogPath", c.proxy.cfg.LogPath})
	rows = append(rows, []string{"LogLevel", c.proxy.cfg.LogLevel})
	rows = append(rows, []string{"LogSql", c.proxy.cfg.LogSql})
	rows = append(rows, []string{"SlowLogTime", strconv.Itoa(c.proxy.cfg.SlowLogTime)})
	rows = append(rows, []string{"Nodes_Count", fmt.Sprintf("%d", len(c.proxy.nodes))})
	rows = append(rows, []string{"Nodes_List", strings.Join(nodeNames, ",")})
	rows = append(rows, []string{"ConnCount", fmt.Sprintf("%d", c.proxy.counter.connCnt)})
	rows = append(rows, []string{"OpCount", fmt.Sprintf("%d", c.proxy.counter.oldOpCnt)})
	rows = append(rows, []string{"ErrorCount", fmt.Sprintf("%d", c.proxy.counter.oldErrorCnt)})
	rows = append(rows, []string{"SlowCount", fmt.Sprintf("%d", c.proxy.counter.oldSlowCnt)})

	var values [][]interface{} = make([][]interface{}, len(rows))
	for i := range rows {
		values[i] = make([]interface{}, Column)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	return c.buildResultset(nil, names, values)
}

func (c *ClientConn) handleShowNodeConfig() (*mysql.Resultset, error) {
	var names []string = []string{
		"Node",
		"Address",
		"Type",
		"State",
		"LastPing",
		"MaxConn",
		"IdleConn",
	}
	var rows [][]string
	const (
		Column = 7
	)

	//var nodeRows [][]string
	for name, node := range c.schema.nodes {
		//"master"
		rows = append(
			rows,
			[]string{
				name,
				node.Master.Addr(),
				"master",
				node.Master.State(),
				fmt.Sprintf("%v", time.Unix(node.LastMasterPing, 0)),
				strconv.Itoa(node.Cfg.MaxConnNum),
				strconv.Itoa(node.Master.IdleConnCount()),
			})
		//"slave"
		for _, slave := range node.Slave {
			if slave != nil {
				rows = append(
					rows,
					[]string{
						name,
						slave.Addr(),
						"slave",
						slave.State(),
						fmt.Sprintf("%v", time.Unix(node.LastSlavePing, 0)),
						strconv.Itoa(node.Cfg.MaxConnNum),
						strconv.Itoa(slave.IdleConnCount()),
					})
			}
		}
	}
	//rows = append(rows, nodeRows...)
	var values [][]interface{} = make([][]interface{}, len(rows))
	for i := range rows {
		values[i] = make([]interface{}, Column)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	return c.buildResultset(nil, names, values)
}

func (c *ClientConn) handleShowSchemaConfig() (*mysql.Resultset, error) {
	var Column = 7
	var rows [][]string
	var names []string = []string{
		"DB",
		"Table",
		"Type",
		"Key",
		"Nodes_List",
		"Locations",
		"TableRowLimit",
	}

	//default Rule
	var defaultRule = c.schema.rule.DefaultRule
	rows = append(
		rows,
		[]string{
			defaultRule.DB,
			defaultRule.Table,
			defaultRule.Type,
			defaultRule.Key,
			strings.Join(defaultRule.Nodes, ", "),
			"",
			"0",
		},
	)

	schemaConfig := c.proxy.cfg.Schema
	shardRule := schemaConfig.ShardRule

	for _, r := range shardRule {
		rows = append(
			rows,
			[]string{
				schemaConfig.DB,
				r.Table,
				r.Type,
				r.Key,
				strings.Join(r.Nodes, ", "),
				arrayToString(r.Locations),
				strconv.Itoa(r.TableRowLimit),
			},
		)
	}

	var values [][]interface{} = make([][]interface{}, len(rows))
	for i := range rows {
		values[i] = make([]interface{}, Column)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	return c.buildResultset(nil, names, values)
}

func arrayToString(array []int) string {
	if len(array) == 0 {
		return ""
	}
	var strArray []string
	for _, v := range array {
		strArray = append(strArray, strconv.FormatInt(int64(v), 10))
	}

	return strings.Join(strArray, ", ")
}
