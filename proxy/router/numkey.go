package router

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	MinNumKey = math.MinInt64
	MaxNumKey = math.MaxInt64
)

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

// ParseShardingSpec parses a string that describes a sharding
// specification. a-b-c-d will be parsed as a-b, b-c, c-d. The empty
// string may serve both as the start and end of the keyspace: -a-b-
// will be parsed as start-a, a-b, b-end.
func ParseNumShardingSpec(spec string) ([]NumKeyRange, error) {
	parts := strings.Split(spec, "-")
	if len(parts) == 1 {
		return nil, fmt.Errorf("malformed spec: doesn't define a range: %q", spec)
	}
	var old int64
	var err error
	if len(parts[0]) != 0 {
		old, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, err
		}
	} else {
		old = MinNumKey
	}

	ranges := make([]NumKeyRange, len(parts)-1)

	var n int64

	for i, p := range parts[1:] {
		if p == "" && i != (len(parts)-2) {
			return nil, fmt.Errorf("malformed spec: MinKey/MaxKey cannot be in the middle of the spec: %q", spec)
		}

		if p != "" {
			n, err = strconv.ParseInt(p, 10, 64)
			if err != nil {
				return nil, err
			}
		} else {
			n = MaxNumKey
		}
		if n <= old {
			return nil, fmt.Errorf("malformed spec: shard limits should be in order: %q", spec)
		}

		ranges[i] = NumKeyRange{Start: old, End: n}
		old = n
	}
	return ranges, nil
}
