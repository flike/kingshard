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
	"strconv"
	"strings"
	"time"

	"github.com/flike/kingshard/core/errors"
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
	return fmt.Sprintf("%d-%d", kr.Start, kr.End)
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

//return date of daynumber by order
//20151201-20151205
//20151201,20151202,20151203,20151204,20151205
func ParseDayRange(dateRange string) ([]int, error) {
	timeFormat := "20060102"
	dateDays := make([]int, 0)
	dateLength := 8

	dateTmp := strings.SplitN(dateRange, "-", 2)
	if len(dateTmp) == 1 {
		if len(dateTmp[0]) != dateLength {
			return nil, errors.ErrDateRangeIllegal
		}
		dateNum, err := strconv.Atoi(dateTmp[0])
		if err != nil {
			return nil, err
		}
		return []int{dateNum}, nil
	}
	if len(dateTmp) != 2 {
		return nil, errors.ErrDateRangeIllegal
	}
	if len(dateTmp[0]) != dateLength || len(dateTmp[1]) != dateLength {
		return nil, errors.ErrDateRangeIllegal
	}
	//change the begin day and the end day
	if dateTmp[1] < dateTmp[0] {
		dateTmp[0], dateTmp[1] = dateTmp[1], dateTmp[0]
	}

	begin, err := time.Parse(timeFormat, dateTmp[0])
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(timeFormat, dateTmp[1])
	if err != nil {
		return nil, err
	}

	duration := end.Sub(begin)
	daysCount := int(duration.Hours() / 24)

	for i := 0; i <= daysCount; i++ {
		date := begin.Add(time.Hour * time.Duration(24*i))
		dateStr := date.Format("20060102")
		dateNum, err := strconv.Atoi(dateStr)
		if err != nil {
			return nil, err
		}
		dateDays = append(dateDays, dateNum)
	}
	return dateDays, nil
}

//return date of month by order
//201510-201512
//201510,201511,201512
func ParseMonthRange(dateRange string) ([]int, error) {
	dateMonth := make([]int, 0)
	dateLength := 6

	dateTmp := strings.SplitN(dateRange, "-", 2)
	if len(dateTmp) == 1 {
		if len(dateTmp[0]) != dateLength {
			return nil, errors.ErrDateRangeIllegal
		}
		dateNum, err := strconv.Atoi(dateTmp[0])
		if err != nil {
			return nil, err
		}
		return []int{dateNum}, nil
	}
	if len(dateTmp) != 2 {
		return nil, errors.ErrDateRangeIllegal
	}
	if len(dateTmp[0]) != dateLength || len(dateTmp[1]) != dateLength {
		return nil, errors.ErrDateRangeIllegal
	}
	//change the begin month and the end month
	if dateTmp[1] < dateTmp[0] {
		dateTmp[0], dateTmp[1] = dateTmp[1], dateTmp[0]
	}

	beginYear, err := strconv.Atoi(dateTmp[0][:4])
	if err != nil {
		return nil, err
	}
	beginMonth, err := strconv.Atoi(dateTmp[0][4:])
	if err != nil {
		return nil, err
	}

	endYear, err := strconv.Atoi(dateTmp[1][:4])
	if err != nil {
		return nil, err
	}
	endMonth, err := strconv.Atoi(dateTmp[1][4:])
	if err != nil {
		return nil, err
	}
	// how many months between the two date
	monthCount := (endYear-beginYear)*12 + endMonth - beginMonth + 1
	monthTmp := beginMonth
	for i := 0; i < monthCount; i++ {
		if 12 < monthTmp {
			monthTmp = monthTmp % 12
			beginYear++
		}

		monthNum := beginYear*100 + monthTmp
		dateMonth = append(dateMonth, monthNum)
		monthTmp++
	}

	return dateMonth, nil
}

//return date of year by order
//2013-2015
//2013,2014,2015
func ParseYearRange(dateRange string) ([]int, error) {
	dateYear := make([]int, 0)
	dateLength := 4

	dateTmp := strings.SplitN(dateRange, "-", 2)
	if len(dateTmp) == 1 {
		if len(dateTmp[0]) != dateLength {
			return nil, errors.ErrDateRangeIllegal
		}
		dateNum, err := strconv.Atoi(dateTmp[0])
		if err != nil {
			return nil, err
		}
		return []int{dateNum}, nil
	}
	if len(dateTmp) != 2 {
		return nil, errors.ErrDateRangeIllegal
	}
	//change the begin year and the end year
	if dateTmp[1] < dateTmp[0] {
		dateTmp[0], dateTmp[1] = dateTmp[1], dateTmp[0]
	}
	beginYear, err := strconv.Atoi(dateTmp[0])
	if err != nil {
		return nil, err
	}
	endYear, err := strconv.Atoi(dateTmp[1])
	if err != nil {
		return nil, err
	}

	for i := beginYear; i <= endYear; i++ {
		dateYear = append(dateYear, i)
	}

	return dateYear, nil
}
