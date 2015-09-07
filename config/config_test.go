package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	var testConfigData = []byte(
		`
addr : 127.0.0.1:3601
user : root
password : root
log_level : error

nodes :
- 
  name : node1 
  down_after_noalive : 300
  idle_conns : 16
  user: root
  password: root
  master : 127.0.0.1:3306
  slave : 127.0.0.1:4306
- 
  name : node2
  user: root
  master : 127.0.0.1:3307

- 
  name : node3 
  down_after_noalive : 300
  idle_conns : 16
  user: root
  password: root
  master : 127.0.0.1:3308

schemas :
-
  db : kingshard 
  nodes: [node1, node2, node3]
  rules:
    default: node1
    shard:
      -   
        table: test_shard_hash
        key: id
        nodes: [node1, node2, node3]
        type: hash
		locations: [4,4,4]

      -   
        table: test_shard_range
        key: id
        type: range
        nodes: [node2, node3]
        locations: [4,4]
		table_row_limit:10000
`)

	cfg, err := ParseConfigData(testConfigData)
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Nodes) != 3 {
		t.Fatal(len(cfg.Nodes))
	}

	if len(cfg.Schemas) != 1 {
		t.Fatal(len(cfg.Schemas))
	}

	testNode := NodeConfig{
		Name:             "node1",
		DownAfterNoAlive: 300,
		IdleConns:        16,

		User:     "root",
		Password: "root",

		Master: "127.0.0.1:3306",
		Slave:  "127.0.0.1:4306",
	}

	if !reflect.DeepEqual(cfg.Nodes[0], testNode) {
		fmt.Printf("%v\n", cfg.Nodes[0])
		t.Fatal("node1 must equal")
	}

	testNode_2 := NodeConfig{
		Name:   "node2",
		User:   "root",
		Master: "127.0.0.1:3307",
	}

	if !reflect.DeepEqual(cfg.Nodes[1], testNode_2) {
		t.Fatal("node2 must equal")
	}

	testShard_1 := ShardConfig{
		Table: "test_shard_hash",
		Key:   "id",
		Nodes: []string{"node1", "node2", "node3"},
		Type:  "hash",
	}
	if !reflect.DeepEqual(cfg.Schemas[0].RulesConifg.ShardRule[0], testShard_1) {
		t.Fatal("ShardConfig0 must equal")
	}

	testShard_2 := ShardConfig{
		Table: "test_shard_range",
		Key:   "id",
		Nodes: []string{"node2", "node3"},
		Type:  "range",
	}
	if !reflect.DeepEqual(cfg.Schemas[0].RulesConifg.ShardRule[1], testShard_2) {
		t.Fatal("ShardConfig1 must equal")
	}

	if 2 != len(cfg.Schemas[0].RulesConifg.ShardRule) {
		t.Fatal("ShardRule must 2")
	}

	testRules := RulesConfig{
		Default:   "node1",
		ShardRule: []ShardConfig{testShard_1, testShard_2},
	}
	if !reflect.DeepEqual(cfg.Schemas[0].RulesConifg, testRules) {
		t.Fatal("RulesConfig must equal")
	}

	testSchema := SchemaConfig{
		DB:          "kingshard",
		Nodes:       []string{"node1", "node2", "node3"},
		RulesConifg: testRules,
	}

	if !reflect.DeepEqual(cfg.Schemas[0], testSchema) {
		t.Fatal("schema must equal")
	}

	if cfg.LogLevel != "error" || cfg.User != "root" || cfg.Password != "root" || cfg.Addr != "127.0.0.1:4000" {
		t.Fatal("Top Config not equal.")
	}
}
