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
	"bytes"
	"fmt"
	"sort"

	"github.com/flike/kingshard/core/hack"
)

const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

type SortKey struct {
	//name of the field
	Name string

	Direction string

	//column index of the field
	column int
}

type resultsetSorter struct {
	*Resultset

	sk []SortKey
}

func newResultsetSorter(r *Resultset, sk []SortKey) (*resultsetSorter, error) {
	s := new(resultsetSorter)

	s.Resultset = r

	for i, k := range sk {
		if column, ok := r.FieldNames[k.Name]; ok {
			sk[i].column = column
		} else {
			return nil, fmt.Errorf("key %s not in resultset fields, can not sort", k.Name)
		}
	}

	s.sk = sk

	return s, nil
}

func (r *resultsetSorter) Len() int {
	return r.RowNumber()
}

func (r *resultsetSorter) Less(i, j int) bool {
	v1 := r.Values[i]
	v2 := r.Values[j]

	for _, k := range r.sk {
		v := cmpValue(v1[k.column], v2[k.column])

		if k.Direction == SortDesc {
			v = -v
		}

		if v < 0 {
			return true
		} else if v > 0 {
			return false
		}

		//equal, cmp next key
	}

	return false
}

//compare value using asc
func cmpValue(v1 interface{}, v2 interface{}) int {
	if v1 == nil && v2 == nil {
		return 0
	} else if v1 == nil {
		return -1
	} else if v2 == nil {
		return 1
	}

	switch v := v1.(type) {
	case string:
		s := v2.(string)
		return bytes.Compare(hack.Slice(v), hack.Slice(s))
	case []byte:
		s := v2.([]byte)
		return bytes.Compare(v, s)
	case int64:
		s := v2.(int64)
		if v < s {
			return -1
		} else if v > s {
			return 1
		} else {
			return 0
		}
	case uint64:
		s := v2.(uint64)
		if v < s {
			return -1
		} else if v > s {
			return 1
		} else {
			return 0
		}
	case float64:
		s := v2.(float64)
		if v < s {
			return -1
		} else if v > s {
			return 1
		} else {
			return 0
		}
	default:
		//can not go here
		panic(fmt.Sprintf("invalid type %T", v))
	}
}

func (r *resultsetSorter) Swap(i, j int) {
	r.Values[i], r.Values[j] = r.Values[j], r.Values[i]

	r.RowDatas[i], r.RowDatas[j] = r.RowDatas[j], r.RowDatas[i]
}

func (r *Resultset) Sort(sk []SortKey) error {
	s, err := newResultsetSorter(r, sk)

	if err != nil {
		return err
	}

	sort.Sort(s)

	return nil
}
