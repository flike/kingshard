package router

import (
	"fmt"
	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/yaml"
	"testing"
)

func TestParseRule(t *testing.T) {
	var s = `
schemas :
-
  db : mixer
  nodes: [node1, node2, node3]
  rules:
    default: node1
    shard:
      -   
        table: mixer_test_shard_hash
        key: id
        nodes: [node2, node3]
        type: hash

      -   
        table: mixer_test_shard_range
        key: id
        type: range
        nodes: [node2, node3]
        range: -10000-
`
	var cfg config.Config
	if err := yaml.Unmarshal([]byte(s), &cfg); err != nil {
		t.Fatal(err)
	}

	rt, err := NewRouter(&cfg.Schemas[0])
	if err != nil {
		t.Fatal(err)
	}
	if rt.DefaultRule.Nodes[0] != "node1" {
		t.Fatal("default rule parse not correct.")
	}

	hashRule := rt.GetRule("mixer_test_shard_hash")
	if hashRule.Type != HashRuleType {
		t.Fatal(hashRule.Type)
	}

	if len(hashRule.Nodes) != 2 || hashRule.Nodes[0] != "node2" || hashRule.Nodes[1] != "node3" {
		t.Fatal("parse nodes not correct.")
	}

	if n := hashRule.FindNode(uint64(11)); n != "node3" {
		t.Fatal(n)
	}

	rangeRule := rt.GetRule("mixer_test_shard_range")
	if rangeRule.Type != RangeRuleType {
		t.Fatal(rangeRule.Type)
	}

	if n := rangeRule.FindNode(10000 - 1); n != "node2" {
		t.Fatal(n)
	}

	defaultRule := rt.GetRule("mixer_defaultRule_table")
	if defaultRule == nil {
		t.Fatal("must not nil")
	}

	if defaultRule.Type != DefaultRuleType {
		t.Fatal(defaultRule.Type)
	}

	if defaultRule.Shard == nil {
		t.Fatal("nil error")
	}

	if n := defaultRule.FindNode(11); n != "node1" {
		t.Fatal(n)
	}
}

func newTestDBRule() *Router {
	var s = `
schemas :
-
  db : mixer
  nodes: [node1,node2,node3,node4,node5,node6,node7,node8,node9,node10]
  rules:
    default: node1
    shard:
      -
        table: test1
        key: id
        nodes: [node1,node2,node3,node4,node5,node6,node7,node8,node9,node10]
        type: hash

      -
        table: test2
        key: id
        type: range
        nodes: [node1,node2,node3]
        range: -10000-20000-
`

	cfg, err := config.ParseConfigData([]byte(s))
	if err != nil {
		println(err.Error())
		panic(err)
	}

	var r *Router

	r, err = NewRouter(&cfg.Schemas[0])
	if err != nil {
		println(err.Error())
		panic(err)
	}

	return r
}

func TestBadUpdateExpr(t *testing.T) {
	var sql string

	r := newTestDBRule()

	sql = "insert into test1 (id) values (5) on duplicate key update  id = 10"

	if _, err := r.GetShardList(sql, nil); err == nil {
		t.Fatal("must err")
	}

	sql = "update test1 set id = 10 where id = 5"

	if _, err := r.GetShardList(sql, nil); err == nil {
		t.Fatal("must err")
	}
}

func checkSharding(t *testing.T, sql string, args []int, checkNodeIndex ...int) {
	r := newTestDBRule()

	bindVars := make(map[string]interface{}, len(args))
	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}
	ns, err := r.GetShardListIndex(sql, bindVars)
	if err != nil {
		t.Fatal(sql, err)
	} else if len(ns) != len(checkNodeIndex) {
		s := fmt.Sprintf("%v %v", ns, checkNodeIndex)
		t.Fatal(sql, s)
	} else {
		for i := range ns {
			if ns[i] != checkNodeIndex[i] {
				s := fmt.Sprintf("%v %v", ns, checkNodeIndex)
				panic(sql)
				t.Fatal(sql, s, i)
			}
		}
	}
}

func TestConditionSharding(t *testing.T) {
	var sql string

	sql = "select * from test1 where id = 5"
	checkSharding(t, sql, nil, 5)

	sql = "select * from test1 where id in (5, 6)"
	checkSharding(t, sql, nil, 5, 6)

	sql = "select * from test1 where id > 5"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6) and id in (5, 6, 7)"
	checkSharding(t, sql, nil, 5, 6)

	sql = "select * from test1 where id in (5, 6) or id in (5, 6, 7,8)"
	checkSharding(t, sql, nil, 5, 6, 7, 8)

	sql = "select * from test1 where id not in (5, 6) or id in (5, 6, 7,8)"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id not in (5, 6)"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6) or (id in (5, 6, 7,8) and id in (1,5,7))"
	checkSharding(t, sql, nil, 5, 6, 7)

	sql = "select * from test2 where id = 10000"
	checkSharding(t, sql, nil, 1)

	sql = "select * from test2 where id between 10000 and 100000"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where id not between 1000 and 100000"
	checkSharding(t, sql, nil, 0, 2)

	sql = "select * from test2 where id not between 10000 and 100000"
	checkSharding(t, sql, nil, 0, 2)

	sql = "select * from test2 where id > 10000"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where id >= 10000"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where id <= 10000"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id < 10000"
	checkSharding(t, sql, nil, 0)

	sql = "select * from test2 where  10000 < id"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where  10000 <= id"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where  10000 > id"
	checkSharding(t, sql, nil, 0)

	sql = "select * from test2 where  10000 >= id"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id >= 10000 and id <= 100000"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where (id >= 10000 and id <= 100000) or id < 100"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where (id >= 10000 and id <= 100000) or (id < 100 and name > 100000)"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id in (1, 10000)"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id not in (1, 10000)"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id in (1000, 10000)"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id > -1"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id > -1 and id < 11000"
	checkSharding(t, sql, nil, 0, 1)
}

func TestConditionVarArgSharding(t *testing.T) {
	var sql string

	sql = "select * from test1 where id = ?"
	checkSharding(t, sql, []int{5}, 5)

	sql = "select * from test1 where id in (?, ?)"
	checkSharding(t, sql, []int{5, 6}, 5, 6)

	sql = "select * from test1 where id > ?"
	checkSharding(t, sql, []int{5}, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (?, ?) and id in (?, ?, ?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7}, 5, 6)

	sql = "select * from test1 where id in (?, ?) or id in (?, ?,?,?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8}, 5, 6, 7, 8)

	sql = "select * from test1 where id not in (?, ?) or id in (?, ?, ?,?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8}, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id not in (?, ?)"
	checkSharding(t, sql, []int{5, 6}, 0, 1, 2, 3, 4, 7, 8, 9)

	sql = "select * from test1 where id in (?, ?) or (id in (?, ?, ?,?) and id in (?,?,?))"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8, 1, 5, 7}, 5, 6, 7)

	sql = "select * from test2 where id = ?"
	checkSharding(t, sql, []int{10000}, 1)

	sql = "select * from test2 where id between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 1, 2)

	sql = "select * from test2 where id not between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 0, 2)

	sql = "select * from test2 where id not between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 0, 2)

	sql = "select * from test2 where id > ?"
	checkSharding(t, sql, []int{10000}, 1, 2)

	sql = "select * from test2 where id >= ?"
	checkSharding(t, sql, []int{10000}, 1, 2)

	sql = "select * from test2 where id <= ?"
	checkSharding(t, sql, []int{10000}, 0, 1)

	sql = "select * from test2 where id < ?"
	checkSharding(t, sql, []int{10000}, 0)

	sql = "select * from test2 where  ? < id"
	checkSharding(t, sql, []int{10000}, 1, 2)

	sql = "select * from test2 where  ? <= id"
	checkSharding(t, sql, []int{10000}, 1, 2)

	sql = "select * from test2 where  ? > id"
	checkSharding(t, sql, []int{10000}, 0)

	sql = "select * from test2 where  ? >= id"
	checkSharding(t, sql, []int{10000}, 0, 1)

	sql = "select * from test2 where id >= ? and id <= ?"
	checkSharding(t, sql, []int{10000, 100000}, 1, 2)

	sql = "select * from test2 where (id >= ? and id <= ?) or id < ?"
	checkSharding(t, sql, []int{10000, 100000, 100}, 0, 1, 2)

	sql = "select * from test2 where (id >= ? and id <= ?) or (id < ? and name > ?)"
	checkSharding(t, sql, []int{10000, 100000, 100, 100000}, 0, 1, 2)

	sql = "select * from test2 where id in (?, ?)"
	checkSharding(t, sql, []int{1, 10000}, 0, 1)

	sql = "select * from test2 where id not in (?, ?)"
	checkSharding(t, sql, []int{1, 10000}, 0, 1, 2)

	sql = "select * from test2 where id in (?, ?)"
	checkSharding(t, sql, []int{1000, 10000}, 0, 1)

	sql = "select * from test2 where id > ?"
	checkSharding(t, sql, []int{-1}, 0, 1, 2)

	sql = "select * from test2 where id > ? and id < ?"
	checkSharding(t, sql, []int{-1, 11000}, 0, 1)
}

func TestValueSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (5)"
	checkSharding(t, sql, nil, 5)

	sql = "insert into test2 (id) values (10000)"
	checkSharding(t, sql, nil, 1)

	sql = "insert into test2 (id) values (20000)"
	checkSharding(t, sql, nil, 2)

	sql = "insert into test2 (id) values (200000)"
	checkSharding(t, sql, nil, 2)
}

func TestValueVarArgSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (?)"
	checkSharding(t, sql, []int{5}, 5)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{10000}, 1)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{20000}, 2)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{200000}, 2)
}
