//	Copyright (c) 2014-2015, Percona LLC and/or its affiliates. All rights reserved.
//	This program is free software: you can redistribute it and/or modify
//	it under the terms of the GNU Affero General Public License as published by
//	the Free Software Foundation, either version 3 of the License, or
//	(at your option) any later version.
//	This program is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even the implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU Affero General Public License for more details.
//	You should have received a copy of the GNU Affero General Public License
//	along with this program.  If not, see <http://www.gnu.org/licenses/>

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

import (
	"testing"
)

func TestFingerprintBasic(t *testing.T) {
	var q, fp string

	// A most basic case.
	q = "SELECT c FROM t WHERE id=1"
	fp = "select c from t where id=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// The values looks like one line -- comments, but they're not.
	q = `UPDATE groups_search SET  charter = '   -------3\'\' XXXXXXXXX.\n    \n    -----------------------------------------------------', show_in_list = 'Y' WHERE group_id='aaaaaaaa'`
	fp = "update groups_search set charter = ?, show_in_list = ? where group_id=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// PT treats this as "mysqldump", but we don't do any special fingerprints.
	q = "SELECT /*!40001 SQL_NO_CACHE */ * FROM `film`"
	fp = "select /*!? sql_no_cache */ * from `film`"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Fingerprints stored procedure calls specially
	q = "CALL foo(1, 2, 3)"
	fp = "call foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Fingerprints admin commands as themselves
	q = "administrator command: Init DB"
	fp = "administrator command: Init DB"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Removes identifier from USE
	q = "use `foo`"
	fp = "use ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Handles bug from perlmonks thread 728718
	q = "select null, 5.001, 5001. from foo"
	fp = "select ?, ?, ? from foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Handles quoted strings
	q = "select 'hello', '\nhello\n', \"hello\", '\\'' from foo"
	fp = "select ?, ?, ?, ? from foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Handles trailing newline
	q = "select 'hello'\n"
	fp = "select ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "select '\\\\' from foo"
	fp = "select ? from foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Collapses whitespace
	q = "select   foo"
	fp = "select foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Lowercases, replaces integer
	q = "SELECT * from foo where a = 5"
	fp = "select * from foo where a = ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Floats
	q = "select 0e0, +6e-30, -6.00 from foo where a = 5.5 or b=0.5 or c=.5"
	fp = "select ?, ?, ? from foo where a = ? or b=? or c=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Hex/bit
	q = "select 0x0, x'123', 0b1010, b'10101' from foo"
	fp = "select ?, ?, ?, ? from foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Collapses whitespace
	q = " select  * from\nfoo where a = 5"
	fp = "select * from foo where a = ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// IN lists
	q = "select * from foo where a in (5) and b in (5, 8,9 ,9 , 10)"
	fp = "select * from foo where a in(?+) and b in(?+)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}
	// Numeric table names.  By default, PT will return foo_?, etc. because
	// match_embedded_numbers is false by default for speed.
	q = "select foo_1 from foo_2_3"
	fp = "select foo_1 from foo_2_3"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Numeric table name prefixes
	q = "select 123foo from 123foo"
	fp = "select 123foo from 123foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Numeric table name prefixes with underscores
	q = "select 123_foo from 123_foo"
	fp = "select 123_foo from 123_foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// A string that needs no changes
	q = "insert into abtemp.coxed select foo.bar from foo"
	fp = "insert into abtemp.coxed select foo.bar from foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// limit alone
	q = "select * from foo limit 5"
	fp = "select * from foo limit ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// limit with comma-offset
	q = "select * from foo limit 5, 10"
	fp = "select * from foo limit ?, ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// limit with offset
	q = "select * from foo limit 5 offset 10"
	fp = "select * from foo limit ? offset ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Fingerprint LOAD DATA INFILE
	q = "LOAD DATA INFILE '/tmp/foo.txt' INTO db.tbl"
	fp = "load data infile ? into db.tbl"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Fingerprint db.tbl<number>name (preserve number)
	q = "SELECT * FROM prices.rt_5min where id=1"
	fp = "select * from prices.rt_5min where id=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Fingerprint /* -- comment */ SELECT (bug 1174956)
	q = "/* -- S++ SU ABORTABLE -- spd_user: rspadim */SELECT SQL_SMALL_RESULT SQL_CACHE DISTINCT centro_atividade FROM est_dia WHERE unidade_id=1001 AND item_id=67 AND item_id_red=573"
	fp = "select sql_small_result sql_cache distinct centro_atividade from est_dia where unidade_id=? and item_id=? and item_id_red=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "INSERT INTO t (ts) VALUES (NOW())"
	fp = "insert into t (ts) values(?+)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "INSERT INTO t (ts) VALUES ('()', '\\(', '\\)')"
	fp = "insert into t (ts) values(?+)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintValueList(t *testing.T) {
	var q, fp string

	// VALUES lists
	q = "insert into foo(a, b, c) values(2, 4, 5)"
	fp = "insert into foo(a, b, c) values(?+)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// VALUES lists with multiple ()
	q = "insert into foo(a, b, c) values(2, 4, 5) , (2,4,5)"
	fp = "insert into foo(a, b, c) values(?+)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// VALUES lists with VALUE()
	q = "insert into foo(a, b, c) value(2, 4, 5)"
	fp = "insert into foo(a, b, c) value(?+)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "insert into foo values (1, '(2)', 'This is a trick: ). More values.', 4)"
	fp = "insert into foo values(?+)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintInList(t *testing.T) {
	var q, fp string

	q = "select * from t where (base.nid IN  ('1412', '1410', '1411'))"
	fp = "select * from t where (base.nid in(?+))"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "SELECT ID, name, parent, type FROM posts WHERE _name IN ('perf','caching') AND (type = 'page' OR type = 'attachment')"
	fp = "select id, name, parent, type from posts where _name in(?+) and (type = ? or type = ?)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "SELECT t FROM field WHERE  (entity_type = 'node') AND (entity_id IN  ('609')) AND (language IN  ('und')) AND (deleted = '0') ORDER BY delta ASC"
	fp = "select t from field where (entity_type = ?) and (entity_id in(?+)) and (language in(?+)) and (deleted = ?) order by delta"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintOrderBy(t *testing.T) {
	var q, fp string

	// Remove ASC from ORDER BY
	// Issue 1030: Fingerprint can remove ORDER BY ASC
	q = "select c from t where i=1 order by c asc"
	fp = "select c from t where i=? order by c"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Remove only ASC from ORDER BY
	q = "select * from t where i=1 order by a, b ASC, d DESC, e asc"
	fp = "select * from t where i=? order by a, b, d desc, e"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Remove ASC from spacey ORDER BY
	q = `select * from t where i=1      order            by
			  a,  b          ASC, d    DESC,

									 e asc`
	fp = "select * from t where i=? order by a, b, d desc, e"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintOneLineComments(t *testing.T) {
	var q, fp string

	// Removes one-line comments in fingerprints
	q = "select \n-- bar\n foo"
	fp = "select foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Removes one-line comments in fingerprint without mushing things together
	q = "select foo-- bar\nfoo"
	fp = "select foo foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Removes one-line EOL comments in fingerprints
	q = "select foo -- bar\n"
	fp = "select foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintTricky(t *testing.T) {
	var q, fp string

	// Full hex can look like an ident if not for the leading 0x.
	q = "SELECT c FROM t WHERE id=0xdeadbeaf"
	fp = "select c from t where id=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Caused a crash.
	q = "SELECT *    FROM t WHERE 1=1 AND id=1"
	fp = "select * from t where ?=? and id=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Caused a crash.
	q = "SELECT `db`.*, (CASE WHEN (`date_start` <=  '2014-09-10 09:17:59' AND `date_end` >=  '2014-09-10 09:17:59') THEN 'open' WHEN (`date_start` >  '2014-09-10 09:17:59' AND `date_end` >  '2014-09-10 09:17:59') THEN 'tbd' ELSE 'none' END) AS `status` FROM `foo` AS `db` WHERE (a_b in ('1', '10101'))"
	fp = "select `db`.*, (case when (`date_start` <= ? and `date_end` >= ?) then ? when (`date_start` > ? and `date_end` > ?) then ? else ? end) as `status` from `foo` as `db` where (a_b in(?+))"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// VALUES() after ON DUPE KEY is not the same as VALUES() for INSERT.
	q = "insert into t values (1) on duplicate key update query_count=COALESCE(query_count, 0) + VALUES(query_count)"
	fp = "insert into t values(?+) on duplicate key update query_count=coalesce(query_count, ?) + values(query_count)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "insert into t values (1), (2), (3)\n\n\ton duplicate key update query_count=1"
	fp = "insert into t values(?+) on duplicate key update query_count=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "select  t.table_schema,t.table_name,engine  from information_schema.tables t  inner join information_schema.columns c  on t.table_schema=c.table_schema and t.table_name=c.table_name group by t.table_schema,t.table_name having  sum(if(column_key in ('PRI','UNI'),1,0))=0"
	fp = "select t.table_schema,t.table_name,engine from information_schema.tables t inner join information_schema.columns c on t.table_schema=c.table_schema and t.table_name=c.table_name group by t.table_schema,t.table_name having sum(if(column_key in(?+),?,?))=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// Empty value list is valid SQL.
	q = "INSERT INTO t () VALUES ()"
	fp = "insert into t () values()"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestNumbersInFunctions(t *testing.T) {
	var q, fp string

	// Full hex can look like an ident if not for the leading 0x.
	q = "select sleep(2) from test.n"
	fp = "select sleep(?) from test.n"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintPanicChallenge1(t *testing.T) {
	q := "SELECT '' '' ''"
	fp := "select ? ? ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "SELECT '' '' '' FROM kamil"
	fp = "select ? ? ? from kamil"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintPanicChallenge2(t *testing.T) {
	q := "SELECT 'a' 'b' 'c' 'd'"
	fp := "select ? ? ? ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "SELECT 'a' 'b' 'c' 'd' FROM kamil"
	fp = "select ? ? ? ? from kamil"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintKeywords(t *testing.T) {
	var q, fp string

	// values is a keyword but value is not.
	q = "SELECT name, value FROM variable"
	fp = "select name, value from variable"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintUseIndex(t *testing.T) {
	var q, fp string

	q = `SELECT 	1 AS one FROM calls USE INDEX(index_name)`
	fp = "select ? as one from calls use index(index_name)"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

}

func TestFingerprintWithNumberInDbName(t *testing.T) {
	var q, fp string
	defaultReplaceNumbersInWords := ReplaceNumbersInWords
	ReplaceNumbersInWords = true
	defer func() {
		// Restore default value for other tests
		ReplaceNumbersInWords = defaultReplaceNumbersInWords
	}()

	q = "SELECT c FROM org235.t WHERE id=0xdeadbeaf"
	fp = "select c from org?.t where id=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "CREATE DATABASE org235_percona345 COLLATE 'utf8_general_ci'"
	fp = "create database org?_percona? collate ?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "select foo_1 from foo_2_3"
	fp = "select foo_? from foo_?_?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	q = "SELECT * FROM prices.rt_5min where id=1"
	fp = "select * from prices.rt_?min where id=?"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}

	// @todo prefixes are not supported, requires more hacks
	q = "select 123foo from 123foo"
	fp = "select 123foo from 123foo"
	if GetFingerprint(q) != fp {
		t.Fatalf("query=%s,and fingerPrint=%s\n", q, fp)
	}
}
