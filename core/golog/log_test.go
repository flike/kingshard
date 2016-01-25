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

package golog

import (
	"os"
	"testing"
)

const (
	ModMySQL = "mysql"
)

func TestLog(t *testing.T) {
	SetLevel(LevelTrace)
	Error(ModMySQL, "sqlx.DB.Select", "test", 0, "select * from qing_user")
	Info(ModMySQL, "sqlx.DB.Select", "test", 0, "select * from qing_user")
	Trace(ModMySQL, "sqlx.DB.Select", "test", 0, "select * from qing_user")
	Debug(ModMySQL, "sqlx.DB.Select", "test", 0, "select * from qing_user")
	Warn(ModMySQL, "sqlx.DB.Select", "test", 0, "select * from qing_user")
	Fatal(ModMySQL, "sqlx.DB.Select", "test", 0, "a", "b", "c", "d", "e")

	Fatal(ModMySQL, "AA", "test", 0, "%3d", 123)
}

func TestEscape(t *testing.T) {
	r := escape("abc= %|", false)
	if r != "abc= %%" {
		t.Fatal("invalid result ", r)
	}

	r = escape("abc= %|", true)
	if r != "abc %%" {
		t.Fatal("invalid result ", r)
	}

	if r := escape("%3d", false); r != "%%3d" {
		t.Fatal("invalid result ", r)
	}
}

func TestRotatingFileLog(t *testing.T) {
	path := "/tmp/test_log"
	os.RemoveAll(path)

	os.Mkdir(path, 0777)
	fileName := path + "/test"

	h, err := NewRotatingFileHandler(fileName, 1024*1024, 2)
	if err != nil {
		t.Fatal(err)
	}

	GlobalSysLogger = New(h, Lfile|Ltime|Llevel)
	GlobalSysLogger.SetLevel(LevelTrace)
	Debug("log", "hello,world", "OK", 0, "fileName", fileName, "fileName2", fileName, "fileName3", fileName)

	GlobalSysLogger.Close()

	//os.RemoveAll(path)
}
