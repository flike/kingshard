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

// Package query provides functions to transform queries.
package mysql

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

const (
	unknown             byte = iota
	inWord                   // \S+
	inNumber                 // [0-9a-fA-Fx.-]
	inSpace                  // space, tab, \r, \n
	inOp                     // [=<>!] (usually precedes a number)
	opOrNumber               // + in 2 + 2 or +3e-9
	inQuote                  // '...' or "..."
	subOrOLC                 // - or start of -- comment
	inDash                   // -- begins a one-line comment if followed by space
	inOLC                    // -- comment (at least one space after dash is required)
	divOrMLC                 // / operator or start of /* comment */
	mlcOrMySQLCode           // /* comment */ or /*! MySQL-specific code */
	inMLC                    // /* comment */
	inValues                 // VALUES (1), ..., (N)
	moreValuesOrUnknown      // , (2nd+) or ON DUPLICATE KEY or end of query
	orderBy                  // ORDER BY
	onDupeKeyUpdate          // ON DUPLICATE KEY UPDATE
	inNumberInWord           // e.g. db23
)

var stateName map[byte]string = map[byte]string{
	0:  "unknown",
	1:  "inWord",
	2:  "inNumber",
	3:  "inSpace",
	4:  "inOp",
	5:  "opOrNumber",
	6:  "inQuote",
	7:  "subOrOLC",
	8:  "inDash",
	9:  "inOLC",
	10: "divOrMLC",
	11: "mlcOrMySQLCode",
	12: "inMLC",
	13: "inValues",
	14: "moreValuesOrUnknown",
	15: "orderBy",
	16: "onDupeKeyUpdate",
	17: "inNumberInWord",
}

// Debug prints very verbose tracing information to STDOUT.
var Debug bool = false

// ReplaceNumbersInWords enables replacing numbers in words. For example:
// `SELECT c FROM org235.t` -> `SELECT c FROM org?.t`. For more examples
// look at test query_test.go/TestFingerprintWithNumberInDbName.
var ReplaceNumbersInWords = false

// Fingerprint returns the canonical form of q. The primary transformations are:
//   - Replace values with ?
//   - Collapse whitespace
//   - Remove comments
//   - Lowercase everything
// Additional trasnformations are performed which change the syntax of the
// original query without affecting its performance characteristics. For
// example, "ORDER BY col ASC" is the same as "ORDER BY col", so "ASC" in the
// fingerprint is removed.
func GetFingerprint(q string) string {
	q += " " // need range to run off end of original query
	prevWord := ""
	f := make([]byte, len(q))
	fi := 0
	pr := rune(0) // previous rune
	s := unknown  // current state
	sqlState := unknown
	quoteChar := rune(0)
	cpFromOffset := 0
	cpToOffset := 0
	addSpace := false
	escape := false
	parOpen := 0
	parOpenTotal := 0
	valueNo := 0
	firstPar := 0

	for qi, r := range q {
		if Debug {
			fmt.Printf("\n%d:%d %s/%s [%d:%d] %x %q\n", qi, fi, stateName[s], stateName[sqlState], cpFromOffset, cpToOffset, r, r)
		}

		//1. Skip parts of the query for certain states.

		if s == inQuote {
			// We're in a 'quoted value' or "quoted value".  The quoted value
			// ends at the first non-escaped matching quote character (' or ").
			if r != quoteChar {
				// The only char inside a quoted value we need to track is \,
				// the escape char.  This allows us to tell that the 2nd ' in
				// '\'' is escaped, not the ending quote char.
				if escape {
					if Debug {
						fmt.Println("Ignore quoted literal")
					}
					escape = false
				} else if r == '\\' {
					if Debug {
						fmt.Println("Escape")
					}
					escape = true
				} else {
					if Debug {
						fmt.Println("Ignore quoted value")
					}
				}
			} else if escape {
				// \' or \"
				if Debug {
					fmt.Println("Quote literal")
				}
				escape = false
			} else {
				// 'foo' -> ?
				// "foo" -> ?
				if Debug {
					fmt.Println("Quote end")
				}
				escape = false

				// qi = the closing quote char, so +1 to ensure we don't copy
				// anything before this, i.e. quoted value is done, move on.
				cpFromOffset = qi + 1

				if sqlState == inValues {
					// ('Hello world!', ...) -> VALUES (, ...)
					// The inValues state uses this state to skip quoted values,
					// so we don't replace them with ?; the inValues blocks will
					// replace the entire value list with ?+.
					s = inValues
				} else {
					f[fi] = '?'
					fi++
					s = unknown
				}
			}
			continue
		} else if s == inNumberInWord {
			// Replaces number in words with ?
			// e.g. `db37` to `db?`
			// Parser can fall into inNumberInWord only if
			// option ReplaceNumbersInWords is turned on
			if r >= '0' && r <= '9' {
				if Debug {
					fmt.Println("Ignore digit in word")
				}
				continue
			}
			// 123 -> ?, 0xff -> ?, 1e-9 -> ?, etc.
			if Debug {
				fmt.Println("Number in word end")
			}
			f[fi] = '?'
			fi++
			cpFromOffset = qi
			if isSpace(r) {
				s = unknown
			} else {
				s = inWord
			}
		} else if s == inNumber {
			// We're in a number which can be something simple like 123 or
			// something trickier like 1e-9 or 0xFF.  The pathological case is
			// like 12ff: this is valid hex number and a valid ident (e.g. table
			// name).  We can't detect this; the best we can do is realize that
			// 12ffz is not a number because of the z.
			if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F') || r == '.' || r == 'x' || r == '-' {
				if Debug {
					fmt.Println("Ignore digit")
				}
				continue
			}
			if (r >= 'g' && r <= 'z') || (r >= 'G' && r <= 'Z') || r == '_' {
				if Debug {
					fmt.Println("Not a number")
				}
				cpToOffset = qi
				s = inWord
			} else {
				// 123 -> ?, 0xff -> ?, 1e-9 -> ?, etc.
				if Debug {
					fmt.Println("Number end")
				}
				f[fi] = '?'
				fi++
				cpFromOffset = qi
				cpToOffset = qi
				s = unknown
			}
		} else if s == inValues {
			// We're in the (val1),...,(valN) after IN or VALUE[S].  A single
			// () value ends when the parenthesis are balanced, but...
			if r == ')' {
				parOpen--
				parOpenTotal++
				if Debug {
					fmt.Println("Close parenthesis", parOpen)
				}
			} else if r == '(' {
				parOpen++
				if Debug {
					fmt.Println("Open parenthesis", parOpen)
				}
				if parOpen == 1 {
					firstPar = qi
				}
			} else if r == '\'' || r == '"' {
				// VALUES ('Hello world!') -> enter inQuote state to skip
				// the quoted value so ')' in 'This ) is a trick' doesn't
				// balance an outer parenthesis.
				if Debug {
					fmt.Println("Quote begin")
				}
				s = inQuote
				quoteChar = r
				continue
			} else if isSpace(r) {
				if Debug {
					fmt.Println("Space")
				}
				continue
			}
			if parOpen > 0 {
				// Parenthesis are not balanced yet; i.e. haven't reached
				// closing ) for this value.
				continue
			}
			if parOpenTotal == 0 {
				// SELECT value FROM t
				if Debug {
					fmt.Println("Literal values not VALUES()")
				}
				s = inWord
				continue
			}
			// (<anything>) -> (?+) only for first value
			if Debug {
				fmt.Println("Values end")
			}
			valueNo++
			if valueNo == 1 {
				if qi-firstPar > 1 {
					copy(f[fi:fi+4], "(?+)")
					fi += 4
				} else {
					// INSERT INTO t VALUES ()
					copy(f[fi:fi+2], "()")
					fi += 2
				}
				firstPar = 0
			}
			// ... the difficult part is that there may be other values, e.g.
			// (1), (2), (3).  So we enter the following state.  The values list
			// ends when the next char is not a comma.
			s = moreValuesOrUnknown
			pr = r
			cpFromOffset = qi + 1
			parOpenTotal = 0
			continue
		} else if s == inMLC {
			// We're in a /* mutli-line comments */.  Skip and ignore it all.
			if pr == '*' && r == '/' {
				// /* foo */ -> (nothing)
				if Debug {
					fmt.Println("Multi-line comment end")
				}
				s = unknown
			} else {
				if Debug {
					fmt.Println("Ignore multi-line comment content")
				}
			}
			continue
		} else if s == mlcOrMySQLCode {
			// We're at the start of either a /* multi-line comment */ or some
			// /*![version] some MySQL-specific code */.  The ! after the /*
			// determines which one.
			if r != '!' {
				if Debug {
					fmt.Println("Multi-line comment")
				}
				s = inMLC
				continue
			} else {
				// /*![version] SQL_NO_CACHE */ -> /*![version] SQL_NO_CACHE */ (no change)
				if Debug {
					fmt.Println("MySQL-specific code")
				}
				s = inWord
			}
		} else if s == inOLC {
			// We're in a -- one line comment.  A space after -- is required.
			// It ends at the end of the line, but there can be more query after
			// it like:
			//   SELECT * -- comment
			//   FROM t
			// is really "SELECT * FROM t".
			if r == 0x0A { // newline
				if Debug {
					fmt.Println("One-line comment end")
				}
				s = unknown
			}
			continue
		} else if isSpace(r) && isSpace(pr) {
			// All space is collapsed into a single space, so if this char is
			// a space and the previous was too, then skip the extra space.
			if Debug {
				fmt.Println("Skip space")
			}
			// +1 here ensures we actually skip the extra space in certain
			// cases like "select \n-- bar\n foo".  When a part of the query
			// triggers a copy of preceding chars, if the only preceding char
			// is a space then it's incorrectly copied, but +1 sets cpFromOffset
			// to the same offset as the trigger char, thus avoiding the copy.
			// For example in that ^ query, the offsets are:
			//   0 's'
			//   1 'e'
			//   2 'l'
			//   3 'e'
			//   4 'c'
			//   5 't'
			//   6 ' '
			//   7 '\n'
			//   8 '-'
			// After copying 'select ', we are here @ 7 and intend to skip the
			// newline.  Next, the '-' @ 8 triggers a copy of any preceding
			// chars.  So here if we set cpFromOffset = 7 then 7:8 is copied,
			// the newline, but setting cpFromOffset = 7 + 1 is 8:8 and so
			// nothing is copied as we want.  Actually, cpToOffset is still 6
			// in this case, but 8:6 avoids the copy too.
			cpFromOffset = qi + 1
			pr = r
			continue
		}

		//2. Change state based on rune and current state.

		switch {
		case r >= 0x30 && r <= 0x39: // 0-9
			switch s {
			case opOrNumber:
				if Debug {
					fmt.Println("+/-First digit")
				}
				cpToOffset = qi - 1
				s = inNumber
			case inOp:
				if Debug {
					fmt.Println("First digit after operator")
				}
				cpToOffset = qi
				s = inNumber
			case inWord:
				if pr == '(' {
					if Debug {
						fmt.Println("Number in function")
					}
					cpToOffset = qi
					s = inNumber
				} else if pr == ',' {
					// foo,4 -- 4 may be a number literal or a word/ident
					if Debug {
						fmt.Println("Number or word")
					}
					s = inNumber
					cpToOffset = qi
				} else {
					if Debug {
						fmt.Println("Number in word")
					}
					if ReplaceNumbersInWords {
						s = inNumberInWord
						cpToOffset = qi
					}
				}
			default:
				if Debug {
					fmt.Println("Number literal")
				}
				s = inNumber
				cpToOffset = qi
			}
		case isSpace(r):
			if s == unknown {
				if Debug {
					fmt.Println("Lost in space")
				}
				if fi > 0 && !isSpace(rune(f[fi-1])) {
					if Debug {
						fmt.Println("Add space")
					}
					f[fi] = ' '
					fi++
					// This is a common case: a space after skipping something,
					// e.g. col = 'foo'<space>. We want only the first space,
					// so advance cpFromOffset to whatever is after the space
					// and if it's more space then space skipping block will
					// handle it.
					cpFromOffset = qi + 1
				}
			} else if s == inDash {
				if Debug {
					fmt.Println("One-line comment begin")
				}
				s = inOLC
				if cpToOffset > 2 {
					cpToOffset = qi - 2
					addSpace = true
				}
			} else if s == moreValuesOrUnknown {
				if Debug {
					fmt.Println("Space after values")
				}
				if valueNo == 1 {
					f[fi] = ' '
					fi++
				}
			} else {
				if Debug {
					fmt.Println("Word end")
				}
				word := strings.ToLower(q[cpFromOffset:qi])
				// Only match USE if it is the first word in the query, otherwise,
				// it could be a USE INDEX
				if word == "use" && prevWord == "" {
					return "use ?"
				} else if (word == "null" && (prevWord != "is" && prevWord != "not")) || word == "null," {
					if Debug {
						fmt.Println("NULL as value")
					}
					f[fi] = '?'
					fi++
					if word[len(word)-1] == ',' {
						f[fi] = ','
						fi++
					}
					f[fi] = ' '
					fi++
					cpFromOffset = qi + 1
				} else if prevWord == "order" && word == "by" {
					if Debug {
						fmt.Println("ORDER BY begin")
					}
					sqlState = orderBy
				} else if sqlState == orderBy && wordIn(word, "asc", "asc,", "asc ") {
					if Debug {
						fmt.Println("ORDER BY ASC")
					}
					cpFromOffset = qi
					if word[len(word)-1] == ',' {
						fi--
						f[fi] = ','
						f[fi+1] = ' '
						fi += 2
					}
				} else if prevWord == "key" && word == "update" {
					if Debug {
						fmt.Println("ON DUPLICATE KEY UPDATE begin")
					}
					sqlState = onDupeKeyUpdate
				}
				s = inSpace
				cpToOffset = qi
				addSpace = true
			}
		case r == '\'' || r == '"':
			if pr != '\\' {
				if s != inQuote {
					if Debug {
						fmt.Println("Quote begin")
					}
					s = inQuote
					quoteChar = r
					cpToOffset = qi
					if pr == 'x' || pr == 'b' {
						if Debug {
							fmt.Println("Hex/binary value")
						}
						// We're at the first quote char of x'0F'
						// (or b'0101', etc.), so -2 for the quote char and
						// the x or b char to copy anything before and up to
						// this value.
						cpToOffset = -2
					}
				}
			}
		case r == '=' || r == '<' || r == '>' || r == '!':
			if Debug {
				fmt.Println("Operator")
			}
			if s != inWord && s != inOp {
				cpFromOffset = qi
			}
			s = inOp
		case r == '/':
			if Debug {
				fmt.Println("Op or multi-line comment")
			}
			s = divOrMLC
		case r == '*' && s == divOrMLC:
			if Debug {
				fmt.Println("Multi-line comment or MySQL-specific code")
			}
			s = mlcOrMySQLCode
		case r == '+':
			if Debug {
				fmt.Println("Operator or number")
			}
			s = opOrNumber
		case r == '-':
			if pr == '-' {
				if Debug {
					fmt.Println("Dash")
				}
				s = inDash
			} else {
				if Debug {
					fmt.Println("Operator or number")
				}
				s = opOrNumber
			}
		case r == '.':
			if s == inNumber || s == inOp {
				if Debug {
					fmt.Println("Floating point number")
				}
				s = inNumber
				cpToOffset = qi
			}
		case r == '(':
			if prevWord == "call" {
				// 'CALL foo(...)' -> 'call foo'
				if Debug {
					fmt.Println("CALL sp_name")
				}
				return "call " + q[cpFromOffset:qi]
			} else if sqlState != onDupeKeyUpdate && (((s == inSpace || s == moreValuesOrUnknown) && (prevWord == "value" || prevWord == "values" || prevWord == "in")) || wordIn(q[cpFromOffset:qi], "value", "values", "in")) {
				// VALUE(, VALUE (, VALUES(, VALUES (, IN(, or IN(
				// but not after ON DUPLICATE KEY UPDATE
				if Debug {
					fmt.Println("Values begin")
				}
				s = inValues
				sqlState = inValues
				parOpen = 1
				firstPar = qi
				if valueNo == 0 {
					cpToOffset = qi
				}
			} else if s != inWord {
				if Debug {
					fmt.Println("Random (")
				}
				valueNo = 0
				cpFromOffset = qi
				s = inWord
			}
		case r == ',' && s == moreValuesOrUnknown:
			if Debug {
				fmt.Println("More values")
			}
		case r == ':' && prevWord == "administrator":
			// 'administrator command: Init DB' -> 'administrator command: Init DB' (no change)
			if Debug {
				fmt.Println("Admin cmd")
			}
			return q[0 : len(q)-1] // original query minus the trailing space we added
		case r == '#':
			if Debug {
				fmt.Println("One-line comment begin")
			}
			s = inOLC
		default:
			if s != inWord && s != inOp {
				// If in a word or operator then keep copying the query, else
				// previous chars were being ignored for some reasons but now
				// we should start copying again, so set cpFromOffset.  Example:
				// col=NOW(). 'col' will be set to copy, but then '=' will put
				// us in inOp state which, if a value follows, will trigger a
				// copy of "col=", but "NOW()" is not a value so "N" is caught
				// here and since s=inOp still we do not copy yet (this block is
				// is not entered).
				if Debug {
					fmt.Println("Random character")
				}
				valueNo = 0
				cpFromOffset = qi

				if sqlState == inValues {
					// Values are comma-separated, so the first random char
					// marks the end of the VALUE() or IN() list.
					if Debug {
						fmt.Println("No more values")
					}
					sqlState = unknown
				}
			}
			s = inWord
		}

		/**
		 * 3. Copy a slice of the query into the fingerprint.
		 */

		if cpToOffset > cpFromOffset {
			l := cpToOffset - cpFromOffset
			prevWord = strings.ToLower(q[cpFromOffset:cpToOffset])
			if Debug {
				fmt.Printf("copy '%s' (%d:%d, %d:%d) %d\n", prevWord, fi, fi+l, cpFromOffset, cpToOffset, l)
			}
			copy(f[fi:fi+l], prevWord)
			fi += l
			cpFromOffset = cpToOffset
			if wordIn(prevWord, "in", "value", "values") && sqlState != onDupeKeyUpdate {
				// IN ()     -> in(?+)
				// VALUES () -> values(?+)
				addSpace = false
				s = inValues
				sqlState = inValues
			} else if addSpace {
				if Debug {
					fmt.Println("Add space")
				}
				f[fi] = ' '
				fi++
				cpFromOffset++
				addSpace = false
			}
		}
		pr = r
	}

	// Remove trailing spaces.
	for fi > 0 && isSpace(rune(f[fi-1])) {
		fi--
	}

	// Return the fingerprint.
	return string(f[0:fi])
}

func isSpace(r rune) bool {
	return r == 0x20 || r == 0x09 || r == 0x0D || r == 0x0A
}

func wordIn(q string, words ...string) bool {
	q = strings.ToLower(q)
	for _, word := range words {
		if q == word {
			return true
		}
	}
	return false
}

func Id(fingerprint string) string {
	id := md5.New()
	io.WriteString(id, fingerprint)
	h := fmt.Sprintf("%x", id.Sum(nil))
	return strings.ToUpper(h[16:32])
}

// GetMd5 returns the MD5 checksum of fingerprint.
func GetMd5(fingerPrint string) string {
	data := []byte(fingerPrint)
	return fmt.Sprintf("%x", md5.Sum(data))
}
