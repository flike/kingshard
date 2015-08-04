package server

import (
	"fmt"
	. "github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
	"strings"
	"time"
)

const (
	Master = "master"
	Slave  = "slave"
	Proxy  = "proxy"
	Config = "config"

	ServerRegion = "server"
	NodeRegion   = "node"

	ADMIN_OPT_ADD  = "add"
	ADMIN_OPT_DEL  = "del"
	ADMIN_OPT_UP   = "up"
	ADMIN_OPT_DOWN = "down"
	ADMIN_OPT_SHOW = "show"
)

var cmdServerOrder = []string{"opt", "k", "v"}
var cmdNodeOrder = []string{"opt", "node", "k", "v"}

func (c *ClientConn) handleNodeCmd(rows sqlparser.InsertRows) error {
	var err error
	var opt, nodeName, role, addr string

	vals := rows.(sqlparser.Values)
	if len(vals) == 0 {
		return ErrCmdUnsupport
	}

	tuple := vals[0].(sqlparser.ValTuple)
	if len(tuple) != len(cmdNodeOrder) {
		return ErrCmdUnsupport
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
		err = ErrCmdUnsupport
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
		return nil, ErrCmdUnsupport
	}

	tuple := vals[0].(sqlparser.ValTuple)
	if len(tuple) != len(cmdServerOrder) {
		return nil, ErrCmdUnsupport
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
		err = ErrCmdUnsupport
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
		return ErrCmdUnsupport
	}

	return c.proxy.AddSlave(nodeName, addr)
}

func (c *ClientConn) DeleteDatabase(nodeName string, role string, addr string) error {
	//can not delete a master database
	if role != Slave {
		return ErrCmdUnsupport
	}

	return c.proxy.DeleteSlave(nodeName, addr)
}

func (c *ClientConn) UpDatabase(nodeName string, role string, addr string) error {
	if role != Master && role != Slave {
		return ErrCmdUnsupport
	}
	if role == Master {
		return c.proxy.UpMaster(nodeName, addr)
	}

	return c.proxy.UpSlave(nodeName, addr)
}

func (c *ClientConn) DownDatabase(nodeName string, role string, addr string) error {
	if role != Master && role != Slave {
		return ErrCmdUnsupport
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
		return ErrCmdUnsupport
	}

	for i := 0; i < len(node); i++ {
		val := sqlparser.String(node[i])
		if val != cmdOrder[i] {
			return ErrCmdUnsupport
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
		return nil, ErrCmdUnsupport
	}
	if k == Proxy && v == Config {
		return c.handleShowProxyConfig()
	}
	return nil, ErrCmdUnsupport
}

func (c *ClientConn) handleShowProxyConfig() (*mysql.Resultset, error) {
	var names []string = []string{"Section", "Key", "Value"}
	var rows [][]string
	const (
		Column = 3
	)

	rows = append(rows, []string{"Global_Config", "Addr", c.proxy.cfg.Addr})
	rows = append(rows, []string{"Global_Config", "User", c.proxy.cfg.User})
	rows = append(rows, []string{"Global_Config", "Password", c.proxy.cfg.Password})
	rows = append(rows, []string{"Global_Config", "LogLevel", c.proxy.cfg.LogLevel})
	rows = append(rows, []string{"Global_Config", "Schemas_Count", fmt.Sprintf("%d", len(c.proxy.schemas))})
	rows = append(rows, []string{"Global_Config", "Nodes_Count", fmt.Sprintf("%d", len(c.proxy.nodes))})

	for db, schema := range c.proxy.schemas {
		rows = append(rows, []string{"Schemas", "DB", db})

		var nodeNames []string
		var nodeRows [][]string
		for name, node := range schema.nodes {
			nodeNames = append(nodeNames, name)
			var nodeSection = fmt.Sprintf("Schemas[%s]-Node[ %v ]", db, name)

			if node.Master != nil {
				nodeRows = append(nodeRows, []string{nodeSection, "Master", node.Master.String()})
			}

			if node.Slave != nil {
				nodeRows = append(nodeRows, []string{nodeSection, "Slave", node.FormatSlave()})
			}
			nodeRows = append(nodeRows, []string{nodeSection, "Last_Master_Ping", fmt.Sprintf("%v", time.Unix(node.LastMasterPing, 0))})

			nodeRows = append(nodeRows, []string{nodeSection, "Last_Slave_Ping", fmt.Sprintf("%v", time.Unix(node.LastSlavePing, 0))})

			nodeRows = append(nodeRows, []string{nodeSection, "down_after_noalive", fmt.Sprintf("%v", node.DownAfterNoAlive)})

		}
		rows = append(rows, []string{fmt.Sprintf("Schemas[%s]", db), "Nodes_List", strings.Join(nodeNames, ",")})

		var defaultRule = schema.rule.DefaultRule
		if defaultRule.DB == db {
			if defaultRule.DB == db {
				rows = append(rows, []string{fmt.Sprintf("Schemas[%s]_Rule_Default", db),
					"Default_Table", defaultRule.String()})
			}
		}
		for tb, r := range schema.rule.Rules {
			if r.DB == db {
				rows = append(rows, []string{fmt.Sprintf("Schemas[%s]_Rule_Table", db),
					fmt.Sprintf("Table[ %s ]", tb), r.String()})
			}
		}

		rows = append(rows, nodeRows...)

	}

	var values [][]interface{} = make([][]interface{}, len(rows))
	for i := range rows {
		values[i] = make([]interface{}, Column)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	return c.buildResultset(names, values)
}
