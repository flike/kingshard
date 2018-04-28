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

package errors

import (
	"errors"
)

var (
	ErrNoMasterConn  = errors.New("no master connection")
	ErrNoSlaveConn   = errors.New("no slave connection")
	ErrNoDefaultNode = errors.New("no default node")
	ErrNoMasterDB    = errors.New("no master database")
	ErrNoSlaveDB     = errors.New("no slave database")
	ErrNoDatabase    = errors.New("no database")

	ErrMasterDown    = errors.New("master is down")
	ErrSlaveDown     = errors.New("slave is down")
	ErrDatabaseClose = errors.New("database is close")
	ErrConnIsNil     = errors.New("connection is nil")
	ErrBadConn       = errors.New("connection was bad")
	ErrIgnoreSQL     = errors.New("ignore this sql")

	ErrAddressNull     = errors.New("address is nil")
	ErrInvalidArgument = errors.New("argument is invalid")
	ErrInvalidCharset  = errors.New("charset is invalid")
	ErrCmdUnsupport    = errors.New("command unsupport")

	ErrLocationsCount = errors.New("locations count is not equal")
	ErrNoCriteria     = errors.New("plan have no criteria")
	ErrNoRouteNode    = errors.New("no route node")
	ErrResultNil      = errors.New("result is nil")
	ErrSumColumnType  = errors.New("sum column type error")
	ErrSelectInInsert = errors.New("select in insert not allowed")
	ErrInsertInMulti  = errors.New("insert in multi node")
	ErrUpdateInMulti  = errors.New("update in multi node")
	ErrDeleteInMulti  = errors.New("delete in multi node")
	ErrReplaceInMulti = errors.New("replace in multi node")
	ErrExecInMulti    = errors.New("exec in multi node")
	ErrTransInMulti   = errors.New("transaction in multi node")

	ErrNoPlan           = errors.New("statement have no plan")
	ErrNoPlanRule       = errors.New("statement have no plan rule")
	ErrUpdateKey        = errors.New("routing key in update expression")
	ErrStmtConvert      = errors.New("statement fail to convert")
	ErrExprConvert      = errors.New("expr fail to convert")
	ErrConnNotEqual     = errors.New("the length of conns not equal sqls")
	ErrKeyOutOfRange    = errors.New("shard key not in key range")
	ErrMultiShard       = errors.New("insert or replace has multiple shard targets")
	ErrIRNoColumns      = errors.New("insert or replace must specify columns")
	ErrIRNoShardingKey  = errors.New("insert or replace not contain sharding key")
	ErrColsLenNotMatch  = errors.New("insert or replace cols and values length not match")
	ErrDateIllegal      = errors.New("date format illegal")
	ErrDateRangeIllegal = errors.New("date range format illegal")
	ErrDateRangeCount   = errors.New("date range count is not equal")
	ErrSlaveExist       = errors.New("slave has exist")
	ErrSlaveNotExist    = errors.New("slave has not exist")
	ErrBlackSqlExist    = errors.New("black sql has exist")
	ErrBlackSqlNotExist = errors.New("black sql has not exist")
	ErrInsertTooComplex = errors.New("insert is too complex")
	ErrSQLNULL          = errors.New("sql is null")

	ErrInternalServer   = errors.New("internal server error")
)
