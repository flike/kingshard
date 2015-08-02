package server

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/core/hack"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
)

func (c *ClientConn) handleShow(sql string, stmt *sqlparser.Show) error {
	var err error
	var r *mysql.Resultset
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

func (c *ClientConn) handleShowDatabases() (*mysql.Resultset, error) {
	dbs := make([]interface{}, 0, len(c.proxy.schemas))
	for key := range c.proxy.schemas {
		dbs = append(dbs, key)
	}

	return c.buildSimpleShowResultset(dbs, "Database")
}

func (c *ClientConn) handleShowTables(sql string, stmt *sqlparser.Show) (*mysql.Resultset, error) {
	s := c.schema
	if stmt.From != nil {
		db := nstring(stmt.From)
		s = c.proxy.GetSchema(db)
	}

	if s == nil {
		return nil, mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
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

func (c *ClientConn) handleShowProxy(sql string, stmt *sqlparser.Show) (*mysql.Resultset, error) {
	var err error
	var r *mysql.Resultset
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

func (c *ClientConn) handleShowProxyStatus(sql string, stmt *sqlparser.Show) (*mysql.Resultset, error) {
	// TODO: handle like_or_where expr
	return nil, nil
}

func (c *ClientConn) buildSimpleShowResultset(values []interface{}, name string) (*mysql.Resultset, error) {

	r := new(mysql.Resultset)

	field := &mysql.Field{}

	field.Name = hack.Slice(name)
	field.Charset = 33
	field.Type = mysql.MYSQL_TYPE_VAR_STRING

	r.Fields = []*mysql.Field{field}

	var row []byte
	var err error

	for _, value := range values {
		row, err = formatValue(value)
		if err != nil {
			return nil, err
		}
		r.RowDatas = append(r.RowDatas,
			mysql.PutLengthEncodedString(row))
	}

	return r, nil
}
