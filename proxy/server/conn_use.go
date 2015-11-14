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
	defer c.closeConn(co, err != nil)
	if err != nil {
		return err
	}

	if err = co.UseDB(string(stmt.DB)); err != nil {
		return err
	}
	c.db = string(stmt.DB)
	return c.writeOK(nil)
}
