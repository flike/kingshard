// Copyright 2015 The kingshard Authors. All rights reserved.
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
	"sync/atomic"
)

type Counter struct {
	// 上一秒请求总数
	oldOpCnt int64
	// 上一秒错误数
	oldErrorCnt int64
	// 上一秒慢查询数
	oldSlowCnt int64
	// 连接数
	connCnt int64
	// 请求总数
	opCnt int64
	// 错误数
	errorCnt int64
	// 慢查询数
	slowCnt int64
}

func NewCounter() *Counter {
	return &Counter{connCnt: 0, opCnt: 0, errorCnt: 0, slowCnt: 0}
}

func (counter *Counter) IncrConnCnt() {
	atomic.AddInt64(&counter.connCnt, 1)
}

func (counter *Counter) DecrConnCnt() {
	atomic.AddInt64(&counter.connCnt, -1)
}

func (counter *Counter) IncrOpCnt() {
	atomic.AddInt64(&counter.opCnt, 1)
}

func (counter *Counter) IncrErrorCnt() {
	atomic.AddInt64(&counter.errorCnt, 1)
}

func (counter *Counter) IncrSlowCnt() {
	atomic.AddInt64(&counter.slowCnt, 1)
}

//每间隔一秒重置数据
func (counter *Counter) FlushCounter() {
	atomic.StoreInt64(&counter.oldOpCnt, counter.opCnt)
	atomic.StoreInt64(&counter.oldErrorCnt, counter.errorCnt)
	atomic.StoreInt64(&counter.oldSlowCnt, counter.slowCnt)

	atomic.StoreInt64(&counter.opCnt, 0)
	atomic.StoreInt64(&counter.errorCnt, 0)
	atomic.StoreInt64(&counter.slowCnt, 0)
}
