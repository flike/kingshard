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

package mysql

const (
	MinProtocolVersion byte   = 10
	MaxPayloadLen      int    = 1<<24 - 1
	TimeFormat         string = "2006-01-02 15:04:05"
	ServerVersion      string = "5.6.20-kingshard"
)

const (
	OK_HEADER          byte = 0x00
	ERR_HEADER         byte = 0xff
	EOF_HEADER         byte = 0xfe
	LocalInFile_HEADER byte = 0xfb
)

const (
	SERVER_STATUS_IN_TRANS             uint16 = 0x0001
	SERVER_STATUS_AUTOCOMMIT           uint16 = 0x0002
	SERVER_MORE_RESULTS_EXISTS         uint16 = 0x0008
	SERVER_STATUS_NO_GOOD_INDEX_USED   uint16 = 0x0010
	SERVER_STATUS_NO_INDEX_USED        uint16 = 0x0020
	SERVER_STATUS_CURSOR_EXISTS        uint16 = 0x0040
	SERVER_STATUS_LAST_ROW_SEND        uint16 = 0x0080
	SERVER_STATUS_DB_DROPPED           uint16 = 0x0100
	SERVER_STATUS_NO_BACKSLASH_ESCAPED uint16 = 0x0200
	SERVER_STATUS_METADATA_CHANGED     uint16 = 0x0400
	SERVER_QUERY_WAS_SLOW              uint16 = 0x0800
	SERVER_PS_OUT_PARAMS               uint16 = 0x1000
)

const (
	COM_SLEEP byte = iota
	COM_QUIT
	COM_INIT_DB
	COM_QUERY
	COM_FIELD_LIST
	COM_CREATE_DB
	COM_DROP_DB
	COM_REFRESH
	COM_SHUTDOWN
	COM_STATISTICS
	COM_PROCESS_INFO
	COM_CONNECT
	COM_PROCESS_KILL
	COM_DEBUG
	COM_PING
	COM_TIME
	COM_DELAYED_INSERT
	COM_CHANGE_USER
	COM_BINLOG_DUMP
	COM_TABLE_DUMP
	COM_CONNECT_OUT
	COM_REGISTER_SLAVE
	COM_STMT_PREPARE
	COM_STMT_EXECUTE
	COM_STMT_SEND_LONG_DATA
	COM_STMT_CLOSE
	COM_STMT_RESET
	COM_SET_OPTION
	COM_STMT_FETCH
	COM_DAEMON
	COM_BINLOG_DUMP_GTID
	COM_RESET_CONNECTION
)

const (
	CLIENT_LONG_PASSWORD uint32 = 1 << iota
	CLIENT_FOUND_ROWS
	CLIENT_LONG_FLAG
	CLIENT_CONNECT_WITH_DB
	CLIENT_NO_SCHEMA
	CLIENT_COMPRESS
	CLIENT_ODBC
	CLIENT_LOCAL_FILES
	CLIENT_IGNORE_SPACE
	CLIENT_PROTOCOL_41
	CLIENT_INTERACTIVE
	CLIENT_SSL
	CLIENT_IGNORE_SIGPIPE
	CLIENT_TRANSACTIONS
	CLIENT_RESERVED
	CLIENT_SECURE_CONNECTION
	CLIENT_MULTI_STATEMENTS
	CLIENT_MULTI_RESULTS
	CLIENT_PS_MULTI_RESULTS
	CLIENT_PLUGIN_AUTH
	CLIENT_CONNECT_ATTRS
	CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA
)

//https://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnType
const (
	MYSQL_TYPE_DECIMAL byte = iota
	MYSQL_TYPE_TINY
	MYSQL_TYPE_SHORT
	MYSQL_TYPE_LONG
	MYSQL_TYPE_FLOAT
	MYSQL_TYPE_DOUBLE
	MYSQL_TYPE_NULL
	MYSQL_TYPE_TIMESTAMP
	MYSQL_TYPE_LONGLONG
	MYSQL_TYPE_INT24
	MYSQL_TYPE_DATE
	MYSQL_TYPE_TIME
	MYSQL_TYPE_DATETIME
	MYSQL_TYPE_YEAR
	MYSQL_TYPE_NEWDATE
	MYSQL_TYPE_VARCHAR
	MYSQL_TYPE_BIT
)

const (
	MYSQL_TYPE_NEWDECIMAL byte = iota + 0xf6
	MYSQL_TYPE_ENUM
	MYSQL_TYPE_SET
	MYSQL_TYPE_TINY_BLOB
	MYSQL_TYPE_MEDIUM_BLOB
	MYSQL_TYPE_LONG_BLOB
	MYSQL_TYPE_BLOB
	MYSQL_TYPE_VAR_STRING
	MYSQL_TYPE_STRING
	MYSQL_TYPE_GEOMETRY
)

const (
	NOT_NULL_FLAG       = 1
	PRI_KEY_FLAG        = 2
	UNIQUE_KEY_FLAG     = 4
	BLOB_FLAG           = 16
	UNSIGNED_FLAG       = 32
	ZEROFILL_FLAG       = 64
	BINARY_FLAG         = 128
	ENUM_FLAG           = 256
	AUTO_INCREMENT_FLAG = 512
	TIMESTAMP_FLAG      = 1024
	SET_FLAG            = 2048
	NUM_FLAG            = 32768
	PART_KEY_FLAG       = 16384
	GROUP_FLAG          = 32768
	UNIQUE_FLAG         = 65536
)

const (
	AUTH_NAME = "mysql_native_password"
)

var (
	TK_ID_INSERT   = 1
	TK_ID_UPDATE   = 2
	TK_ID_DELETE   = 3
	TK_ID_REPLACE  = 4
	TK_ID_SET      = 5
	TK_ID_BEGIN    = 6
	TK_ID_COMMIT   = 7
	TK_ID_ROLLBACK = 8
	TK_ID_ADMIN    = 9
	TK_ID_USE      = 10

	TK_ID_SELECT      = 11
	TK_ID_START       = 12
	TK_ID_TRANSACTION = 13
	TK_ID_SHOW        = 14
	TK_ID_TRUNCATE    = 15

	PARSE_TOKEN_MAP = map[string]int{
		"insert":      TK_ID_INSERT,
		"update":      TK_ID_UPDATE,
		"delete":      TK_ID_DELETE,
		"replace":     TK_ID_REPLACE,
		"set":         TK_ID_SET,
		"begin":       TK_ID_BEGIN,
		"commit":      TK_ID_COMMIT,
		"rollback":    TK_ID_ROLLBACK,
		"admin":       TK_ID_ADMIN,
		"select":      TK_ID_SELECT,
		"use":         TK_ID_USE,
		"start":       TK_ID_START,
		"transaction": TK_ID_TRANSACTION,
		"show":        TK_ID_SHOW,
		"truncate":    TK_ID_TRUNCATE,
	}
	// '*'
	COMMENT_PREFIX uint8 = 42
	COMMENT_STRING       = "*"

	//
	TK_STR_SELECT = "select"
	TK_STR_FROM   = "from"
	TK_STR_INTO   = "into"
	TK_STR_SET    = "set"

	TK_STR_TRANSACTION    = "transaction"
	TK_STR_LAST_INSERT_ID = "last_insert_id()"
	TK_STR_MASTER_HINT    = "*master*"
	//show
	TK_STR_COLUMNS = "columns"
	TK_STR_FIELDS  = "fields"

	SET_KEY_WORDS = map[string]struct{}{
		"names": struct{}{},

		"character_set_results":           struct{}{},
		"@@character_set_results":         struct{}{},
		"@@session.character_set_results": struct{}{},

		"character_set_client":           struct{}{},
		"@@character_set_client":         struct{}{},
		"@@session.character_set_client": struct{}{},

		"character_set_connection":           struct{}{},
		"@@character_set_connection":         struct{}{},
		"@@session.character_set_connection": struct{}{},

		"autocommit":           struct{}{},
		"@@autocommit":         struct{}{},
		"@@session.autocommit": struct{}{},
	}
)
