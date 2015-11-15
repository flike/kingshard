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
	"testing"

	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/yaml"
	"github.com/flike/kingshard/sqlparser"
)

func TestParseRule(t *testing.T) {
	var s = `
schemas:
-
  db: kingshard
  nodes: [node1, node2, node3]
  rules:
    default: node1
    shard:      
     -
      table: test_shard_hash
      key: id
      nodes: [node2, node3]
      locations: [16,16]
      type: hash
     -
      table: test_shard_range
      key: id
      type: range
      nodes: [node2, node3]
      locations: [16,16]
      table_row_limit: 10000
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

	hashRule := rt.GetRule("test_shard_hash")
	if hashRule.Type != HashRuleType {
		t.Fatal(hashRule.Type)
	}

	if len(hashRule.Nodes) != 2 || hashRule.Nodes[0] != "node2" || hashRule.Nodes[1] != "node3" {
		t.Fatal("parse nodes not correct.")
	}

	if n, _ := hashRule.FindNode(uint64(11)); n != "node2" {
		t.Fatal(n)
	}

	rangeRule := rt.GetRule("test_shard_range")
	if rangeRule.Type != RangeRuleType {
		t.Fatal(rangeRule.Type)
	}

	if n, _ := rangeRule.FindNode(10000 - 1); n != "node2" {
		t.Fatal(n)
	}

	defaultRule := rt.GetRule("defaultRule_table")
	if defaultRule == nil {
		t.Fatal("must not nil")
	}

	if defaultRule.Type != DefaultRuleType {
		t.Fatal(defaultRule.Type)
	}

	if defaultRule.Shard == nil {
		t.Fatal("nil error")
	}

	if n, _ := defaultRule.FindNode(11); n != "node1" {
		t.Fatal(n)
	}
}

func newTestRouter() *Router {
	var s = `
schemas :
-
  db : kingshard
  nodes: [node1,node2,node3,node4,node5,node6,node7,node8,node9,node10]
  rules:
    default: node1
    shard:
      -
        table: test1
        key: id
        nodes: [node1,node2,node3]
        locations: [4,4,4]
        type: hash

      -
        table: test2
        key: id
        type: range
        nodes: [node1,node2,node3]
        locations: [4,4,4]
        table_row_limit: 10000
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

func newTestDBRule() *Router {
	var s = `
schemas :
-
  db : kingshard
  nodes: [node1,node2,node3,node4,node5,node6,node7,node8,node9,node10]
  rules:
    default: node1
    shard:
      -
        table: test1
        key: id
        nodes: [node1,node2,node3]
        locations: [1,2,3]
        type: hash

      -
        table: test2
        key: id
        type: range
        nodes: [node1,node2,node3]
        locations: [8,8,8]
        table_row_limit: 100
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

	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		t.Fatal(err.Error())
	}

	if _, err := r.BuildPlan(stmt); err == nil {
		t.Fatal("must err")
	}

	sql = "update test1 set id = 10 where id = 5"

	stmt, err = sqlparser.Parse(sql)
	if err != nil {
		t.Fatal(err.Error())
	}

	if _, err := r.BuildPlan(stmt); err == nil {
		t.Fatal("must err")
	}
}

func isListEqual(l1 []int, l2 []int) bool {
	var i, j int
	if len(l1) != len(l2) {
		return false
	}
	if len(l1) == 0 {
		return true
	}
	for i = 0; i < len(l1); i++ {
		for j = 0; j < len(l2); j++ {
			if l1[i] == l2[j] {
				break
			}
		}
		if j == len(l2) {
			return false
		}
	}
	return true
}

func checkPlan(t *testing.T, sql string, tableIndexs []int, nodeIndexs []int) {
	r := newTestRouter()
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		t.Fatal(err.Error())
	}
	plan, err := r.BuildPlan(stmt)
	if err != nil {
		t.Fatal(err.Error())
	}

	if isListEqual(plan.RouteTableIndexs, tableIndexs) == false {
		err := fmt.Errorf("RouteTableIndexs=%v but tableIndexs=%v",
			plan.RouteTableIndexs, tableIndexs)
		t.Fatal(err.Error())
	}
	if isListEqual(plan.RouteNodeIndexs, nodeIndexs) == false {
		err := fmt.Errorf("RouteNodeIndexs=%v but nodeIndexs=%v",
			plan.RouteNodeIndexs, nodeIndexs)
		t.Fatal(err.Error())
	}
	t.Logf("rewritten_sql=%v", plan.RewrittenSqls)

}

func TestSelectPlan(t *testing.T) {
	var sql string
	t1 := makeList(0, 12)

	sql = "select/*master*/ * from test1 where id = 5"
	checkPlan(t, sql, []int{5}, []int{1}) //table_5 node1

	sql = "select * from test1 where id in (5, 8)"
	checkPlan(t, sql, []int{5, 8}, []int{1, 2})

	sql = "select * from test1 where id > 5"

	checkPlan(t, sql, t1, []int{0, 1, 2})

	sql = "select * from test1 where id in (5, 6) and id in (5, 6, 7)"
	checkPlan(t, sql, []int{5, 6}, []int{1})

	sql = "select * from test1 where id in (5, 6) or id in (5, 6, 7,8)"
	checkPlan(t, sql, []int{5, 6, 7, 8}, []int{1, 2})

	sql = "select * from test1 where id not in (5, 6) or id in (5, 6, 7,8)"
	checkPlan(t, sql, t1, []int{0, 1, 2})

	sql = "select * from test1 where id not in (5, 6)"
	checkPlan(t, sql, []int{0, 1, 2, 3, 4, 7, 8, 9, 10, 11}, []int{0, 1, 2})

	sql = "select * from test1 where id in (5, 6) or (id in (5, 6, 7,8) and id in (1,5,7))"
	checkPlan(t, sql, []int{5, 6, 7}, []int{1})

	sql = "select * from test2 where id = 10000"
	checkPlan(t, sql, []int{1}, []int{0})

	sql = "select * from test2 where id between 10000 and 20000"
	checkPlan(t, sql, []int{1, 2}, []int{0})

	sql = "select * from test2 where id not between 1000 and 100000"
	checkPlan(t, sql, []int{0, 10, 11}, []int{0, 2})

	sql = "select * from test2 where id > 10000"
	checkPlan(t, sql, makeList(1, 12), []int{0, 1, 2})

	sql = "select * from test2 where id >= 9999"
	checkPlan(t, sql, t1, []int{0, 1, 2})

	sql = "select * from test2 where id <= 10000"
	checkPlan(t, sql, []int{0, 1}, []int{0})

	sql = "select * from test2 where id < 10000"
	checkPlan(t, sql, []int{0}, []int{0})

	sql = "select * from test2 where id >= 10000 and id <= 30000"
	checkPlan(t, sql, []int{1, 2, 3}, []int{0})

	sql = "select * from test2 where (id >= 10000 and id <= 30000) or id < 100"
	checkPlan(t, sql, []int{0, 1, 2, 3}, []int{0})

	sql = "select * from test2 where id in (1, 10000)"
	checkPlan(t, sql, []int{0, 1}, []int{0})

	sql = "select * from test2 where id not in (1, 10000)"
	checkPlan(t, sql, makeList(0, 12), []int{0, 1, 2})
}

func TestValueSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (5)"
	checkPlan(t, sql, []int{5}, []int{1})

	sql = "insert into test2 (id) values (10000)"
	checkPlan(t, sql, []int{1}, []int{0})

	sql = "insert into test2 (id) values (20000)"
	checkPlan(t, sql, []int{2}, []int{0})

	sql = "update test1 set a =10 where id =12"
	checkPlan(t, sql, []int{0}, []int{0})

	sql = "update test2 set a =10 where id < 30000 and 10000< id"
	checkPlan(t, sql, []int{1, 2}, []int{0})

	sql = "delete from test2 where id < 30000 and 10000< id"
	checkPlan(t, sql, []int{1, 2}, []int{0})

	sql = "replace into test1(id) values(5)"
	checkPlan(t, sql, []int{5}, []int{1})
}
