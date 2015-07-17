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
