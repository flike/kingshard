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
)

func TestStmt_DropTable(t *testing.T) {
	str := `drop table if exists kingshard_test_stmt`

	c := newTestConn()

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestStmt_CreateTable(t *testing.T) {
	str := `CREATE TABLE IF NOT EXISTS kingshard_test_stmt (
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

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err = s.Execute(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestStmt_Delete(t *testing.T) {
	str := `delete from kingshard_test_stmt`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestStmt_Insert(t *testing.T) {
	str := `insert into kingshard_test_stmt (id, str, f, e, u, i) values (?, ?, ?, ?, ?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Execute(1, "a", 3.14, "test1", 255, -127); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()
}

func TestStmt_Select(t *testing.T) {
	str := `select str, f, e from kingshard_test_stmt where id = ?`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if result, err := s.Execute(1); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Values) != 1 {
			t.Fatal(len(result.Values))
		}

		if len(result.Fields) != 3 {
			t.Fatal(len(result.Fields))
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

	s.Close()
}

func TestStmt_NULL(t *testing.T) {
	str := `insert into kingshard_test_stmt (id, str, f, e) values (?, ?, ?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Execute(2, nil, 3.14, nil); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()

	str = `select * from kingshard_test_stmt where id = ?`
	s, err = c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if r, err := s.Execute(2); err != nil {
		t.Fatal(err)
	} else {
		if b, err := r.IsNullByName(0, "id"); err != nil {
			t.Fatal(err)
		} else if b == true {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "str"); err != nil {
			t.Fatal(err)
		} else if b == false {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "f"); err != nil {
			t.Fatal(err)
		} else if b == true {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "e"); err != nil {
			t.Fatal(err)
		} else if b == false {
			t.Fatal(b)
		}
	}

	s.Close()
}

func TestStmt_Unsigned(t *testing.T) {
	str := `insert into kingshard_test_stmt (id, u) values (?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Execute(3, uint8(255)); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()

	str = `select u from kingshard_test_stmt where id = ?`

	s, err = c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if r, err := s.Execute(3); err != nil {
		t.Fatal(err)
	} else {
		if u, err := r.GetUint(0, 0); err != nil {
			t.Fatal(err)
		} else if u != uint64(255) {
			t.Fatal(u)
		}
	}

	s.Close()
}

func TestStmt_Signed(t *testing.T) {
	str := `insert into kingshard_test_stmt (id, i) values (?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(4, 127); err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(uint64(18446744073709551516), int8(-128)); err != nil {
		t.Fatal(err)
	}

	s.Close()

}

func TestStmt_Trans(t *testing.T) {
	c := newTestConn()
	defer c.Close()

	if _, err := c.Execute(`insert into kingshard_test_stmt (id, str) values (1002, "abc")`); err != nil {
		t.Fatal(err)
	}

	if err := c.Begin(); err != nil {
		t.Fatal(err)
	}

	str := `select str from kingshard_test_stmt where id = ?`

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(1002); err != nil {
		t.Fatal(err)
	}

	if err := c.Commit(); err != nil {
		t.Fatal(err)
	}

	if r, err := s.Execute(1002); err != nil {
		t.Fatal(err)
	} else {
		if str, _ := r.GetString(0, 0); str != `abc` {
			t.Fatal(str)
		}
	}

	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
}
