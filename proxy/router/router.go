// Copyright 2015 The kingshard Authors. All rights reserved.
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
	"fmt"
	"strings"

	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/sqlparser"
)

var (
	DefaultRuleType = "default"
	HashRuleType    = "hash"
	RangeRuleType   = "range"
)

type Rule struct {
	DB    string
	Table string
	Key   string

	Type string

	Nodes       []string
	TableToNode []int //index is table index,value is node index
	Shard       Shard
}

type Router struct {
	DB          string
	Rules       map[string]*Rule //key is <table name>
	DefaultRule *Rule
	Nodes       []string //just for human saw
}

func NewDefaultRule(db string, node string) *Rule {
	var r *Rule = &Rule{
		DB:          db,
		Type:        DefaultRuleType,
		Nodes:       []string{node},
		Shard:       new(DefaultShard),
		TableToNode: []int{0},
	}
	return r
}

func (r *Rule) FindNode(key interface{}) (string, error) {
	tableIndex, err := r.Shard.FindForKey(key)
	if err != nil {
		return "", err
	}
	nodeIndex := r.TableToNode[tableIndex]
	return r.Nodes[nodeIndex], nil
}

func (r *Rule) FindNodeIndex(key interface{}) (int, error) {
	tableIndex, err := r.Shard.FindForKey(key)
	if err != nil {
		return -1, err
	}
	return r.TableToNode[tableIndex], nil
}

func (r *Rule) FindTableIndex(key interface{}) (int, error) {
	return r.Shard.FindForKey(key)
}

/*UpdateExprs对应set后面的表达式*/
func (r *Rule) checkUpdateExprs(exprs sqlparser.UpdateExprs) error {
	if r.Type == DefaultRuleType {
		return nil
	} else if len(r.Nodes) == 1 {
		return nil
	}

	for _, e := range exprs {
		if string(e.Name.Name) == r.Key {
			return errors.ErrUpdateKey
		}
	}
	return nil
}

//router相关
/*根据配置文件建立路由规则*/
func NewRouter(schemaConfig *config.SchemaConfig) (*Router, error) {
	//default节点是否是节点列表中的一个
	if !includeNode(schemaConfig.Nodes, schemaConfig.RulesConfig.Default) {
		return nil, fmt.Errorf("default node[%s] not in the nodes list.",
			schemaConfig.RulesConfig.Default)
	}

	rt := new(Router)
	rt.DB = schemaConfig.DB       //对应schema中的db
	rt.Nodes = schemaConfig.Nodes //对应schema中的nodes
	rt.Rules = make(map[string]*Rule, len(schemaConfig.RulesConfig.ShardRule))
	rt.DefaultRule = NewDefaultRule(rt.DB, schemaConfig.RulesConfig.Default)

	for _, shard := range schemaConfig.RulesConfig.ShardRule {
		//rc := &RuleConfig{shard}
		for _, node := range shard.Nodes { //rules中的nodes是不是都在schema中的nodes
			if !includeNode(rt.Nodes, node) {
				return nil, fmt.Errorf("shard table[%s] node[%s] not in the schema.nodes list:[%s].",
					shard.Table, node, strings.Join(shard.Nodes, ","))
			}
		}
		rule, err := parseRule(rt.DB, &shard)
		if err != nil {
			return nil, err
		}

		if rule.Type == DefaultRuleType {
			return nil, fmt.Errorf("[default-rule] duplicate, must only one.")
		} else {
			if _, ok := rt.Rules[rule.Table]; ok {
				return nil, fmt.Errorf("table %s rule in %s duplicate", rule.Table, rule.DB)
			}
			rt.Rules[rule.Table] = rule
		}
	}
	return rt, nil
}

func (r *Router) GetRule(table string) *Rule {
	arry := strings.Split(table, ".")
	if len(arry) == 2 {
		if strings.Trim(arry[0], "`") == r.DB {
			table = strings.Trim(arry[1], "`")
		}
	}
	rule := r.Rules[table]
	if rule == nil {
		return r.DefaultRule
	} else {
		return rule
	}
}

func parseRule(db string, cfg *config.ShardConfig) (*Rule, error) {
	r := new(Rule)
	r.DB = db
	r.Table = cfg.Table
	r.Key = cfg.Key
	r.Type = cfg.Type
	r.Nodes = cfg.Nodes //将ruleconfig中的nodes赋值给rule
	r.TableToNode = make([]int, 0)

	if len(cfg.Locations) != len(r.Nodes) {
		return nil, errors.ErrLocationsCount
	}
	for i := 0; i < len(cfg.Locations); i++ {
		for j := 0; j < cfg.Locations[i]; j++ {
			r.TableToNode = append(r.TableToNode, i)
		}
	}

	if err := parseShard(r, cfg); err != nil {
		return nil, err
	}

	return r, nil
}

func parseShard(r *Rule, cfg *config.ShardConfig) error {
	if r.Type == HashRuleType {
		//hash shard
		r.Shard = &HashShard{ShardNum: len(r.TableToNode)}
	} else if r.Type == RangeRuleType {
		rs, err := ParseNumSharding(cfg.Locations, cfg.TableRowLimit)
		if err != nil {
			return err
		}

		if len(rs) != len(r.TableToNode) {
			return fmt.Errorf("range space %d not equal tables %d", len(rs), len(r.TableToNode))
		}

		r.Shard = &NumRangeShard{Shards: rs}
	} else {
		r.Shard = &DefaultShard{}
	}

	return nil
}

func includeNode(nodes []string, node string) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

/*生成一个route plan*/
func (r *Router) BuildPlan(statement sqlparser.Statement) (*Plan, error) {
	//因为实现Statement接口的方法都是指针类型，所以type对应类型也是指针类型
	switch stmt := statement.(type) {
	case *sqlparser.Insert:
		return r.buildInsertPlan(stmt)
	case *sqlparser.Replace:
		return r.buildReplacePlan(stmt)
	case *sqlparser.Select:
		return r.buildSelectPlan(stmt)
	case *sqlparser.Update:
		return r.buildUpdatePlan(stmt)
	case *sqlparser.Delete:
		return r.buildDeletePlan(stmt)
	}
	return nil, errors.ErrNoPlan
}

func (r *Router) buildSelectPlan(statement sqlparser.Statement) (*Plan, error) {
	plan := &Plan{}
	var where *sqlparser.Where
	var tableName string
	stmt := statement.(*sqlparser.Select)
	switch v := (stmt.From[0]).(type) {
	case *sqlparser.AliasedTableExpr:
		tableName = sqlparser.String(v.Expr)
	case *sqlparser.JoinTableExpr:
		if ate, ok := (v.LeftExpr).(*sqlparser.AliasedTableExpr); ok {
			tableName = sqlparser.String(ate.Expr)
		} else {
			tableName = sqlparser.String(v)
		}
	default:
		tableName = sqlparser.String(v)
	}

	plan.Rule = r.GetRule(tableName) //根据表名获得分表规则
	where = stmt.Where

	if where != nil {
		plan.Criteria = where.Expr /*路由条件*/
	} else {
		plan.Rule = r.DefaultRule
	}
	plan.TableIndexs = makeList(0, len(plan.Rule.TableToNode))

	err := plan.calRouteIndexs()
	if err != nil {
		golog.Error("Route", "BuildSelectPlan", err.Error(), 0)
		return nil, err
	}

	if plan.Rule.Type != DefaultRuleType && len(plan.RouteTableIndexs) == 0 {
		golog.Error("Route", "BuildSelectPlan", errors.ErrNoCriteria.Error(), 0)
		return nil, errors.ErrNoCriteria
	}
	//generate sql,如果routeTableindexs为空则表示不分表，不分表则发default node
	err = r.generateSelectSql(plan, stmt)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (r *Router) buildInsertPlan(statement sqlparser.Statement) (*Plan, error) {
	plan := &Plan{}
	stmt := statement.(*sqlparser.Insert)
	if _, ok := stmt.Rows.(sqlparser.SelectStatement); ok {
		return nil, errors.ErrSelectInInsert
	}
	/*根据sql语句的表，获得对应的分片规则*/
	plan.Rule = r.GetRule(sqlparser.String(stmt.Table))

	if stmt.OnDup != nil {
		err := plan.Rule.checkUpdateExprs(sqlparser.UpdateExprs(stmt.OnDup))
		if err != nil {
			return nil, err
		}
	}

	plan.Criteria = plan.checkValuesType(stmt.Rows.(sqlparser.Values))
	plan.TableIndexs = makeList(0, len(plan.Rule.TableToNode))

	err := plan.calRouteIndexs()
	if err != nil {
		golog.Error("Route", "BuildInsertPlan", err.Error(), 0)
		return nil, err
	}

	err = r.generateInsertSql(plan, stmt)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (r *Router) buildUpdatePlan(statement sqlparser.Statement) (*Plan, error) {
	plan := &Plan{}
	var where *sqlparser.Where

	stmt := statement.(*sqlparser.Update)
	plan.Rule = r.GetRule(sqlparser.String(stmt.Table))
	err := plan.Rule.checkUpdateExprs(stmt.Exprs)
	if err != nil {
		return nil, err
	}

	where = stmt.Where
	if where != nil {
		plan.Criteria = where.Expr /*路由条件*/
	} else {
		plan.Rule = r.DefaultRule
	}

	plan.TableIndexs = makeList(0, len(plan.Rule.TableToNode))

	err = plan.calRouteIndexs()
	if err != nil {
		golog.Error("Route", "BuildUpdatePlan", err.Error(), 0)
		return nil, err
	}

	if plan.Rule.Type != DefaultRuleType && len(plan.RouteTableIndexs) == 0 {
		golog.Error("Route", "BuildUpdatePlan", errors.ErrNoCriteria.Error(), 0)
		return nil, errors.ErrNoCriteria
	}
	//generate sql,如果routeTableindexs为空则表示不分表，不分表则发default node
	err = r.generateUpdateSql(plan, stmt)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (r *Router) buildDeletePlan(statement sqlparser.Statement) (*Plan, error) {
	plan := &Plan{}
	var where *sqlparser.Where

	stmt := statement.(*sqlparser.Delete)
	plan.Rule = r.GetRule(sqlparser.String(stmt.Table))
	where = stmt.Where

	if where != nil {
		plan.Criteria = where.Expr /*路由条件*/
	} else {
		plan.Rule = r.DefaultRule
	}

	plan.TableIndexs = makeList(0, len(plan.Rule.TableToNode))

	err := plan.calRouteIndexs()
	if err != nil {
		golog.Error("Route", "BuildDeletePlan", err.Error(), 0)
		return nil, err
	}

	if plan.Rule.Type != DefaultRuleType && len(plan.RouteTableIndexs) == 0 {
		golog.Error("Route", "BuildDeletePlan", errors.ErrNoCriteria.Error(), 0)
		return nil, errors.ErrNoCriteria
	}
	//generate sql,如果routeTableindexs为空则表示不分表，不分表则发default node
	err = r.generateDeleteSql(plan, stmt)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (r *Router) buildReplacePlan(statement sqlparser.Statement) (*Plan, error) {
	plan := &Plan{}

	stmt := statement.(*sqlparser.Replace)
	if _, ok := stmt.Rows.(sqlparser.SelectStatement); ok {
		panic(sqlparser.NewParserError("select in replace not allowed"))
	}

	plan.Rule = r.GetRule(sqlparser.String(stmt.Table))
	plan.Criteria = plan.checkValuesType(stmt.Rows.(sqlparser.Values))

	plan.TableIndexs = makeList(0, len(plan.Rule.TableToNode))

	err := plan.calRouteIndexs()
	if err != nil {
		golog.Error("Route", "BuildReplacePlan", err.Error(), 0)
		return nil, err
	}

	err = r.generateReplaceSql(plan, stmt)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (r *Router) generateSelectSql(plan *Plan, stmt sqlparser.Statement) error {
	sqls := make(map[string][]string)
	node, ok := stmt.(*sqlparser.Select)
	if ok == false {
		return errors.ErrStmtConvert
	}
	if len(plan.RouteNodeIndexs) == 0 {
		return errors.ErrNoRouteNode
	}
	if len(plan.RouteTableIndexs) == 0 {
		buf := sqlparser.NewTrackedBuffer(nil)
		stmt.Format(buf)
		nodeName := r.Nodes[0]
		sqls[nodeName] = []string{buf.String()}
	} else {
		tableCount := len(plan.RouteTableIndexs)
		for i := 0; i < tableCount; i++ {
			buf := sqlparser.NewTrackedBuffer(nil)
			buf.Fprintf("select %v%s%v from ",
				node.Comments,
				node.Distinct,
				node.SelectExprs,
			)
			switch v := (node.From[0]).(type) {
			case *sqlparser.AliasedTableExpr:
				if len(v.As) != 0 {
					fmt.Fprintf(buf, "%s_%04d AS %s",
						sqlparser.String(v.Expr),
						plan.RouteTableIndexs[i],
						string(v.As),
					)
				} else {
					fmt.Fprintf(buf, "%s_%04d",
						sqlparser.String(v.Expr),
						plan.RouteTableIndexs[i],
					)
				}
			case *sqlparser.JoinTableExpr:
				if ate, ok := (v.LeftExpr).(*sqlparser.AliasedTableExpr); ok {
					if len(ate.As) != 0 {
						fmt.Fprintf(buf, "%s_%04d AS %s",
							sqlparser.String(ate.Expr),
							plan.RouteTableIndexs[i],
							string(ate.As),
						)
					} else {
						fmt.Fprintf(buf, "%s_%04d",
							sqlparser.String(ate.Expr),
							plan.RouteTableIndexs[i],
						)
					}
				} else {
					fmt.Fprintf(buf, "%s_%04d",
						sqlparser.String(v.LeftExpr),
						plan.RouteTableIndexs[i],
					)
				}
				buf.Fprintf(" %s %v", v.Join, v.RightExpr)
				if v.On != nil {
					buf.Fprintf(" on %v", v.On)
				}
			default:
				fmt.Fprintf(buf, "%s_%04d",
					sqlparser.String(node.From[0]),
					plan.RouteTableIndexs[i],
				)
			}
			buf.Fprintf("%v%v%v%v%v%s",
				node.Where,
				node.GroupBy,
				node.Having,
				node.OrderBy,
				node.Limit,
				node.Lock,
			)

			tableIndex := plan.RouteTableIndexs[i]
			nodeIndex := plan.Rule.TableToNode[tableIndex]
			nodeName := r.Nodes[nodeIndex]
			if _, ok := sqls[nodeName]; ok == false {
				sqls[nodeName] = make([]string, 0, tableCount)
			}
			sqls[nodeName] = append(sqls[nodeName], buf.String())
		}

	}
	plan.RewrittenSqls = sqls
	return nil
}

func (r *Router) generateInsertSql(plan *Plan, stmt sqlparser.Statement) error {
	sqls := make(map[string][]string)
	node, ok := stmt.(*sqlparser.Insert)
	if ok == false {
		return errors.ErrStmtConvert
	}
	if len(plan.RouteNodeIndexs) == 0 {
		return errors.ErrNoRouteNode
	}
	if len(plan.RouteTableIndexs) == 0 {
		buf := sqlparser.NewTrackedBuffer(nil)
		stmt.Format(buf)
		nodeName := r.Nodes[0]
		sqls[nodeName] = []string{buf.String()}
	} else {
		nodeCount := len(plan.RouteNodeIndexs)
		if 1 < nodeCount {
			golog.Error("Router", "generateInsertSql", errors.ErrInsertInMulti.Error(), 0)
			return errors.ErrInsertInMulti
		}
		tableCount := len(plan.RouteTableIndexs)
		for i := 0; i < tableCount; i++ {
			buf := sqlparser.NewTrackedBuffer(nil)
			buf.Fprintf("insert %vinto %v", node.Comments, node.Table)
			fmt.Fprintf(buf, "_%04d", plan.RouteTableIndexs[i])
			buf.Fprintf("%v %v%v",
				node.Columns,
				node.Rows,
				node.OnDup)

			tableIndex := plan.RouteTableIndexs[i]
			nodeIndex := plan.Rule.TableToNode[tableIndex]
			nodeName := r.Nodes[nodeIndex]
			if _, ok := sqls[nodeName]; ok == false {
				sqls[nodeName] = make([]string, 0, tableCount)
			}
			sqls[nodeName] = append(sqls[nodeName], buf.String())
		}

	}
	plan.RewrittenSqls = sqls
	return nil
}

func (r *Router) generateUpdateSql(plan *Plan, stmt sqlparser.Statement) error {
	sqls := make(map[string][]string)
	node, ok := stmt.(*sqlparser.Update)
	if ok == false {
		return errors.ErrStmtConvert
	}
	if len(plan.RouteNodeIndexs) == 0 {
		return errors.ErrNoRouteNode
	}
	if len(plan.RouteTableIndexs) == 0 {
		buf := sqlparser.NewTrackedBuffer(nil)
		stmt.Format(buf)
		nodeName := r.Nodes[0]
		sqls[nodeName] = []string{buf.String()}
	} else {
		nodeCount := len(plan.RouteNodeIndexs)
		if 1 < nodeCount {
			golog.Error("Router", "generateUpdateSql", errors.ErrUpdateInMulti.Error(), 0,
				"RouteNodeIndexs", plan.RouteNodeIndexs)
			return errors.ErrUpdateInMulti
		}
		tableCount := len(plan.RouteTableIndexs)
		for i := 0; i < tableCount; i++ {
			buf := sqlparser.NewTrackedBuffer(nil)
			buf.Fprintf("update %v%v",
				node.Comments,
				node.Table,
			)
			fmt.Fprintf(buf, "_%04d", plan.RouteTableIndexs[i])
			buf.Fprintf(" set %v%v%v%v",
				node.Exprs,
				node.Where,
				node.OrderBy,
				node.Limit,
			)
			tableIndex := plan.RouteTableIndexs[i]
			nodeIndex := plan.Rule.TableToNode[tableIndex]
			nodeName := r.Nodes[nodeIndex]
			if _, ok := sqls[nodeName]; ok == false {
				sqls[nodeName] = make([]string, 0, tableCount)
			}
			sqls[nodeName] = append(sqls[nodeName], buf.String())
		}

	}
	plan.RewrittenSqls = sqls
	return nil
}

func (r *Router) generateDeleteSql(plan *Plan, stmt sqlparser.Statement) error {
	sqls := make(map[string][]string)
	node, ok := stmt.(*sqlparser.Delete)
	if ok == false {
		return errors.ErrStmtConvert
	}
	if len(plan.RouteNodeIndexs) == 0 {
		return errors.ErrNoRouteNode
	}
	if len(plan.RouteTableIndexs) == 0 {
		buf := sqlparser.NewTrackedBuffer(nil)
		stmt.Format(buf)
		nodeName := r.Nodes[0]
		sqls[nodeName] = []string{buf.String()}
	} else {
		nodeCount := len(plan.RouteNodeIndexs)
		if 1 < nodeCount {
			golog.Error("Router", "generateDeleteSql", errors.ErrDeleteInMulti.Error(), 0)
			return errors.ErrUpdateInMulti
		}
		tableCount := len(plan.RouteTableIndexs)
		for i := 0; i < tableCount; i++ {
			buf := sqlparser.NewTrackedBuffer(nil)
			buf.Fprintf("delete %vfrom %v",
				node.Comments,
				node.Table,
			)
			fmt.Fprintf(buf, "_%04d", plan.RouteTableIndexs[i])
			buf.Fprintf("%v%v%v",
				node.Where,
				node.OrderBy,
				node.Limit,
			)
			tableIndex := plan.RouteTableIndexs[i]
			nodeIndex := plan.Rule.TableToNode[tableIndex]
			nodeName := r.Nodes[nodeIndex]
			if _, ok := sqls[nodeName]; ok == false {
				sqls[nodeName] = make([]string, 0, tableCount)
			}
			sqls[nodeName] = append(sqls[nodeName], buf.String())
		}

	}
	plan.RewrittenSqls = sqls
	return nil
}

func (r *Router) generateReplaceSql(plan *Plan, stmt sqlparser.Statement) error {
	sqls := make(map[string][]string)
	node, ok := stmt.(*sqlparser.Replace)
	if ok == false {
		return errors.ErrStmtConvert
	}
	if len(plan.RouteNodeIndexs) == 0 {
		return errors.ErrNoRouteNode
	}
	if len(plan.RouteTableIndexs) == 0 {
		buf := sqlparser.NewTrackedBuffer(nil)
		stmt.Format(buf)
		nodeName := r.Nodes[0]
		sqls[nodeName] = []string{buf.String()}
	} else {
		nodeCount := len(plan.RouteNodeIndexs)
		if 1 < nodeCount {
			golog.Error("Router", "generateReplaceSql", errors.ErrReplaceInMulti.Error(), 0)
			return errors.ErrUpdateInMulti
		}
		tableCount := len(plan.RouteTableIndexs)
		for i := 0; i < tableCount; i++ {
			buf := sqlparser.NewTrackedBuffer(nil)
			buf.Fprintf("replace %vinto %v",
				node.Comments,
				node.Table,
			)
			fmt.Fprintf(buf, "_%04d", plan.RouteTableIndexs[i])
			buf.Fprintf("%v %v",
				node.Columns,
				node.Rows,
			)
			tableIndex := plan.RouteTableIndexs[i]
			nodeIndex := plan.Rule.TableToNode[tableIndex]
			nodeName := r.Nodes[nodeIndex]
			if _, ok := sqls[nodeName]; ok == false {
				sqls[nodeName] = make([]string, 0, tableCount)
			}
			sqls[nodeName] = append(sqls[nodeName], buf.String())
		}

	}
	plan.RewrittenSqls = sqls
	return nil
}
