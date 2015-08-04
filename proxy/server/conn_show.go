package server

import (
	"fmt"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/core/hack"
	. "github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
	"sort"
	"strings"
)

func (c *ClientConn) handleShow(sql string, stmt *sqlparser.Show) error {
	var err error
	var r *Resultset
	switch strings.ToLower(stmt.Section) {
	case "databases":
		r, err = c.handleShowDatabases()
	case "tables":
		r, err = c.handleShowTables(sql, stmt)
	case "proxy":
		r, err = c.handleShowProxy(sql, stmt)
	default:
		err = fmt.Errorf("unsupport show %s now", sql)
	}

	if err != nil {
		return err
	}

	return c.writeResultset(c.status, r)
}

func (c *ClientConn) handleShowDatabases() (*Resultset, error) {
	dbs := make([]interface{}, 0, len(c.proxy.schemas))
	for key := range c.proxy.schemas {
		dbs = append(dbs, key)
	}

	return c.buildSimpleShowResultset(dbs, "Database")
}

func (c *ClientConn) handleShowTables(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	s := c.schema
	if stmt.From != nil {
		db := nstring(stmt.From)
		s = c.proxy.GetSchema(db)
	}

	if s == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	var tables []string
	tmap := map[string]struct{}{}
	for _, n := range s.nodes {
		co, err := n.GetMasterConn()
		if err != nil {
			return nil, err
		}

		if err := co.UseDB(s.db); err != nil {
			co.Close()
			return nil, err
		}

		if r, err := co.Execute(sql); err != nil {
			co.Close()
			return nil, err
		} else {
			co.Close()
			for i := 0; i < r.RowNumber(); i++ {
				n, _ := r.GetString(i, 0)
				if _, ok := tmap[n]; !ok {
					tables = append(tables, n)
				}
			}
		}
	}

	sort.Strings(tables)

	values := make([]interface{}, len(tables))
	for i := range tables {
		values[i] = tables[i]
	}

	return c.buildSimpleShowResultset(values, fmt.Sprintf("Tables_in_%s", s.db))
}

func (c *ClientConn) handleShowProxy(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	var err error
	var r *Resultset
	switch strings.ToLower(stmt.Key) {
	case "config":
		r, err = c.handleShowProxyConfig()
	case "status":
		r, err = c.handleShowProxyStatus(sql, stmt)
	default:
		err = fmt.Errorf("Unsupport show proxy [%v] yet, just support [config|status] now.", stmt.Key)
		golog.Warn("ClientConn", "handleShowProxy", err.Error(), 0)
		return nil, err
	}
	return r, err
}

func (c *ClientConn) handleShowProxyStatus(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	// TODO: handle like_or_where expr
	return nil, nil
}

func (c *ClientConn) buildSimpleShowResultset(values []interface{}, name string) (*Resultset, error) {

	r := new(Resultset)

	field := &Field{}

	field.Name = hack.Slice(name)
	field.Charset = 33
	field.Type = MYSQL_TYPE_VAR_STRING

	r.Fields = []*Field{field}

	var row []byte
	var err error

	for _, value := range values {
		row, err = formatValue(value)
		if err != nil {
			return nil, err
		}
		r.RowDatas = append(r.RowDatas,
			PutLengthEncodedString(row))
	}

	return r, nil
}
