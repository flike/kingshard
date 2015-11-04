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
addr : 127.0.0.1:3601
user : root
password : 

nodes :
- 
    name : node1 
    down_after_noalive : 300
    idle_conns : 16
    user: root
    password:
    master : 127.0.0.1:3306
    slave : 

schemas :
-
    db : kingshard 
    nodes: [node1]
    rules:
        default: node1 
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
		testDB, _ = backend.Open("127.0.0.1:3601", "root", "", "kingshard", 100)
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
