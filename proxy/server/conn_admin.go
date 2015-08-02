package server

import (
	"fmt"
	. "github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/sqlparser"
	"strings"
)

const (
	Master       = "master"
	Slave        = "slave"
	ServerRegion = "server"
	NodeRegion   = "node"

	ADMIN_OPT_ADD  = "add"
	ADMIN_OPT_DEL  = "del"
	ADMIN_OPT_UP   = "up"
	ADMIN_OPT_DOWN = "down"
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
	region := sqlparser.String(admin.Region)

	err = c.checkCmdOrder(region, admin.Columns)
	if err != nil {
		return err
	}

	switch strings.ToLower(region) {
	case NodeRegion:
		err = c.handleNodeCmd(admin.Rows)
	default:
		return fmt.Errorf("admin %s not supported now", region)
	}

	if err != nil {
		golog.Error("ClientConn", "handleAdmin", err.Error(),
			c.connectionId, "sql", sqlparser.String(admin))
		return err
	}

	return c.writeOK(nil)
}
