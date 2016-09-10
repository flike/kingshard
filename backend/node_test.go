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

package backend

import (
	"testing"

	"github.com/flike/kingshard/config"
)

func TestParse(t *testing.T) {
	node := new(Node)
	nodeConfig := config.NodeConfig{
		Name:             "node1",
		DownAfterNoAlive: 100,
		MaxConnNum:       16,
		User:             "root",
		Password:         "",
		Master:           "127.0.0.1:3306",
		Slave:            "127.0.0.1:3306@2,127.0.0.1:3307@4,127.0.0.1:3308@8",
	}
	node.Cfg = nodeConfig
	err := node.ParseMaster(nodeConfig.Master)
	if err != nil {
		t.Fatal(err.Error())
	}
	if node.Master.addr != "127.0.0.1:3306" {
		t.Fatal(node.Master)
	}
	err = node.ParseSlave(nodeConfig.Slave)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%v\n", node.RoundRobinQ)
	t.Logf("%v\n", node.SlaveWeights)
	t.Logf("%v\n", node.Master)
	t.Logf("%v\n", node.Slave)
}
