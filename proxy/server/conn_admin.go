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
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/core/hack"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
)

const (
	Master = "master"
	Slave  = "slave"

	ServerRegion = "server"
	NodeRegion   = "node"

	//op
	ADMIN_OPT_ADD     = "add"
	ADMIN_OPT_DEL     = "del"
	ADMIN_OPT_UP      = "up"
	ADMIN_OPT_DOWN    = "down"
	ADMIN_OPT_SHOW    = "show"
	ADMIN_OPT_CHANGE  = "change"
	ADMIN_SAVE_CONFIG = "save"

	ADMIN_PROXY         = "proxy"
	ADMIN_NODE          = "node"
	ADMIN_SCHEMA        = "schema"
	ADMIN_LOG_SQL       = "log_sql"
	ADMIN_SLOW_LOG_TIME = "slow_log_time"
	ADMIN_ALLOW_IP      = "allow_ip"
	ADMIN_BLACK_SQL     = "black_sql"

	ADMIN_CONFIG = "config"
	ADMIN_STATUS = "status"
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
	case ADMIN_OPT_CHANGE:
		err = c.handleAdminChange(k, v)
	case ADMIN_OPT_ADD:
		err = c.handleAdminAdd(k, v)
	case ADMIN_OPT_DEL:
		err = c.handleAdminDelete(k, v)
	case ADMIN_SAVE_CONFIG:
		err = c.handleAdminSave(k, v)
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

func (c *ClientConn) handleAdminHelp(ah *sqlparser.AdminHelp) error {
	var Column = 2
	var rows [][]string
	var names []string = []string{"command", "description"}
	relativePath := "/doc/KingDoc/command_help"

	execPath, err := os.Getwd()
	if err != nil {
		return err
	}

	helpFilePath := execPath + relativePath
	file, err := os.Open(helpFilePath)
	if err != nil {
		return err
	}

	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		//end of file
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//parse the command description with '|' separating
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			cmdStr := strings.SplitN(line, "|", 2)
			if len(cmdStr) == 2 {
				rows = append(rows,
					[]string{
						strings.TrimSpace(cmdStr[0]),
						strings.TrimSpace(cmdStr[1]),
					},
				)
			}
		}
	}

	var values [][]interface{} = make([][]interface{}, len(rows))
	for i := range rows {
		values[i] = make([]interface{}, Column)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	result, err := c.buildResultset(nil, names, values)
	if err != nil {
		return err
	}
	return c.writeResultset(c.status, result)
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

	if k == ADMIN_PROXY && v == ADMIN_STATUS {
		return c.handleShowProxyStatus()
	}

	if k == ADMIN_NODE && v == ADMIN_CONFIG {
		return c.handleShowNodeConfig()
	}

	if k == ADMIN_SCHEMA && v == ADMIN_CONFIG {
		return c.handleShowSchemaConfig()
	}

	if k == ADMIN_ALLOW_IP && v == ADMIN_CONFIG {
		return c.handleShowAllowIPConfig()
	}

	if k == ADMIN_BLACK_SQL && v == ADMIN_CONFIG {
		return c.handleShowBlackSqlConfig()
	}

	return nil, errors.ErrCmdUnsupport
}

func (c *ClientConn) handleAdminChange(k, v string) error {
	if len(k) == 0 || len(v) == 0 {
		return errors.ErrCmdUnsupport
	}
	if k == ADMIN_LOG_SQL {
		return c.handleChangeLogSql(v)
	}

	if k == ADMIN_SLOW_LOG_TIME {
		return c.handleChangeSlowLogTime(v)
	}

	if k == ADMIN_PROXY {
		return c.handleChangeProxy(v)
	}

	return errors.ErrCmdUnsupport
}

func (c *ClientConn) handleAdminAdd(k, v string) error {
	if len(k) == 0 || len(v) == 0 {
		return errors.ErrCmdUnsupport
	}
	if k == ADMIN_ALLOW_IP {
		return c.handleAddAllowIP(v)
	}

	if k == ADMIN_BLACK_SQL {
		return c.handleAddBlackSql(v)
	}

	return errors.ErrCmdUnsupport
}

func (c *ClientConn) handleAdminDelete(k, v string) error {
	if len(k) == 0 || len(v) == 0 {
		return errors.ErrCmdUnsupport
	}
	if k == ADMIN_ALLOW_IP {
		return c.handleDelAllowIP(v)
	}

	if k == ADMIN_BLACK_SQL {
		return c.handleDelBlackSql(v)
	}

	return errors.ErrCmdUnsupport
}

func (c *ClientConn) handleShowProxyConfig() (*mysql.Resultset, error) {
	var names []string = []string{"Key", "Value"}
	var rows [][]string
	var nodeNames []string
	var users []string

	const (
		Column = 2
	)
	for name := range c.schema.nodes {
		nodeNames = append(nodeNames, name)
	}
	for user, _ := range c.proxy.users {
		users = append(users, user)
	}

	rows = append(rows, []string{"Addr", c.proxy.cfg.Addr})
	rows = append(rows, []string{"User_List", strings.Join(users, ",")})
	rows = append(rows, []string{"LogPath", c.proxy.cfg.LogPath})
	rows = append(rows, []string{"LogLevel", c.proxy.cfg.LogLevel})
	rows = append(rows, []string{"LogSql", c.proxy.logSql[c.proxy.logSqlIndex]})
	rows = append(rows, []string{"SlowLogTime", strconv.Itoa(c.proxy.slowLogTime[c.proxy.slowLogTimeIndex])})
	rows = append(rows, []string{"Nodes_Count", fmt.Sprintf("%d", len(c.proxy.nodes))})
	rows = append(rows, []string{"Nodes_List", strings.Join(nodeNames, ",")})
	rows = append(rows, []string{"ClientConns", fmt.Sprintf("%d", c.proxy.counter.ClientConns)})
	rows = append(rows, []string{"ClientQPS", fmt.Sprintf("%d", c.proxy.counter.OldClientQPS)})
	rows = append(rows, []string{"ErrLogTotal", fmt.Sprintf("%d", c.proxy.counter.OldErrLogTotal)})
	rows = append(rows, []string{"SlowLogTotal", fmt.Sprintf("%d", c.proxy.counter.OldSlowLogTotal)})

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
		"CacheConns",
		"PushConnCount",
		"PopConnCount",
	}
	var rows [][]string
	const (
		Column = 10
	)

	//var nodeRows [][]string
	for name, node := range c.schema.nodes {
		//"master"
		idleConns,cacheConns,pushConnCount,popConnCount := node.Master.ConnCount()
		
		rows = append(
			rows,
			[]string{
				name,
				node.Master.Addr(),
				"master",
				node.Master.State(),
				fmt.Sprintf("%v", time.Unix(node.Master.GetLastPing(), 0)),
				strconv.Itoa(node.Cfg.MaxConnNum),
				strconv.Itoa(idleConns),
				strconv.Itoa(cacheConns),
				strconv.FormatInt(pushConnCount, 10),
				strconv.FormatInt(popConnCount, 10),
			})
		//"slave"
		for _, slave := range node.Slave {
			if slave != nil {
				idleConns,cacheConns,pushConnCount,popConnCount := slave.ConnCount()

				rows = append(
					rows,
					[]string{
						name,
						slave.Addr(),
						"slave",
						slave.State(),
						fmt.Sprintf("%v", time.Unix(slave.GetLastPing(), 0)),
						strconv.Itoa(node.Cfg.MaxConnNum),
						strconv.Itoa(idleConns),
						strconv.Itoa(cacheConns),
						strconv.FormatInt(pushConnCount, 10),
						strconv.FormatInt(popConnCount, 10),
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
	var rows [][]string
	var names []string = []string{
		"User",
		"DB",
		"Table",
		"Type",
		"Key",
		"Nodes_List",
		"Locations",
		"TableRowLimit",
	}
	var Column = len(names)

	for _, schemaConfig := range c.proxy.cfg.SchemaList {
		//default Rule
		var defaultRule = c.schema.rule.DefaultRule
		if defaultRule != nil {
			rows = append(
				rows,
				[]string{
					schemaConfig.User,
					defaultRule.DB,
					defaultRule.Table,
					defaultRule.Type,
					defaultRule.Key,
					strings.Join(defaultRule.Nodes, ", "),
					"",
					"0",
				},
			)
		}

		shardRule := schemaConfig.ShardRule
		for _, r := range shardRule {
			rows = append(
				rows,
				[]string{
					schemaConfig.User,
					r.DB,
					r.Table,
					r.Type,
					r.Key,
					strings.Join(r.Nodes, ", "),
					hack.ArrayToString(r.Locations),
					strconv.Itoa(r.TableRowLimit),
				},
			)
		}
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

func (c *ClientConn) handleShowAllowIPConfig() (*mysql.Resultset, error) {
	var Column = 1
	var rows [][]string
	var names []string = []string{
		"AllowIP",
	}

	//allow ips
	current, _, _ := c.proxy.allowipsIndex.Get()
	var allowips = c.proxy.allowips[current]
	if len(allowips) != 0 {
		for _, v := range allowips {
			if v.Info() != "" {
				rows = append(rows,
					[]string{
						v.Info(),
					})
			}
		}
	}

	if len(rows) == 0 {
		rows = append(rows, []string{""})
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

func (c *ClientConn) handleShowProxyStatus() (*mysql.Resultset, error) {
	var Column = 1
	var rows [][]string
	var names []string = []string{
		"status",
	}

	var status string
	status = c.proxy.Status()
	rows = append(rows,
		[]string{
			status,
		})

	var values [][]interface{} = make([][]interface{}, len(rows))
	for i := range rows {
		values[i] = make([]interface{}, Column)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	return c.buildResultset(nil, names, values)
}

func (c *ClientConn) handleShowBlackSqlConfig() (*mysql.Resultset, error) {
	var Column = 1
	var rows [][]string
	var names []string = []string{
		"BlackListSql",
	}

	//black sql
	var blackListSqls = c.proxy.blacklistSqls[c.proxy.blacklistSqlsIndex].sqls
	if len(blackListSqls) != 0 {
		for _, v := range blackListSqls {
			rows = append(rows,
				[]string{
					v,
				})
		}
	}

	if len(rows) == 0 {
		rows = append(rows, []string{""})
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

func (c *ClientConn) handleChangeProxy(v string) error {
	return c.proxy.ChangeProxy(v)
}

func (c *ClientConn) handleChangeLogSql(v string) error {
	return c.proxy.ChangeLogSql(v)
}

func (c *ClientConn) handleChangeSlowLogTime(v string) error {
	return c.proxy.ChangeSlowLogTime(v)
}

func (c *ClientConn) handleAddAllowIP(v string) error {
	v = strings.TrimSpace(v)
	err := c.proxy.AddAllowIP(v)
	return err
}

func (c *ClientConn) handleDelAllowIP(v string) error {
	v = strings.TrimSpace(v)
	err := c.proxy.DelAllowIP(v)
	return err
}

func (c *ClientConn) handleAddBlackSql(v string) error {
	v = strings.TrimSpace(v)
	err := c.proxy.AddBlackSql(v)
	return err
}

func (c *ClientConn) handleDelBlackSql(v string) error {
	v = strings.TrimSpace(v)
	err := c.proxy.DelBlackSql(v)
	return err
}

func (c *ClientConn) handleAdminSave(k string, v string) error {
	if len(k) == 0 || len(v) == 0 {
		return errors.ErrCmdUnsupport
	}
	if k == ADMIN_PROXY && v == ADMIN_CONFIG {
		return c.proxy.SaveProxyConfig()
	}

	return errors.ErrCmdUnsupport
}
