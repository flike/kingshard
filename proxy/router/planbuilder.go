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

package router

import (
	"sort"
	"strconv"

	"strings"

	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/sqlparser"
)

const (
	EID_NODE = iota
	VALUE_NODE
	LIST_NODE
	OTHER_NODE
)

type Plan struct {
	Rule *Rule

	Criteria sqlparser.SQLNode
	KeyIndex int //used for insert/replace to find shard key idx
	//used for insert/replace values,key is table index,and value is
	//the rows for insert or replace.
	Rows map[int]sqlparser.Values

	SubTableValueGroups map[int]sqlparser.ValTuple //按照tableIndex存放ValueExpr
	InRightToReplace    *sqlparser.ComparisonExpr  //记录in的右边Expr,用来动态替换不同table in的值
	RouteTableIndexs    []int
	RouteNodeIndexs     []int
	RewrittenSqls       map[string][]string
}

func (plan *Plan) rewriteWhereIn(tableIndex int) (sqlparser.ValExpr, error) {
	var oldright sqlparser.ValExpr
	if plan.InRightToReplace != nil && plan.SubTableValueGroups[tableIndex] != nil {
		//assign corresponding values to different table index
		oldright = plan.InRightToReplace.Right
		plan.InRightToReplace.Right = plan.SubTableValueGroups[tableIndex]
	}
	return oldright, nil
}

func (plan *Plan) notList(l []int) []int {
	return differentList(plan.Rule.SubTableIndexs, l)
}

func (plan *Plan) getTableIndexs(expr sqlparser.BoolExpr) ([]int, error) {
	switch plan.Rule.Type {
	case HashRuleType:
		return plan.getHashShardTableIndex(expr)
	case RangeRuleType:
		return plan.getRangeShardTableIndex(expr)
	case DateYearRuleType, DateMonthRuleType, DateDayRuleType:
		return plan.getDateShardTableIndex(expr)
	default:
		return plan.Rule.SubTableIndexs, nil
	}
	return nil, nil
}

//Get the table index of hash shard type
func (plan *Plan) getHashShardTableIndex(expr sqlparser.BoolExpr) ([]int, error) {
	var index int
	var err error
	switch criteria := expr.(type) {
	case *sqlparser.ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>": //=对应的分片
			if plan.getValueType(criteria.Left) == EID_NODE {
				index, err = plan.getTableIndexByValue(criteria.Right)
			} else {
				index, err = plan.getTableIndexByValue(criteria.Left)
			}
			if err != nil {
				return nil, err
			}
			return []int{index}, nil
		case "<", "<=", ">", ">=", "not in":
			return plan.Rule.SubTableIndexs, nil
		case "in":
			return plan.getTableIndexsByTuple(criteria.Right)
		}
	case *sqlparser.RangeCond: //between ... and ...
		return plan.Rule.SubTableIndexs, nil
	default:
		return plan.Rule.SubTableIndexs, nil
	}

	return plan.RouteTableIndexs, nil
}

//Get the table index of range shard type
func (plan *Plan) getRangeShardTableIndex(expr sqlparser.BoolExpr) ([]int, error) {
	var index int
	var err error
	switch criteria := expr.(type) {
	case *sqlparser.ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>": //=对应的分片
			if plan.getValueType(criteria.Left) == EID_NODE {
				index, err = plan.getTableIndexByValue(criteria.Right)
			} else {
				index, err = plan.getTableIndexByValue(criteria.Left)
			}
			if err != nil {
				return nil, err
			}
			return []int{index}, nil
		case "<", "<=":
			if plan.getValueType(criteria.Left) == EID_NODE {
				index, err = plan.getTableIndexByValue(criteria.Right)
				if err != nil {
					return nil, err
				}
				if criteria.Operator == "<" {
					//调整边界值，当shard[index].start等于criteria.Right 则index--
					index = plan.adjustShardIndex(criteria.Right, index)
				}

				return makeList(0, index+1), nil
			} else {
				index, err = plan.getTableIndexByValue(criteria.Left)
				if err != nil {
					return nil, err
				}
				return makeList(index, len(plan.Rule.SubTableIndexs)), nil
			}
		case ">", ">=":
			if plan.getValueType(criteria.Left) == EID_NODE {
				index, err = plan.getTableIndexByValue(criteria.Right)
				if err != nil {
					return nil, err
				}
				return makeList(index, len(plan.Rule.SubTableIndexs)), nil
			} else { // 10 > id，这种情况
				index, err = plan.getTableIndexByValue(criteria.Left)
				if err != nil {
					return nil, err
				}
				if criteria.Operator == ">" {
					index = plan.adjustShardIndex(criteria.Left, index)
				}
				return makeList(0, index+1), nil
			}
		case "in":
			return plan.getTableIndexsByTuple(criteria.Right)
		case "not in":
			return plan.Rule.SubTableIndexs, nil
		}
	case *sqlparser.RangeCond:
		var start, last int
		start, err = plan.getTableIndexByValue(criteria.From)
		if err != nil {
			return nil, err
		}
		last, err = plan.getTableIndexByValue(criteria.To)
		if err != nil {
			return nil, err
		}
		if criteria.Operator == "between" { //对应between ...and ...
			if last < start {
				start, last = last, start
			}
			return makeList(start, last+1), nil
		} else { //对应not between ....and
			if last < start {
				start, last = last, start
				start = plan.adjustShardIndex(criteria.To, start)
			} else {
				start = plan.adjustShardIndex(criteria.From, start)
			}

			l1 := makeList(0, start+1)
			l2 := makeList(last, len(plan.Rule.SubTableIndexs))
			return unionList(l1, l2), nil
		}
	default:
		return plan.Rule.SubTableIndexs, nil
	}

	return plan.RouteTableIndexs, nil
}

//Get the table index of date shard type(date_year,date_month,date_day).
func (plan *Plan) getDateShardTableIndex(expr sqlparser.BoolExpr) ([]int, error) {
	var index int
	var err error
	switch criteria := expr.(type) {
	case *sqlparser.ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>": //=对应的分片
			if plan.getValueType(criteria.Left) == EID_NODE {
				index, err = plan.getTableIndexByValue(criteria.Right)
			} else {
				index, err = plan.getTableIndexByValue(criteria.Left)
			}
			if err != nil {
				return nil, err
			}
			return []int{index}, nil
		case "<", "<=":
			if plan.getValueType(criteria.Left) == EID_NODE {
				index, err = plan.getTableIndexByValue(criteria.Right)
				if err != nil {
					return nil, err
				}
				return makeLeList(index, plan.Rule.SubTableIndexs), nil
			} else {
				index, err = plan.getTableIndexByValue(criteria.Left)
				if err != nil {
					return nil, err
				}
				return makeGeList(index, plan.Rule.SubTableIndexs), nil
			}
		case ">", ">=":
			if plan.getValueType(criteria.Left) == EID_NODE {
				index, err = plan.getTableIndexByValue(criteria.Right)
				if err != nil {
					return nil, err
				}
				return makeGeList(index, plan.Rule.SubTableIndexs), nil
			} else { // 10 > id，这种情况
				index, err = plan.getTableIndexByValue(criteria.Left)
				if err != nil {
					return nil, err
				}
				return makeLeList(index, plan.Rule.SubTableIndexs), nil
			}
		case "in":
			return plan.getTableIndexsByTuple(criteria.Right)
		case "not in":
			l, err := plan.getTableIndexsByTuple(criteria.Right)
			if err != nil {
				return nil, err
			}
			return plan.notList(l), nil
		}
	case *sqlparser.RangeCond:
		var start, last int
		start, err = plan.getTableIndexByValue(criteria.From)
		if err != nil {
			return nil, err
		}
		last, err = plan.getTableIndexByValue(criteria.To)
		if err != nil {
			return nil, err
		}
		if last < start {
			start, last = last, start
		}
		if criteria.Operator == "between" { //对应between ...and ...
			return makeBetweenList(start, last, plan.Rule.SubTableIndexs), nil
		} else { //对应not between ....and
			l := makeBetweenList(start, last, plan.Rule.SubTableIndexs)
			return plan.notList(l), nil
		}
	default:
		return plan.Rule.SubTableIndexs, nil
	}

	return plan.RouteTableIndexs, nil
}

//计算表下标和node下标
func (plan *Plan) calRouteIndexs() error {
	var err error
	nodesCount := len(plan.Rule.Nodes)

	if plan.Rule.Type == DefaultRuleType {
		plan.RouteNodeIndexs = []int{0}
		return nil
	}
	if plan.Criteria == nil { //如果没有分表条件，则是全子表扫描
		if plan.Rule.Type != DefaultRuleType {
			golog.Error("Plan", "calRouteIndexs", "plan have no criteria", 0,
				"type", plan.Rule.Type)
			return errors.ErrNoCriteria
		}
	}

	switch criteria := plan.Criteria.(type) {
	case sqlparser.Values: //代表insert中values
		plan.RouteTableIndexs, err = plan.getInsertTableIndex(criteria)
		if err != nil {
			return err
		}
		plan.RouteNodeIndexs = plan.TindexsToNindexs(plan.RouteTableIndexs)
		return nil
	case sqlparser.BoolExpr:
		plan.RouteTableIndexs, err = plan.getTableIndexByBoolExpr(criteria)
		if err != nil {
			return err
		}
		plan.RouteNodeIndexs = plan.TindexsToNindexs(plan.RouteTableIndexs)
		return nil
	default:
		plan.RouteTableIndexs = plan.Rule.SubTableIndexs
		plan.RouteNodeIndexs = makeList(0, nodesCount)
		return nil
	}
}

func (plan *Plan) checkValuesType(vals sqlparser.Values) (sqlparser.Values, error) {
	// Analyze first value of every item in the list
	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case sqlparser.ValTuple:
			result := plan.getValueType(tuple[0])
			if result != VALUE_NODE {
				return nil, errors.ErrInsertTooComplex
			}
		default:
			//panic(sqlparser.NewParserError("insert is too complex"))
			return nil, errors.ErrInsertTooComplex
		}
	}
	return vals, nil
}

/*返回valExpr表达式对应的类型*/
func (plan *Plan) getValueType(valExpr sqlparser.ValExpr) int {
	switch node := valExpr.(type) {
	case *sqlparser.ColName:
		//remove table name
		if string(node.Qualifier) == plan.Rule.Table {
			node.Qualifier = nil
		}
		if strings.ToLower(string(node.Name)) == plan.Rule.Key {
			return EID_NODE //表示这是分片id对应的node
		}
	case sqlparser.ValTuple:
		for _, n := range node {
			if plan.getValueType(n) != VALUE_NODE {
				return OTHER_NODE
			}
		}
		return LIST_NODE //列表节点
	case sqlparser.StrVal, sqlparser.NumVal, sqlparser.ValArg: //普通的值节点，字符串值，绑定变量参数
		return VALUE_NODE
	}
	return OTHER_NODE
}

func (plan *Plan) getTableIndexByBoolExpr(node sqlparser.BoolExpr) ([]int, error) {
	switch node := node.(type) {
	case *sqlparser.AndExpr:
		left, err := plan.getTableIndexByBoolExpr(node.Left)
		if err != nil {
			return nil, err
		}
		right, err := plan.getTableIndexByBoolExpr(node.Right)
		if err != nil {
			return nil, err
		}
		return interList(left, right), nil
	case *sqlparser.OrExpr:
		left, err := plan.getTableIndexByBoolExpr(node.Left)
		if err != nil {
			return nil, err
		}
		right, err := plan.getTableIndexByBoolExpr(node.Right)
		if err != nil {
			return nil, err
		}
		return unionList(left, right), nil
	case *sqlparser.ParenBoolExpr: //加上括号的BoolExpr，node.Expr去掉了括号
		return plan.getTableIndexByBoolExpr(node.Expr)
	case *sqlparser.ComparisonExpr:
		switch {
		case sqlparser.StringIn(node.Operator, "=", "<", ">", "<=", ">=", "<=>"):
			left := plan.getValueType(node.Left)
			right := plan.getValueType(node.Right)
			if (left == EID_NODE && right == VALUE_NODE) || (left == VALUE_NODE && right == EID_NODE) {
				return plan.getTableIndexs(node)
			}
		case sqlparser.StringIn(node.Operator, "in", "not in"):
			left := plan.getValueType(node.Left)
			right := plan.getValueType(node.Right)
			if left == EID_NODE && right == LIST_NODE {
				if strings.EqualFold(node.Operator, "in") { //only deal with in expr, it's impossible to process not in here.
					plan.InRightToReplace = node
				}
				return plan.getTableIndexs(node)
			}
		}
	case *sqlparser.RangeCond:
		left := plan.getValueType(node.Left)
		from := plan.getValueType(node.From)
		to := plan.getValueType(node.To)
		if left == EID_NODE && from == VALUE_NODE && to == VALUE_NODE {
			return plan.getTableIndexs(node)
		}
	}
	return plan.Rule.SubTableIndexs, nil
}

//获得(12,14,23)对应的table index
func (plan *Plan) getTableIndexsByTuple(valExpr sqlparser.ValExpr) ([]int, error) {
	shardset := make(map[int]sqlparser.ValTuple)
	switch node := valExpr.(type) {
	case sqlparser.ValTuple:
		for _, n := range node {
			//n.Format()
			index, err := plan.getTableIndexByValue(n)

			if err != nil {
				return nil, err
			}
			valExprs := shardset[index]

			if valExprs == nil {
				valExprs = make([]sqlparser.ValExpr, 0)
			}
			valExprs = append(valExprs, n)
			shardset[index] = valExprs
		}
	}
	plan.SubTableValueGroups = shardset
	shardlist := make([]int, len(shardset))
	index := 0
	for k := range shardset {
		shardlist[index] = k
		index++
	}

	sort.Ints(shardlist)
	return shardlist, nil
}

//get the insert table index and set plan.Rows
func (plan *Plan) getInsertTableIndex(vals sqlparser.Values) ([]int, error) {
	tableIndexs := make([]int, 0, len(vals))
	rowsToTindex := make(map[int][]sqlparser.Tuple)
	for i := 0; i < len(vals); i++ {
		valueExpression := vals[i].(sqlparser.ValTuple)
		if len(valueExpression) < (plan.KeyIndex + 1) {
			return nil, errors.ErrColsLenNotMatch
		}

		tableIndex, err := plan.getTableIndexByValue(valueExpression[plan.KeyIndex])
		if err != nil {
			return nil, err
		}

		tableIndexs = append(tableIndexs, tableIndex)
		//get the rows insert into this table
		rowsToTindex[tableIndex] = append(rowsToTindex[tableIndex], valueExpression)
	}
	for k, v := range rowsToTindex {
		plan.Rows[k] = (sqlparser.Values)(v)
	}

	return cleanList(tableIndexs), nil
}

// find shard key index in insert or replace SQL
// plan.Rule cols must not nil
func (plan *Plan) GetIRKeyIndex(cols sqlparser.Columns) error {
	if plan.Rule == nil {
		return errors.ErrNoPlanRule
	}
	plan.KeyIndex = -1
	for i, _ := range cols {
		colname := string(cols[i].(*sqlparser.NonStarExpr).Expr.(*sqlparser.ColName).Name)

		if strings.ToLower(colname) == plan.Rule.Key {
			plan.KeyIndex = i
			break
		}
	}
	if plan.KeyIndex == -1 {
		return errors.ErrIRNoShardingKey
	}
	return nil
}

func (plan *Plan) getTableIndexByValue(valExpr sqlparser.ValExpr) (int, error) {
	value := plan.getBoundValue(valExpr)
	return plan.Rule.FindTableIndex(value)
}

func (plan *Plan) adjustShardIndex(valExpr sqlparser.ValExpr, index int) int {
	value := plan.getBoundValue(valExpr)
	//生成一个范围的接口,[100,120)
	s, ok := plan.Rule.Shard.(RangeShard)
	if !ok {
		return index
	}
	//value是否和shard[index].Start相等
	if s.EqualStart(value, index) {
		index--
		if index < 0 {
			panic(sqlparser.NewParserError("invalid range sharding"))
		}
	}
	return index
}

/*获得valExpr对应的值*/
func (plan *Plan) getBoundValue(valExpr sqlparser.ValExpr) interface{} {
	switch node := valExpr.(type) {
	case sqlparser.ValTuple: //ValTuple可以是一个slice
		if len(node) != 1 {
			panic(sqlparser.NewParserError("tuples not allowed as insert values"))
		}
		// TODO: Change parser to create single value tuples into non-tuples.
		return plan.getBoundValue(node[0])
	case sqlparser.StrVal:
		return string(node)
	case sqlparser.NumVal:
		val, err := strconv.ParseInt(string(node), 10, 64)
		if err != nil {
			panic(sqlparser.NewParserError("%s", err.Error()))
		}
		return val
	case sqlparser.ValArg:
		panic("Unexpected token")
	}
	panic("Unexpected token")
}

/*2,5 ==> [2,3,4]*/
func makeList(start, end int) []int {
	list := make([]int, end-start)
	for i := start; i < end; i++ {
		list[i-start] = i
	}
	return list
}

//if value is 2016, and indexs is [2015,2016,2017]
//the result is [2015,2016]
func makeLeList(value int, indexs []int) []int {
	sort.Ints(indexs)
	for k, v := range indexs {
		if v == value {
			return indexs[:k+1]
		}
	}
	return nil
}

//if value is 2016, and indexs is [2015,2016,2017,2018]
//the result is [2016,2017,2018]
func makeGeList(value int, indexs []int) []int {
	sort.Ints(indexs)
	for k, v := range indexs {
		if v == value {
			return indexs[k:]
		}
	}
	return nil
}

//if value is 2016, and indexs is [2015,2016,2017,2018]
//the result is [2015]
func makeLtList(value int, indexs []int) []int {
	sort.Ints(indexs)
	for k, v := range indexs {
		if v == value {
			return indexs[:k]
		}
	}
	return nil
}

//if value is 2016, and indexs is [2015,2016,2017,2018]
//the result is [2017,2018]
func makeGtList(value int, indexs []int) []int {
	sort.Ints(indexs)
	for k, v := range indexs {
		if v == value {
			return indexs[k+1:]
		}
	}
	return nil
}

//if start is 2016, end is 2017. indexs is [2015,2016,2017,2018]
//the result is [2016,2017]
func makeBetweenList(start, end int, indexs []int) []int {
	var startIndex, endIndex int
	var SetStart bool
	if end < start {
		start, end = end, start
	}
	sort.Ints(indexs)
	for k, v := range indexs {
		if v == start {
			startIndex = k
			SetStart = true
		}
		if v == end {
			endIndex = k
			if SetStart {
				return indexs[startIndex : endIndex+1]
			}
		}
	}
	return nil
}

// l1 & l2
func interList(l1 []int, l2 []int) []int {
	if len(l1) == 0 || len(l2) == 0 {
		return []int{}
	}

	l3 := make([]int, 0, len(l1)+len(l2))
	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] == l2[j] {
			l3 = append(l3, l1[i])
			i++
			j++
		} else if l1[i] < l2[j] {
			i++
		} else {
			j++
		}
	}

	return l3
}

// l1 | l2
func unionList(l1 []int, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]int, 0, len(l1)+len(l2))

	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l3 = append(l3, l1[i])
			i++
		} else if l1[i] > l2[j] {
			l3 = append(l3, l2[j])
			j++
		} else {
			l3 = append(l3, l1[i])
			i++
			j++
		}
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	} else if j != len(l2) {
		l3 = append(l3, l2[j:]...)
	}

	return l3
}

// l1 - l2
func differentList(l1 []int, l2 []int) []int {
	if len(l1) == 0 {
		return []int{}
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]int, 0, len(l1))

	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l3 = append(l3, l1[i])
			i++
		} else if l1[i] > l2[j] {
			j++
		} else {
			i++
			j++
		}
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	}

	return l3
}

func cleanList(l []int) []int {
	s := make(map[int]struct{})
	listLen := len(l)
	l2 := make([]int, 0, listLen)

	for i := 0; i < listLen; i++ {
		k := l[i]
		s[k] = struct{}{}
	}
	for k := range s {
		l2 = append(l2, k)
	}
	return l2
}

func (plan *Plan) TindexsToNindexs(tableIndexs []int) []int {
	count := len(tableIndexs)
	nodeIndes := make([]int, 0, count)
	for i := 0; i < count; i++ {
		tx := tableIndexs[i]
		nodeIndes = append(nodeIndes, plan.Rule.TableToNode[tx])
	}

	return cleanList(nodeIndes)
}
