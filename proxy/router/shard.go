// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

package router

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"strconv"
	"time"

	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/hack"
)

/*由分片ID找到分片，可用文件中的函数*/
type KeyError string

func NewKeyError(format string, args ...interface{}) KeyError {
	return KeyError(fmt.Sprintf(format, args...))
}

func (ke KeyError) Error() string {
	return string(ke)
}

func handleError(err *error) {
	if x := recover(); x != nil {
		*err = x.(KeyError)
	}
}

// Uint64Key is a uint64 that can be converted into a KeyspaceId.
type Uint64Key uint64

func (i Uint64Key) String() string {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint64(i))
	return buf.String()
}

func EncodeValue(value interface{}) string {
	switch val := value.(type) {
	case int:
		return Uint64Key(val).String()
	case uint64:
		return Uint64Key(val).String()
	case int64:
		return Uint64Key(val).String()
	case string:
		return val
	case []byte:
		return hack.String(val)
	}
	panic(NewKeyError("Unexpected key variable type %T", value))
}

func HashValue(value interface{}) uint64 {
	switch val := value.(type) {
	case int:
		return uint64(val)
	case uint64:
		return uint64(val)
	case int64:
		return uint64(val)
	case string:
		if v, err := strconv.ParseUint(val, 10, 64); err != nil {
			return uint64(crc32.ChecksumIEEE(hack.Slice(val)))
		} else {
			return uint64(v)
		}
	case []byte:
		return uint64(crc32.ChecksumIEEE(val))
	}
	panic(NewKeyError("Unexpected key variable type %T", value))
}

func NumValue(value interface{}) int64 {
	switch val := value.(type) {
	case int:
		return int64(val)
	case uint64:
		return int64(val)
	case int64:
		return int64(val)
	case string:
		if v, err := strconv.ParseInt(val, 10, 64); err != nil {
			panic(NewKeyError("invalid num format %s", v))
		} else {
			return v
		}
	case []byte:
		if v, err := strconv.ParseInt(hack.String(val), 10, 64); err != nil {
			panic(NewKeyError("invalid num format %s", v))
		} else {
			return v
		}
	}
	panic(NewKeyError("Unexpected key variable type %T", value))
}

type Shard interface {
	FindForKey(key interface{}) (int, error)
}

/*一个范围的分片,例如[start,end)*/
type RangeShard interface {
	Shard
	EqualStart(key interface{}, index int) bool
	EqualStop(key interface{}, index int) bool
}

type HashShard struct {
	ShardNum int
}

func (s *HashShard) FindForKey(key interface{}) (int, error) {
	h := HashValue(key)

	return int(h % uint64(s.ShardNum)), nil
}

type NumRangeShard struct {
	Shards []NumKeyRange
}

func (s *NumRangeShard) FindForKey(key interface{}) (int, error) {
	v := NumValue(key)
	for i, r := range s.Shards {
		if r.Contains(v) {
			return i, nil
		}
	}
	return -1, errors.ErrKeyOutOfRange
}

func (s *NumRangeShard) EqualStart(key interface{}, index int) bool {
	v := NumValue(key)
	return s.Shards[index].Start == v
}
func (s *NumRangeShard) EqualStop(key interface{}, index int) bool {
	v := NumValue(key)
	return s.Shards[index].End == v
}

type DateYearShard struct {
}

//the format of date is: YYYY-MM-DD HH:MM:SS,YYYY-MM-DD or unix timestamp(int)
func (s *DateYearShard) FindForKey(key interface{}) (int, error) {
	switch val := key.(type) {
	case int:
		tm := time.Unix(int64(val), 0)
		return tm.Year(), nil
	case uint64:
		tm := time.Unix(int64(val), 0)
		return tm.Year(), nil
	case int64:
		tm := time.Unix(val, 0)
		return tm.Year(), nil
	case string:
		if v, err := strconv.Atoi(val[:4]); err != nil {
			panic(NewKeyError("invalid num format %s", v))
		} else {
			return v, nil
		}
	}
	panic(NewKeyError("Unexpected key variable type %T", key))
}

type DateMonthShard struct {
}

//the format of date is: YYYY-MM-DD HH:MM:SS,YYYY-MM-DD or unix timestamp(int)
func (s *DateMonthShard) FindForKey(key interface{}) (int, error) {
	timeFormat := "2006-01-02"
	switch val := key.(type) {
	case int:
		tm := time.Unix(int64(val), 0)
		dateStr := tm.Format(timeFormat)
		s := dateStr[:4] + dateStr[5:7]
		yearMonth, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return yearMonth, nil
	case uint64:
		tm := time.Unix(int64(val), 0)
		dateStr := tm.Format(timeFormat)
		s := dateStr[:4] + dateStr[5:7]
		yearMonth, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return yearMonth, nil
	case int64:
		tm := time.Unix(val, 0)
		dateStr := tm.Format(timeFormat)
		s := dateStr[:4] + dateStr[5:7]
		yearMonth, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return yearMonth, nil
	case string:
		if len(val) < len(timeFormat) {
			return 0, fmt.Errorf("invalid date format %s", val)
		}
		s := val[:4] + val[5:7]
		if v, err := strconv.Atoi(s); err != nil {
			return 0, fmt.Errorf("invalid date format %s", val)
		} else {
			return v, nil
		}
	}
	panic(NewKeyError("Unexpected key variable type %T", key))
}

type DateDayShard struct {
}

//the format of date is: YYYY-MM-DD HH:MM:SS,YYYY-MM-DD or unix timestamp(int)
func (s *DateDayShard) FindForKey(key interface{}) (int, error) {
	timeFormat := "2006-01-02"
	switch val := key.(type) {
	case int:
		tm := time.Unix(int64(val), 0)
		dateStr := tm.Format(timeFormat)
		s := dateStr[:4] + dateStr[5:7] + dateStr[8:10]
		yearMonthDay, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return yearMonthDay, nil
	case uint64:
		tm := time.Unix(int64(val), 0)
		dateStr := tm.Format(timeFormat)
		s := dateStr[:4] + dateStr[5:7] + dateStr[8:10]
		yearMonthDay, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return yearMonthDay, nil
	case int64:
		tm := time.Unix(val, 0)
		dateStr := tm.Format(timeFormat)
		s := dateStr[:4] + dateStr[5:7] + dateStr[8:10]
		yearMonthDay, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return yearMonthDay, nil
	case string:
		if len(val) < len(timeFormat) {
			return 0, fmt.Errorf("invalid date format %s", val)
		}
		s := val[:4] + val[5:7] + val[8:10]
		if v, err := strconv.Atoi(s); err != nil {
			return 0, fmt.Errorf("invalid date format %s", val)
		} else {
			return v, nil
		}
	}
	panic(NewKeyError("Unexpected key variable type %T", key))
}

type DefaultShard struct {
}

func (s *DefaultShard) FindForKey(key interface{}) (int, error) {
	return 0, nil
}
