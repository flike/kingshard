package server

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
)

func (c *ClientConn) handleSimpleSelect(sql string, stmt *sqlparser.SimpleSelect) error {
	if len(stmt.SelectExprs) != 1 {
		return fmt.Errorf("support select one informaction function, %s", sql)
	}

	expr, ok := stmt.SelectExprs[0].(*sqlparser.NonStarExpr)
	if !ok {
		return fmt.Errorf("support select informaction function, %s", sql)
	}

	var f *sqlparser.FuncExpr
	f, ok = expr.Expr.(*sqlparser.FuncExpr)
	if !ok {
		return fmt.Errorf("support select informaction function, %s", sql)
	}

	var r *mysql.Resultset
	var err error

	switch strings.ToLower(string(f.Name)) {
	case "last_insert_id":
		r, err = c.buildSimpleSelectResult(c.lastInsertId, f.Name, expr.As)
	case "row_count":
		r, err = c.buildSimpleSelectResult(c.affectedRows, f.Name, expr.As)
	case "version":
		r, err = c.buildSimpleSelectResult(mysql.ServerVersion, f.Name, expr.As)
	case "connection_id":
		r, err = c.buildSimpleSelectResult(c.connectionId, f.Name, expr.As)
	case "database":
		if c.schema != nil {
			r, err = c.buildSimpleSelectResult(c.schema.db, f.Name, expr.As)
		} else {
			r, err = c.buildSimpleSelectResult("NULL", f.Name, expr.As)
		}
	default:
		return fmt.Errorf("function %s not support", f.Name)
	}

	if err != nil {
		return err
	}

	return c.writeResultset(c.status, r)
}

func (c *ClientConn) buildSimpleSelectResult(value interface{}, name []byte, asName []byte) (*mysql.Resultset, error) {
	field := &mysql.Field{}

	field.Name = name

	if asName != nil {
		field.Name = asName
	}

	field.OrgName = name

	formatField(field, value)

	r := &mysql.Resultset{Fields: []*mysql.Field{field}}
	row, err := formatValue(value)
	if err != nil {
		return nil, err
	}
	r.RowDatas = append(r.RowDatas, mysql.PutLengthEncodedString(row))

	return r, nil
}

func (c *ClientConn) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := string(data[0:index])
	wildcard := string(data[index+1:])

	if c.schema == nil {
		return mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
	}

	nodeName := c.schema.rule.GetRule(table).Nodes[0]

	n := c.proxy.GetNode(nodeName)

	co, err := n.GetMasterConn()
	if err != nil {
		return err
	}
	defer co.Close()

	if err = co.UseDB(c.db); err != nil {
		return err
	}

	if fs, err := co.FieldList(table, wildcard); err != nil {
		return err
	} else {
		return c.writeFieldList(c.status, fs)
	}
}

func (c *ClientConn) writeFieldList(status uint16, fs []*mysql.Field) error {
	c.affectedRows = int64(-1)

	data := make([]byte, 4, 1024)

	for _, v := range fs {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := c.writePacket(data); err != nil {
			return err
		}
	}

	if err := c.writeEOF(status); err != nil {
		return err
	}
	return nil
}
