package router

import (
	"sort"
	"strconv"

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

	TableIndexs      []int //value is table index
	RouteTableIndexs []int
	RouteNodeIndexs  []int
	RewrittenSqls    map[string][]string
}

func (plan *Plan) notList(l []int) []int {
	return differentList(plan.TableIndexs, l)
}

func (plan *Plan) getTableIndexs(expr sqlparser.BoolExpr) []int {
	var index int
	switch criteria := expr.(type) {
	case *sqlparser.ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>": //=对应的分片
			if plan.getValueType(criteria.Left) == EID_NODE {
				index = plan.getTableIndexByValue(criteria.Right)
			} else {
				index = plan.getTableIndexByValue(criteria.Left)
			}

			return []int{index}
		case "<", "<=":
			if plan.Rule.Type == HashRuleType {
				return plan.TableIndexs
			}

			if plan.getValueType(criteria.Left) == EID_NODE {
				index = plan.getTableIndexByValue(criteria.Right)
				if criteria.Operator == "<" {
					//调整边界值，当shard[index].start等于criteria.Right 则index--
					index = plan.adjustShardIndex(criteria.Right, index)
				}

				return makeList(0, index+1)
			} else {
				index = plan.getTableIndexByValue(criteria.Left)
				return makeList(index, len(plan.TableIndexs))
			}
		case ">", ">=":
			if plan.Rule.Type == HashRuleType {
				return plan.TableIndexs
			}

			if plan.getValueType(criteria.Left) == EID_NODE {
				index = plan.getTableIndexByValue(criteria.Right)
				return makeList(index, len(plan.TableIndexs))
			} else { // 10 > id，这种情况
				index = plan.getTableIndexByValue(criteria.Left)

				if criteria.Operator == ">" {
					index = plan.adjustShardIndex(criteria.Left, index)
				}
				return makeList(0, index+1)
			}
		case "in":
			return plan.getTableIndexsByTuple(criteria.Right)
		case "not in":
			if plan.Rule.Type == RangeRuleType {
				return plan.TableIndexs
			}

			l := plan.getTableIndexsByTuple(criteria.Right)
			return plan.notList(l)
		}
	case *sqlparser.RangeCond:
		if plan.Rule.Type == HashRuleType {
			return plan.TableIndexs
		}

		start := plan.getTableIndexByValue(criteria.From)
		last := plan.getTableIndexByValue(criteria.To)

		if criteria.Operator == "between" { //对应between ...and ...
			if last < start {
				start, last = last, start
			}
			return makeList(start, last+1)
		} else { //对应not between ....and
			if last < start {
				start, last = last, start
				start = plan.adjustShardIndex(criteria.To, start)
			} else {
				start = plan.adjustShardIndex(criteria.From, start)
			}

			l1 := makeList(0, start+1)
			l2 := makeList(last, len(plan.TableIndexs))
			return unionList(l1, l2)
		}
	default:
		return plan.TableIndexs
	}

	return plan.RouteTableIndexs
}

/*计算表下标和node下标 */
func (plan *Plan) calRouteIndexs() error {
	nodesCount := len(plan.Rule.Nodes)

	if plan.Rule.Type == DefaultRuleType {
		plan.RouteNodeIndexs = []int{0}
		return nil
	}
	if plan.Criteria == nil { //如果没有分表条件，则是全表扫描
		if plan.Rule.Type != DefaultRuleType {
			golog.Error("Plan", "calRouteIndexs", "plan have no criteria", 0,
				"type", plan.Rule.Type)
			return errors.ErrNoCriteria
		}

	}

	switch criteria := plan.Criteria.(type) {
	case sqlparser.Values: //代表insert中values
		tindex := plan.getInsertTableIndex(criteria)
		plan.RouteTableIndexs = []int{tindex}
		plan.RouteNodeIndexs = plan.TindexsToNindexs([]int{tindex})

		return nil
	case sqlparser.BoolExpr:
		plan.RouteTableIndexs = plan.getTableIndexByBoolExpr(criteria)
		plan.RouteNodeIndexs = plan.TindexsToNindexs(plan.RouteTableIndexs)

		return nil
	default:
		plan.RouteTableIndexs = plan.TableIndexs
		plan.RouteNodeIndexs = makeList(0, nodesCount)
		return nil
	}
}

func (plan *Plan) checkValuesType(vals sqlparser.Values) sqlparser.Values {
	// Analyze first value of every item in the list
	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case sqlparser.ValTuple:
			result := plan.getValueType(tuple[0])
			if result != VALUE_NODE {
				panic(sqlparser.NewParserError("insert is too complex"))
			}
		default:
			panic(sqlparser.NewParserError("insert is too complex"))
		}
	}
	return vals
}

/*返回valExpr表达式对应的类型*/
func (plan *Plan) getValueType(valExpr sqlparser.ValExpr) int {
	switch node := valExpr.(type) {
	case *sqlparser.ColName:
		if string(node.Name) == plan.Rule.Key {
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

func (plan *Plan) getTableIndexByBoolExpr(node sqlparser.BoolExpr) []int {
	switch node := node.(type) {
	case *sqlparser.AndExpr:
		left := plan.getTableIndexByBoolExpr(node.Left)
		right := plan.getTableIndexByBoolExpr(node.Right)

		return interList(left, right)
	case *sqlparser.OrExpr:
		left := plan.getTableIndexByBoolExpr(node.Left)
		right := plan.getTableIndexByBoolExpr(node.Right)
		return unionList(left, right)
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
	return plan.TableIndexs
}

/*获得(12,14,23)对应的table index*/
func (plan *Plan) getTableIndexsByTuple(valExpr sqlparser.ValExpr) []int {
	shardset := make(map[int]bool)
	switch node := valExpr.(type) {
	case sqlparser.ValTuple:
		for _, n := range node {
			index := plan.getTableIndexByValue(n)
			shardset[index] = true
		}
	}
	shardlist := make([]int, len(shardset))
	index := 0
	for k := range shardset {
		shardlist[index] = k
		index++
	}

	sort.Ints(shardlist)
	return shardlist
}

func (plan *Plan) getInsertTableIndex(vals sqlparser.Values) int {
	index := -1
	for i := 0; i < len(vals); i++ {
		first_value_expression := vals[i].(sqlparser.ValTuple)[0]
		newIndex := plan.getTableIndexByValue(first_value_expression)
		if index == -1 {
			index = newIndex
		} else if index != newIndex {
			panic(sqlparser.NewParserError("insert or replace has multiple shard targets"))
		}
	}
	return index
}

func (plan *Plan) getTableIndexByValue(valExpr sqlparser.ValExpr) int {
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
