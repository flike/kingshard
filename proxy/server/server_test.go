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

package server

import (
	"sync"
	"testing"
	"time"

	"github.com/flike/kingshard/backend"
	"github.com/flike/kingshard/config"
)

var testServerOnce sync.Once
var testServer *Server
var testDBOnce sync.Once
var testDB *backend.DB

var testConfigData = []byte(`
addr : 127.0.0.1:9696
user : root
password : 

nodes :
- 
    name : node1 
    down_after_noalive : 300
    idle_conns : 16
    user: root
    password: flike
    master : 127.0.0.1:3306
    slave : 

schema :
    default: node1  
    nodes: [node1]
    rules:
        shard:
            -
`)

func newTestServer(t *testing.T) *Server {
	f := func() {
		cfg, err := config.ParseConfigData(testConfigData)
		if err != nil {
			t.Fatal(err.Error())
		}

		testServer, err = NewServer(cfg)
		if err != nil {
			t.Fatal(err)
		}

		go testServer.Run()

		time.Sleep(1 * time.Second)
	}

	testServerOnce.Do(f)

	return testServer
}

func newTestDB(t *testing.T) *backend.DB {
	newTestServer(t)

	f := func() {
		testDB, _ = backend.Open("127.0.0.1:3306", "root", "flike", "kingshard", 100)
	}

	testDBOnce.Do(f)
	return testDB
}

func newTestDBConn(t *testing.T) *backend.BackendConn {
	db := newTestDB(t)

	c, err := db.GetConn()

	if err != nil {
		t.Fatal(err)
	}

	return c
}

func TestServer(t *testing.T) {
	newTestServer(t)
}
