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
	"sort"
	"testing"
)

func testCheckList(t *testing.T, l []int, checkList ...int) {
	if len(l) != len(checkList) {
		t.Fatal("invalid list len", len(l), len(checkList))
	}

	for i := 0; i < len(l); i++ {
		if l[i] != checkList[i] {
			t.Fatal("invalid list item", l[i], i)
		}
	}
}

func TestListSet(t *testing.T) {
	var l1 []int
	var l2 []int
	var l3 []int

	l1 = []int{1, 2, 3}
	l2 = []int{2}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 3}
	l2 = []int{2, 3}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2, 3)

	l1 = []int{1, 2, 4}
	l2 = []int{2, 3}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 4}
	l2 = []int{}

	l3 = interList(l1, l2)
	testCheckList(t, l3)

	l1 = []int{1, 2, 3}
	l2 = []int{2}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3)

	l1 = []int{1, 2, 4}
	l2 = []int{3}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{2, 3, 4}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{2}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 3, 4)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{1, 3, 5}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 2, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{1, 3, 5, 6}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{2, 3}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 4)

	l1 = []int{1, 2, 2, 1, 5, 3, 5, 2}
	l2 = cleanList(l1)
	sort.Sort(sort.IntSlice(l2))
	testCheckList(t, l2, 1, 2, 3, 5)
}

func TestMakeLeList(t *testing.T) {
	l1 := []int{20150802, 20150812, 20150822, 20150823, 20150825, 20150828}
	l2 := makeLeList(20150822, l1)
	testCheckList(t, l2, 20150802, 20150812, 20150822)
	l3 := makeLeList(20150824, l1)
	testCheckList(t, l3, []int{}...)
}

func TestmakeLtList(t *testing.T) {
	l1 := []int{20150802, 20150812, 20150822, 20150823, 20150825, 20150828}
	l2 := makeLeList(20150822, l1)
	testCheckList(t, l2, 20150802, 20150812)
	l3 := makeLeList(20150824, l1)
	testCheckList(t, l3, []int{}...)
	l4 := makeLeList(20150802, l1)
	testCheckList(t, l4, []int{}...)
}

func TestmakeGeList(t *testing.T) {
	l1 := []int{20150802, 20150812, 20150822, 20150823, 20150825, 20150828}
	l2 := makeLeList(20150822, l1)
	testCheckList(t, l2, 20150822, 20150823, 20150825, 20150828)
	l3 := makeLeList(20150828, l1)
	testCheckList(t, l3, 20150828)
}

func TestmakeGtList(t *testing.T) {
	l1 := []int{20150802, 20150812, 20150822, 20150823, 20150825, 20150828}
	l2 := makeLeList(20150822, l1)
	testCheckList(t, l2, 20150823, 20150825, 20150828)
	l3 := makeLeList(20150824, l1)
	testCheckList(t, l3, []int{}...)
	l4 := makeLeList(20150828, l1)
	testCheckList(t, l4, []int{}...)
}

func TestParseYearRange(t *testing.T) {
	dateRange := "2014-2017"
	years, err := ParseYearRange(dateRange)
	if err != nil {
		t.Fatal(err)
	}
	testCheckList(t, years, 2014, 2015, 2016, 2017)

	dateRange = "2017-2013"
	years, err = ParseYearRange(dateRange)
	if err != nil {
		t.Fatal(err)
	}
	testCheckList(t, years, 2013, 2014, 2015, 2016, 2017)
}

func TestParseMonthRange(t *testing.T) {
	dateRange := "201602-201610"
	months, err := ParseMonthRange(dateRange)
	if err != nil {
		t.Fatal(err)
	}

	testCheckList(t, months,
		201602,
		201603,
		201604,
		201605,
		201606,
		201607,
		201608,
		201609,
		201610,
	)

	dateRange = "201603-201511"
	months, err = ParseMonthRange(dateRange)
	if err != nil {
		t.Fatal(err)
	}

	testCheckList(t, months,
		201511,
		201512,
		201601,
		201602,
		201603,
	)
}

func TestParseDayRange(t *testing.T) {
	dateRange := "20160227-20160304"
	days, err := ParseDayRange(dateRange)
	if err != nil {
		t.Fatal(err)
	}

	testCheckList(t, days,
		20160227,
		20160228,
		20160229,
		20160301,
		20160302,
		20160303,
		20160304,
	)

	dateRange = "20160304-20160301"
	days, err = ParseDayRange(dateRange)
	if err != nil {
		t.Fatal(err)
	}

	testCheckList(t, days,
		20160301,
		20160302,
		20160303,
		20160304,
	)
}
