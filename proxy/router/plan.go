// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/flike/kingshard/sqlparser"
	"sort"
	"strconv"
)

const (
	EID_NODE = iota
	VALUE_NODE
	LIST_NODE
	OTHER_NODE
)

type Plan struct {
	rule *Rule

	criteria sqlparser.SQLNode

	fullList []int

	bindVars map[string]interface{}
}

/*
	Limitation:

	where, eg, key name is id:

		where id = 1
		where id in (1, 2, 3)
		where id > 1
		where id >= 1
		where id < 1
		where id <= 1
		where id between 1 and 10
		where id >= 1 and id < 10
*/

func (plan *Plan) notList(l []int) []int {
	return differentList(plan.fullList, l)
}

/*根据条件表达式得到shard list*/
func (plan *Plan) findConditionShard(expr sqlparser.BoolExpr) (shardList []int) {
	var index int
	switch criteria := expr.(type) {
	case *sqlparser.ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>": //=对应的分片
			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				index = plan.findShard(criteria.Right)
			} else {
				index = plan.findShard(criteria.Left)
			}
			return []int{index}
		case "<", "<=":
			if plan.rule.Type == HashRuleType {
				return plan.fullList
			}

			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				index = plan.findShard(criteria.Right)
				if criteria.Operator == "<" {
					index = plan.adjustShardIndex(criteria.Right, index) //调整边界值，当shard[index].start等于criteria.Right 则index--
				}

				return makeList(0, index+1)
			} else {
				index = plan.findShard(criteria.Left)
				return makeList(index, len(plan.rule.Nodes))
			}
		case ">", ">=":
			if plan.rule.Type == HashRuleType {
				return plan.fullList
			}

			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				index = plan.findShard(criteria.Right)
				return makeList(index, len(plan.rule.Nodes))
			} else { // 10 > id，这种情况
				index = plan.findShard(criteria.Left)

				if criteria.Operator == ">" {
					index = plan.adjustShardIndex(criteria.Left, index)
				}
				return makeList(0, index+1)
			}
		case "in":
			return plan.findShardList(criteria.Right)
		case "not in":
			if plan.rule.Type == RangeRuleType {
				return plan.fullList
			}

			l := plan.findShardList(criteria.Right)
			return plan.notList(l)
		}
	case *sqlparser.RangeCond:
		if plan.rule.Type == HashRuleType {
			return plan.fullList
		}

		start := plan.findShard(criteria.From)
		last := plan.findShard(criteria.To)

		if criteria.Operator == "between" { //对应between ...and ...
			if last < start {
				start, last = last, start
			}
			l := makeList(start, last+1)
			return l
		} else { //对应not between ....and
			if last < start {
				start, last = last, start
				start = plan.adjustShardIndex(criteria.To, start)
			} else {
				start = plan.adjustShardIndex(criteria.From, start)
			}

			l1 := makeList(0, start+1)
			l2 := makeList(last, len(plan.rule.Nodes))
			return unionList(l1, l2)
		}
	default:
		return plan.fullList
	}

	return plan.fullList
}

/*从plan中得到shard list*/
func (plan *Plan) shardListFromPlan() (shardList []int) {
	if plan.criteria == nil { //如果没有分表条件，则是全表扫描
		return plan.fullList
	}

	//default rule will route all sql to one node
	//if rule has one node, we also can route directly
	if plan.rule.Type == DefaultRuleType || len(plan.rule.Nodes) == 1 {
		if len(plan.fullList) != 1 {
			panic(sqlparser.NewParserError("invalid rule nodes num %d, must 1", plan.fullList))
		}
		return plan.fullList
	}

	switch criteria := plan.criteria.(type) {
	case sqlparser.Values: //代表insert中values
		index := plan.findInsertShard(criteria)
		return []int{index}
	case sqlparser.BoolExpr:
		return plan.routingAnalyzeBoolean(criteria)
	default:
		return plan.fullList
	}
}

func (plan *Plan) routingAnalyzeValues(vals sqlparser.Values) sqlparser.Values {
	// Analyze first value of every item in the list
	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case sqlparser.ValTuple:
			result := plan.routingAnalyzeValue(tuple[0])
			if result != VALUE_NODE {
				panic(sqlparser.NewParserError("insert is too complex"))
			}
		default:
			panic(sqlparser.NewParserError("insert is too complex"))
		}
	}
	return vals
}

//bool表达式类型的分表规则，返回bool表达式对应的shard list
func (plan *Plan) routingAnalyzeBoolean(node sqlparser.BoolExpr) []int {
	switch node := node.(type) {
	case *sqlparser.AndExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)

		return interList(left, right)
	case *sqlparser.OrExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)
		return unionList(left, right)
	case *sqlparser.ParenBoolExpr:
		return plan.routingAnalyzeBoolean(node.Expr)
	case *sqlparser.ComparisonExpr:
		switch {
		case sqlparser.StringIn(node.Operator, "=", "<", ">", "<=", ">=", "<=>"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if (left == EID_NODE && right == VALUE_NODE) || (left == VALUE_NODE && right == EID_NODE) {
				return plan.findConditionShard(node)
			}
		case sqlparser.StringIn(node.Operator, "in", "not in"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if left == EID_NODE && right == LIST_NODE {
				return plan.findConditionShard(node)
			}
		}
	case *sqlparser.RangeCond:
		left := plan.routingAnalyzeValue(node.Left)
		from := plan.routingAnalyzeValue(node.From)
		to := plan.routingAnalyzeValue(node.To)
		if left == EID_NODE && from == VALUE_NODE && to == VALUE_NODE {
			return plan.findConditionShard(node)
		}
	}
	return plan.fullList
}

/*返回valExpr表达式对应的类型*/
func (plan *Plan) routingAnalyzeValue(valExpr sqlparser.ValExpr) int {
	switch node := valExpr.(type) {
	case *sqlparser.ColName:
		if string(node.Name) == plan.rule.Key {
			return EID_NODE //表示这是分片id对应的node
		}
	case sqlparser.ValTuple:
		for _, n := range node {
			if plan.routingAnalyzeValue(n) != VALUE_NODE {
				return OTHER_NODE
			}
		}
		return LIST_NODE //列表节点
	case sqlparser.StrVal, sqlparser.NumVal, sqlparser.ValArg: //普通的值节点，字符串值，绑定变量参数
		return VALUE_NODE
	}
	return OTHER_NODE
}

/*获得(12,14,23)对应的shard node*/
func (plan *Plan) findShardList(valExpr sqlparser.ValExpr) []int {
	shardset := make(map[int]bool)
	switch node := valExpr.(type) {
	case sqlparser.ValTuple:
		for _, n := range node {
			index := plan.findShard(n)
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

func (plan *Plan) findInsertShard(vals sqlparser.Values) int {
	index := -1
	for i := 0; i < len(vals); i++ {
		first_value_expression := vals[i].(sqlparser.ValTuple)[0]
		newIndex := plan.findShard(first_value_expression)
		if index == -1 {
			index = newIndex
		} else if index != newIndex {
			panic(sqlparser.NewParserError("insert has multiple shard targets"))
		}
	}
	return index
}

func (plan *Plan) findShard(valExpr sqlparser.ValExpr) int {
	value := plan.getBoundValue(valExpr)
	return plan.rule.FindNodeIndex(value)
}

func (plan *Plan) adjustShardIndex(valExpr sqlparser.ValExpr, index int) int {
	value := plan.getBoundValue(valExpr)
	//生成一个范围的接口,[100,120]
	s, ok := plan.rule.Shard.(RangeShard) //Shard是一个interface，在这强转为RangeShard,此处语法问题？？
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
		return plan.bindVars[string(node[1:])]
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
