package router

import (
	"fmt"
	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/sqlparser"
	"strings"
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
	TableToNode []int
	Shard       Shard
}

type Router struct {
	DB          string
	Rules       map[string]*Rule //key is <table name>
	DefaultRule *Rule
	nodes       []string //just for human saw
}

func NewDefaultRule(db string, node string) *Rule {
	var r *Rule = &Rule{
		DB:    db,
		Type:  DefaultRuleType,
		Nodes: []string{node},
		Shard: new(DefaultShard),
	}
	return r
}

func (r *Rule) FindNode(key interface{}) string {
	i := r.Shard.FindForKey(key)
	return r.Nodes[i]
}

func (r *Rule) FindNodeIndex(key interface{}) int {
	return r.Shard.FindForKey(key)
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s.%s?key=%v&shard=%s&nodes=%s",
		r.DB, r.Table, r.Key, r.Type, strings.Join(r.Nodes, ", "))
}

/*UpdateExprs对应set后面的表达式*/
func (r *Rule) checkUpdateExprs(exprs sqlparser.UpdateExprs) {
	if r.Type == DefaultRuleType {
		return
	} else if len(r.Nodes) == 1 {
		return
	}

	for _, e := range exprs {
		if string(e.Name.Name) == r.Key {
			panic(sqlparser.NewParserError("routing key can not in update expression"))
		}
	}
}

//router相关
/*根据配置文件建立路由规则*/
func NewRouter(schemaConfig *config.SchemaConfig) (*Router, error) {
	//default节点是否是节点列表中的一个
	if !includeNode(schemaConfig.Nodes, schemaConfig.RulesConifg.Default) {
		return nil, fmt.Errorf("default node[%s] not in the nodes list.",
			schemaConfig.RulesConifg.Default)
	}

	rt := new(Router)
	rt.DB = schemaConfig.DB       //对应schema中的db
	rt.nodes = schemaConfig.Nodes //对应schema中的nodes
	rt.Rules = make(map[string]*Rule, len(schemaConfig.RulesConifg.ShardRule))
	rt.DefaultRule = NewDefaultRule(rt.DB, schemaConfig.RulesConifg.Default)

	for _, shard := range schemaConfig.RulesConifg.ShardRule {
		//rc := &RuleConfig{shard}
		for _, node := range shard.Nodes { //rules中的nodes是不是都在schema中的nodes
			if !includeNode(rt.nodes, node) {
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
	rule := r.Rules[table]
	if rule == nil {
		return r.DefaultRule
	} else {
		return rule
	}
}

//修改这个函数，添加TableToNode
func parseRule(db string, cfg *config.ShardConfig) (*Rule, error) {
	r := new(Rule)
	r.DB = db
	r.Table = cfg.Table
	r.Key = cfg.Key
	r.Type = cfg.Type
	r.Nodes = cfg.Nodes //将ruleconfig中的nodes赋值给rule

	if err := parseShard(r, cfg); err != nil {
		return nil, err
	}

	return r, nil
}

func parseShard(r *Rule, cfg *config.ShardConfig) error {
	if r.Type == HashRuleType {
		//hash shard
		r.Shard = &HashShard{ShardNum: len(r.Nodes)}
	} else if r.Type == RangeRuleType {
		rs, err := ParseNumShardingSpec(cfg.Range)
		if err != nil {
			return err
		}

		if len(rs) != len(r.Nodes) {
			return fmt.Errorf("range space %d not equal nodes %d", len(rs), len(r.Nodes))
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

/*由sql语句获得shard node list*/
func (r *Router) GetShardList(sql string, bindVars map[string]interface{}) (nodes []string, err error) {
	var stmt sqlparser.Statement
	stmt, err = sqlparser.Parse(sql)
	if err != nil {
		return nil, err
	}

	return r.GetStmtShardList(stmt, bindVars)
}

/*由sql语句获得shard node index*/
func (r *Router) GetShardListIndex(sql string, bindVars map[string]interface{}) (nodes []int, err error) {
	var stmt sqlparser.Statement
	stmt, err = sqlparser.Parse(sql)
	if err != nil {
		return nil, err
	}

	return r.GetStmtShardListIndex(stmt, bindVars)
}

/*由sql语法树获得shard node*/
func (r *Router) GetStmtShardList(stmt sqlparser.Statement, bindVars map[string]interface{}) (nodes []string, err error) {
	defer handleError(&err)
	//获得了分表plan
	plan := r.GetPlan(stmt)

	plan.bindVars = bindVars

	ns := plan.shardListFromPlan()

	nodes = make([]string, 0, len(ns))
	for _, i := range ns {
		nodes = append(nodes, plan.rule.Nodes[i])
	}

	return nodes, nil
}

func (r *Router) GetStmtShardListIndex(stmt sqlparser.Statement, bindVars map[string]interface{}) (nodes []int, err error) {
	defer handleError(&err)

	plan := r.GetPlan(stmt)

	plan.bindVars = bindVars

	ns := plan.shardListFromPlan()

	return ns, nil
}

/*生成一个route plan*/
func (r *Router) GetPlan(statement sqlparser.Statement) (plan *Plan) {
	plan = &Plan{}
	var where *sqlparser.Where
	switch stmt := statement.(type) {
	case *sqlparser.Insert:
		if _, ok := stmt.Rows.(sqlparser.SelectStatement); ok {
			panic(sqlparser.NewParserError("select in insert not allowed"))
		}
		/*根据sql语句的表，获得对应的分片规则*/
		plan.rule = r.GetRule(sqlparser.String(stmt.Table))

		if stmt.OnDup != nil {
			plan.rule.checkUpdateExprs(sqlparser.UpdateExprs(stmt.OnDup))
		}

		plan.criteria = plan.routingAnalyzeValues(stmt.Rows.(sqlparser.Values))
		plan.fullList = makeList(0, len(plan.rule.Nodes))
		return plan
	case *sqlparser.Replace:
		if _, ok := stmt.Rows.(sqlparser.SelectStatement); ok {
			panic(sqlparser.NewParserError("select in replace not allowed"))
		}

		plan.rule = r.GetRule(sqlparser.String(stmt.Table))
		plan.criteria = plan.routingAnalyzeValues(stmt.Rows.(sqlparser.Values))
		plan.fullList = makeList(0, len(plan.rule.Nodes))
		return plan

	case *sqlparser.Select:
		plan.rule = r.GetRule(sqlparser.String(stmt.From[0])) //根据表名获得分表规则
		where = stmt.Where
	case *sqlparser.Update:
		plan.rule = r.GetRule(sqlparser.String(stmt.Table))

		plan.rule.checkUpdateExprs(stmt.Exprs)

		where = stmt.Where
	case *sqlparser.Delete:
		plan.rule = r.GetRule(sqlparser.String(stmt.Table))
		where = stmt.Where
	}

	if where != nil {
		plan.criteria = where.Expr /*路由条件*/
	} else {
		plan.rule = r.DefaultRule
	}
	plan.fullList = makeList(0, len(plan.rule.Nodes))

	return plan
}
