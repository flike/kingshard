package server

import (
	"fmt"
	"strings"

	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
)

var nstring = sqlparser.String

func (c *ClientConn) handleSet(stmt *sqlparser.Set, sql string) error {
	if len(stmt.Exprs) != 1 {
		return fmt.Errorf("must set one item once, not %s", nstring(stmt))
	}

	k := string(stmt.Exprs[0].Name.Name)

	switch strings.ToUpper(k) {
	case `AUTOCOMMIT`:
		return c.handleSetAutoCommit(stmt.Exprs[0].Expr)
	case `NAMES`, `CHARACTER_SET_RESULTS`, `CHARACTER_SET_CLIENT`, `CHARACTER_SET_CONNECTION`:
		return c.handleSetNames(stmt.Exprs[0].Expr)
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
	case '0':
		c.status &= ^mysql.SERVER_STATUS_AUTOCOMMIT
	default:
		return fmt.Errorf("invalid autocommit flag %s", value)
	}

	return c.writeOK(nil)
}

func (c *ClientConn) handleSetNames(val sqlparser.ValExpr) error {
	value := sqlparser.String(val)
	value = strings.Trim(value, "'`\"")

	charset := strings.ToLower(value)
	if charset == "null" {
		return c.writeOK(nil)
	}
	cid, ok := mysql.CharsetIds[charset]
	if !ok {
		return fmt.Errorf("invalid charset %s", charset)
	}

	c.charset = charset
	c.collation = cid

	return c.writeOK(nil)
}
