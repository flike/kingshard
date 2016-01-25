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
	"fmt"
	"math"
)

const (
	MinNumKey = math.MinInt64
	MaxNumKey = math.MaxInt64
)

//[start,end)
type NumKeyRange struct {
	Start int64
	End   int64
}

func (kr NumKeyRange) MapKey() string {
	return fmt.Sprintf("%d-%d", kr.String(), kr.End)
}

func (kr NumKeyRange) Contains(i int64) bool {
	return kr.Start <= i && (kr.End == MaxNumKey || i < kr.End)
}

func (kr NumKeyRange) String() string {
	return fmt.Sprintf("{Start: %d, End: %d}", kr.Start, kr.End)
}

func ParseNumSharding(Locations []int, TableRowLimit int) ([]NumKeyRange, error) {
	tableCount := 0
	length := len(Locations)

	for i := 0; i < length; i++ {
		tableCount += Locations[i]
	}

	ranges := make([]NumKeyRange, tableCount)
	for i := 0; i < tableCount; i++ {
		ranges[i].Start = int64(i * TableRowLimit)
		ranges[i].End = int64((i + 1) * TableRowLimit)
	}
	return ranges, nil
}
