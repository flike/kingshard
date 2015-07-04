package server

import (
	"fmt"
	"github.com/flike/kingshard/sqlparser"
	"strings"
)

const (
	Master = "master"
	Slave  = "slave"
)

func (c *ClientConn) handleAdmin(admin *sqlparser.Admin) error {
	name := string(admin.Name)

	var err error
	switch strings.ToLower(name) {
	case "upnode":
		err = c.adminUpNodeServer(admin.Values)
	case "downnode":
		err = c.adminDownNodeServer(admin.Values)
	default:
		return fmt.Errorf("admin %s not supported now", name)
	}

	if err != nil {
		return err
	}

	return c.writeOK(nil)
}

func (c *ClientConn) adminUpNodeServer(values sqlparser.ValExprs) error {
	if len(values) != 3 {
		return fmt.Errorf("upnode needs 3 args, not %d", len(values))
	}

	nodeName := nstring(values[0])
	sType := strings.ToLower(nstring(values[1]))
	addr := strings.ToLower(nstring(values[2]))

	switch sType {
	case Master:
		return c.proxy.UpMaster(nodeName, addr)
	case Slave:
		return c.proxy.UpSlave(nodeName, addr)
	default:
		return fmt.Errorf("invalid server type %s", sType)
	}
}

func (c *ClientConn) adminDownNodeServer(values sqlparser.ValExprs) error {
	if len(values) != 2 {
		return fmt.Errorf("upnode needs 2 args, not %d", len(values))
	}

	nodeName := nstring(values[0])
	sType := strings.ToLower(nstring(values[1]))

	switch sType {
	case Master:
		return c.proxy.DownMaster(nodeName)
	case Slave:
		return c.proxy.DownSlave(nodeName)
	default:
		return fmt.Errorf("invalid server type %s", sType)
	}
}
