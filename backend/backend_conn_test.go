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
	"fmt"
	"testing"

	. "github.com/flike/kingshard/mysql"
)

func newTestConn() *Conn {
	c := new(Conn)

	if err := c.Connect("127.0.0.1:3306", "root", "", "kingshard"); err != nil {
		panic(err)
	}

	return c
}

func TestConn_Connect(t *testing.T) {
	c := newTestConn()
	defer c.Close()
}

func TestConn_Ping(t *testing.T) {
	c := newTestConn()
	defer c.Close()

	if err := c.Ping(); err != nil {
		t.Fatal(err)
	}
}

func TestConn_DeleteTable(t *testing.T) {
	c := newTestConn()
	defer c.Close()

	if _, err := c.Execute("drop table if exists kingshard_test_conn"); err != nil {
		t.Fatal(err)
	}
}

func TestConn_CreateTable(t *testing.T) {
	s := `CREATE TABLE IF NOT EXISTS kingshard_test_conn (
          id BIGINT(64) UNSIGNED  NOT NULL,
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          u tinyint unsigned,
          i tinyint,
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	c := newTestConn()
	defer c.Close()

	if _, err := c.Execute(s); err != nil {
		t.Fatal(err)
	}
}

func TestConn_Insert(t *testing.T) {
	s := `insert into kingshard_test_conn (id, str, f, e) values(1, "a", 3.14, "test1")`

	c := newTestConn()
	defer c.Close()

	if pkg, err := c.Execute(s); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}
}

func TestConn_Select(t *testing.T) {
	s := `select str, f, e from kingshard_test_conn where id = 1`

	c := newTestConn()
	defer c.Close()

	if result, err := c.Execute(s); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Fields) != 3 {
			t.Fatal(len(result.Fields))
		}

		if len(result.Values) != 1 {
			t.Fatal(len(result.Values))
		}

		if str, _ := result.GetString(0, 0); str != "a" {
			t.Fatal("invalid str", str)
		}

		if f, _ := result.GetFloat(0, 1); f != float64(3.14) {
			t.Fatal("invalid f", f)
		}

		if e, _ := result.GetString(0, 2); e != "test1" {
			t.Fatal("invalid e", e)
		}

		if str, _ := result.GetStringByName(0, "str"); str != "a" {
			t.Fatal("invalid str", str)
		}

		if f, _ := result.GetFloatByName(0, "f"); f != float64(3.14) {
			t.Fatal("invalid f", f)
		}

		if e, _ := result.GetStringByName(0, "e"); e != "test1" {
			t.Fatal("invalid e", e)
		}

	}
}

func TestConn_Escape(t *testing.T) {
	c := newTestConn()
	defer c.Close()

	e := `""''\abc`
	s := fmt.Sprintf(`insert into kingshard_test_conn (id, str) values(5, "%s")`,
		Escape(e))

	if _, err := c.Execute(s); err != nil {
		t.Fatal(err)
	}

	s = `select str from kingshard_test_conn where id = ?`

	if r, err := c.Execute(s, 5); err != nil {
		t.Fatal(err)
	} else {
		str, _ := r.GetString(0, 0)
		if str != e {
			t.Fatal(str)
		}
	}
}

func TestConn_SetCharset(t *testing.T) {
	c := newTestConn()
	defer c.Close()

	if err := c.SetCharset("gb2312", 24); err != nil {
		t.Fatal(err)
	}
}

func TestConn_SetAutoCommit(t *testing.T) {
	c := newTestConn()
	defer c.Close()

	if err := c.SetAutoCommit(0); err != nil {
		t.Fatal(err)
	}

	if err := c.SetAutoCommit(1); err != nil {
		t.Fatal(err)
	}
}
