package router

import (
	"strconv"
	"strings"
)

func StrValue(value interface{}) string {
	switch val := value.(type) {
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case string:
		return val
	case []byte:
		return string(val[:])
	}
	panic(NewKeyError("Unexpected key variable type %T", value))
}

type StrSepRangeShard struct {
	Seps []string
}

func (s *StrSepRangeShard) FindForKey(key interface{}) (int, error) {
	v := StrValue(key)
	for i, sep := range s.Seps  {
		switch strings.Compare(v, sep) {
		case 0:
			return i + 1, nil
		case -1:
			return i, nil
		}
	}
	return len(s.Seps), nil
}

type IntSepRangeShard struct {
	Seps []int64
}

func (s *IntSepRangeShard) FindForKey(key interface{}) (int, error) {
	v := NumValue(key)
	for i, sep := range s.Seps  {
		if (sep == v) {
			return i + 1, nil
		} else if (sep > v) {
			return i, nil
		}
	}
	return len(s.Seps), nil
}

func ParseIntSepSharding(Seps []string) ([]int64, error) {
	length := len(Seps)
	ranges := make([]int64, length)

	for i := 0; i < length; i++ {
		ranges[i] = NumValue(Seps[i])
	}
	return ranges, nil
}