package server

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/core/hack"
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

/////

/*处理select语句*/
func (c *ClientConn) handleSelect(stmt *sqlparser.Select, args []interface{}) error {
	var fromSlave bool = true
	plan, err := c.schema.rule.BuildPlan(stmt)
	if err != nil {
		return err
	}
	if 0 < len(stmt.Comments) {
		comment := string(stmt.Comments[0])
		if 0 < len(comment) && strings.ToLower(comment) == MasterComment {
			fromSlave = false
		}
	}

	conns, err := c.getShardConns(fromSlave, plan)
	if err != nil {
		golog.Error("ClientConn", "handleSelect", err.Error(), c.connectionId)
		return err
	}
	if conns == nil {
		r := c.newEmptyResultset(stmt)
		return c.writeResultset(c.status, r)
	}

	var rs []*mysql.Result
	rs, err = c.executeInMultiNodes(conns, plan.RewrittenSqls, args)
	c.closeShardConns(conns, false)
	if err != nil {
		golog.Error("ClientConn", "handleSelect", err.Error(), c.connectionId)
		return err
	}

	err = c.mergeSelectResult(rs, stmt)
	if err != nil {
		golog.Error("ClientConn", "handleSelect", err.Error(), c.connectionId)
	}

	return err
}

func (c *ClientConn) mergeSelectResult(rs []*mysql.Result, stmt *sqlparser.Select) error {
	var err error
	r := rs[0].Resultset
	status := c.status | rs[0].Status

	funcExprs := c.getFuncExprs(stmt)
	if len(funcExprs) == 0 {
		for i := 1; i < len(rs); i++ {
			status |= rs[i].Status
			for j := range rs[i].Values {
				r.Values = append(r.Values, rs[i].Values[j])
				r.RowDatas = append(r.RowDatas, rs[i].RowDatas[j])
			}
		}
	} else {
		r, err = c.buildFuncExprResult(stmt, rs, funcExprs, &status)
		if err != nil {
			return err
		}
	}

	//to do order by, group by, limit offset
	c.sortSelectResult(r, stmt)
	//to do, add log here, sort may error because order by key not exist in resultset fields

	if err := c.limitSelectResult(r, stmt); err != nil {
		return err
	}

	return c.writeResultset(status, r)
}

func (c *ClientConn) sortSelectResult(r *mysql.Resultset, stmt *sqlparser.Select) error {
	if stmt.OrderBy == nil {
		return nil
	}

	sk := make([]mysql.SortKey, len(stmt.OrderBy))

	for i, o := range stmt.OrderBy {
		sk[i].Name = nstring(o.Expr)
		sk[i].Direction = o.Direction
	}

	return r.Sort(sk)
}

func (c *ClientConn) limitSelectResult(r *mysql.Resultset, stmt *sqlparser.Select) error {
	if stmt.Limit == nil {
		return nil
	}

	var offset, count int64
	var err error
	if stmt.Limit.Offset == nil {
		offset = 0
	} else {
		if o, ok := stmt.Limit.Offset.(sqlparser.NumVal); !ok {
			return fmt.Errorf("invalid select limit %s", nstring(stmt.Limit))
		} else {
			if offset, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
				return err
			}
		}
	}

	if o, ok := stmt.Limit.Rowcount.(sqlparser.NumVal); !ok {
		return fmt.Errorf("invalid limit %s", nstring(stmt.Limit))
	} else {
		if count, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
			return err
		} else if count < 0 {
			return fmt.Errorf("invalid limit %s", nstring(stmt.Limit))
		}
	}

	if offset+count > int64(len(r.Values)) {
		count = int64(len(r.Values)) - offset
	}

	r.Values = r.Values[offset : offset+count]
	r.RowDatas = r.RowDatas[offset : offset+count]

	return nil
}

func (c *ClientConn) buildFuncExprResult(stmt *sqlparser.Select,
	rs []*mysql.Result, funcExprs map[int]string, status *uint16) (*mysql.Resultset, error) {

	var names []string
	var err error
	r := rs[0].Resultset
	funcExprValues := make(map[int]interface{})

	for index, funcName := range funcExprs {
		funcExprValue, err := c.calFuncExprValue(
			funcName,
			rs,
			index,
		)
		if err != nil {
			return nil, err
		}
		funcExprValues[index] = funcExprValue
	}

	r.Values, err = c.buildFuncExprValues(rs, funcExprs, funcExprValues, status)

	if 0 < len(r.Values) {
		for _, field := range rs[0].Fields {
			names = append(names, string(field.Name))
		}
		r, err = c.buildResultset(names, r.Values)
		if err != nil {
			return nil, err
		}
	} else {
		r = c.newEmptyResultset(stmt)
	}

	return r, nil
}

//get the index of funcExpr
func (c *ClientConn) getFuncExprs(stmt *sqlparser.Select) map[int]string {
	var f *sqlparser.FuncExpr
	funcExprs := make(map[int]string)

	for i, expr := range stmt.SelectExprs {
		nonStarExpr, ok := expr.(*sqlparser.NonStarExpr)
		if !ok {
			continue
		}

		f, ok = nonStarExpr.Expr.(*sqlparser.FuncExpr)
		if !ok {
			continue
		} else {
			funcExprs[i] = string(f.Name)
		}
	}
	return funcExprs
}

//calculate the the value funcExpr(sum or count)
func (c *ClientConn) calFuncExprValue(funcName string,
	rs []*mysql.Result, index int) (interface{}, error) {

	var num int64
	switch strings.ToLower(funcName) {
	case SumFuncName, CountFuncName:
		for _, r := range rs {
			for k := range r.Values {
				result, err := r.GetInt(k, index)
				if err != nil {
					return nil, err
				}
				num += result
			}
		}
		return num, nil
	default:
		//get a non-null value of funcExpr
		for _, r := range rs {
			for k := range r.Values {
				result, err := r.GetValue(k, index)
				if err != nil {
					return nil, err
				}
				if result != nil {
					return result, nil
				}
			}
		}
	}

	return nil, nil
}

//build values of resultset
func (c *ClientConn) buildFuncExprValues(
	rs []*mysql.Result,
	funcExprs map[int]string,
	funcExprValues map[int]interface{},
	status *uint16,
) ([][]interface{}, error) {
	var rowEmpty bool

	values := make([][]interface{}, 0, len(rs))
	for i := 0; i < len(rs); i++ {
		*status |= rs[i].Status
		//iterate every row in the resultset(rs[i])
		for j := range rs[i].Values {
			//ignore the empty row
			rowEmpty = true
			for k := range funcExprs {
				rs[i].Values[j][k] = nil
			}
			for valueIndex := range rs[i].Values[j] {
				if rs[i].Values[j][valueIndex] != nil {
					rowEmpty = false
					break
				}
			}
			if rowEmpty {
				continue
			}
			//set fucExpr value for every non-empty row
			for k := range funcExprs {
				rs[i].Values[j][k] = funcExprValues[k]
			}
			values = append(values, rs[i].Values[j])
		}
	}

	//generate one row just for sum or count
	if len(values) == 0 {
		value := make([]interface{}, len(rs[0].Fields))
		for k := range funcExprs {
			value[k] = funcExprValues[k]
		}
		values = append(values, value)
	}

	return values, nil
}
