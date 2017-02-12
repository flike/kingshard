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
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/flike/kingshard/core/hack"
)

type RowData []byte

func (p RowData) Parse(f []*Field, binary bool) ([]interface{}, error) {
	if binary {
		return p.ParseBinary(f)
	} else {
		return p.ParseText(f)
	}
}

func (p RowData) ParseText(f []*Field) ([]interface{}, error) {
	data := make([]interface{}, len(f))

	var err error
	var v []byte
	var isNull, isUnsigned bool
	var pos int = 0
	var n int = 0

	for i := range f {
		v, isNull, n, err = LengthEnodedString(p[pos:])
		if err != nil {
			return nil, err
		}

		pos += n

		if isNull {
			data[i] = nil
		} else {
			isUnsigned = (f[i].Flag&UNSIGNED_FLAG > 0)
			switch f[i].Type {
			case MYSQL_TYPE_TINY, MYSQL_TYPE_SHORT, MYSQL_TYPE_LONG, MYSQL_TYPE_INT24,
				MYSQL_TYPE_LONGLONG, MYSQL_TYPE_YEAR:
				if isUnsigned {
					data[i], err = strconv.ParseUint(string(v), 10, 64)
				} else {
					data[i], err = strconv.ParseInt(string(v), 10, 64)
				}
			case MYSQL_TYPE_FLOAT, MYSQL_TYPE_DOUBLE, MYSQL_TYPE_NEWDECIMAL:
				data[i], err = strconv.ParseFloat(string(v), 64)
			case MYSQL_TYPE_VARCHAR, MYSQL_TYPE_VAR_STRING,
				MYSQL_TYPE_STRING, MYSQL_TYPE_DATETIME,
				MYSQL_TYPE_DATE, MYSQL_TYPE_TIME, MYSQL_TYPE_TIMESTAMP:
				data[i] = string(v)
			default:
				data[i] = v
			}

			if err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}

func (p RowData) ParseBinary(f []*Field) ([]interface{}, error) {
	data := make([]interface{}, len(f))

	if p[0] != OK_HEADER {
		return nil, ErrMalformPacket
	}

	pos := 1 + ((len(f) + 7 + 2) >> 3)

	nullBitmap := p[1:pos]

	var isUnsigned bool
	var isNull bool
	var n int
	var err error
	var v []byte
	for i := range data {
		if nullBitmap[(i+2)/8]&(1<<(uint(i+2)%8)) > 0 {
			data[i] = nil
			continue
		}

		isUnsigned = f[i].Flag&UNSIGNED_FLAG > 0

		switch f[i].Type {
		case MYSQL_TYPE_NULL:
			data[i] = nil
			continue

		case MYSQL_TYPE_TINY:
			if isUnsigned {
				data[i] = uint64(p[pos])
			} else {
				data[i] = int64(p[pos])
			}
			pos++
			continue

		case MYSQL_TYPE_SHORT, MYSQL_TYPE_YEAR:
			if isUnsigned {
				data[i] = uint64(binary.LittleEndian.Uint16(p[pos : pos+2]))
			} else {
				var n int16
				err = binary.Read(bytes.NewBuffer(p[pos:pos+2]), binary.LittleEndian, &n)
				if err != nil {
					return nil, err
				}
				data[i] = int64(n)
			}
			pos += 2
			continue

		case MYSQL_TYPE_INT24, MYSQL_TYPE_LONG:
			if isUnsigned {
				data[i] = uint64(binary.LittleEndian.Uint32(p[pos : pos+4]))
			} else {
				var n int32
				err = binary.Read(bytes.NewBuffer(p[pos:pos+4]), binary.LittleEndian, &n)
				if err != nil {
					return nil, err
				}
				data[i] = int64(n)
			}
			pos += 4
			continue

		case MYSQL_TYPE_LONGLONG:
			if isUnsigned {
				data[i] = binary.LittleEndian.Uint64(p[pos : pos+8])
			} else {
				var n int64
				err = binary.Read(bytes.NewBuffer(p[pos:pos+8]), binary.LittleEndian, &n)
				if err != nil {
					return nil, err
				}
				data[i] = int64(n)
			}
			pos += 8
			continue

		case MYSQL_TYPE_FLOAT:
			//data[i] = float64(math.Float32frombits(binary.LittleEndian.Uint32(p[pos : pos+4])))
			var n float32
			err = binary.Read(bytes.NewBuffer(p[pos:pos+4]), binary.LittleEndian, &n)
			if err != nil {
				return nil, err
			}
			data[i] = float64(n)
			pos += 4
			continue

		case MYSQL_TYPE_DOUBLE:
			var n float64
			err = binary.Read(bytes.NewBuffer(p[pos:pos+8]), binary.LittleEndian, &n)
			if err != nil {
				return nil, err
			}
			data[i] = n
			pos += 8
			continue

		case MYSQL_TYPE_DECIMAL, MYSQL_TYPE_NEWDECIMAL, MYSQL_TYPE_VARCHAR,
			MYSQL_TYPE_BIT, MYSQL_TYPE_ENUM, MYSQL_TYPE_SET, MYSQL_TYPE_TINY_BLOB,
			MYSQL_TYPE_MEDIUM_BLOB, MYSQL_TYPE_LONG_BLOB, MYSQL_TYPE_BLOB,
			MYSQL_TYPE_VAR_STRING, MYSQL_TYPE_STRING, MYSQL_TYPE_GEOMETRY:
			v, isNull, n, err = LengthEnodedString(p[pos:])
			pos += n
			if err != nil {
				return nil, err
			}

			if !isNull {
				data[i] = v
				continue
			} else {
				data[i] = nil
				continue
			}
		case MYSQL_TYPE_DATE, MYSQL_TYPE_NEWDATE:
			var num uint64
			num, isNull, n = LengthEncodedInt(p[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryDate(int(num), p[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		case MYSQL_TYPE_TIMESTAMP, MYSQL_TYPE_DATETIME:
			var num uint64
			num, isNull, n = LengthEncodedInt(p[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryDateTime(int(num), p[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		case MYSQL_TYPE_TIME:
			var num uint64
			num, isNull, n = LengthEncodedInt(p[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryTime(int(num), p[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("Stmt Unknown FieldType %d %s", f[i].Type, f[i].Name)
		}
	}

	return data, nil
}

type Result struct {
	Status uint16

	InsertId     uint64
	AffectedRows uint64

	*Resultset
}

type Resultset struct {
	Fields     []*Field
	FieldNames map[string]int
	Values     [][]interface{}

	RowDatas []RowData
}

func (r *Resultset) RowNumber() int {
	return len(r.Values)
}

func (r *Resultset) ColumnNumber() int {
	return len(r.Fields)
}

func (r *Resultset) GetValue(row, column int) (interface{}, error) {
	if row >= len(r.Values) || row < 0 {
		return nil, fmt.Errorf("invalid row index %d", row)
	}

	if column >= len(r.Fields) || column < 0 {
		return nil, fmt.Errorf("invalid column index %d", column)
	}

	return r.Values[row][column], nil
}

func (r *Resultset) NameIndex(name string) (int, error) {
	if column, ok := r.FieldNames[name]; ok {
		return column, nil
	} else {
		return 0, fmt.Errorf("invalid field name %s", name)
	}
}

func (r *Resultset) GetValueByName(row int, name string) (interface{}, error) {
	if column, err := r.NameIndex(name); err != nil {
		return nil, err
	} else {
		return r.GetValue(row, column)
	}
}

func (r *Resultset) IsNull(row, column int) (bool, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return false, err
	}

	return d == nil, nil
}

func (r *Resultset) IsNullByName(row int, name string) (bool, error) {
	if column, err := r.NameIndex(name); err != nil {
		return false, err
	} else {
		return r.IsNull(row, column)
	}
}

func (r *Resultset) GetUint(row, column int) (uint64, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case uint64:
		return v, nil
	case int64:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	case string:
		return strconv.ParseUint(v, 10, 64)
	case []byte:
		return strconv.ParseUint(string(v), 10, 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetUintByName(row int, name string) (uint64, error) {
	if column, err := r.NameIndex(name); err != nil {
		return 0, err
	} else {
		return r.GetUint(row, column)
	}
}

func (r *Resultset) GetIntByName(row int, name string) (int64, error) {
	if column, err := r.NameIndex(name); err != nil {
		return 0, err
	} else {
		return r.GetInt(row, column)
	}
}

func (r *Resultset) GetInt(row, column int) (int64, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case uint64:
		return int64(v), nil
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	case []byte:
		return strconv.ParseInt(string(v), 10, 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetFloat(row, column int) (float64, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case float64:
		return v, nil
	case uint64:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case []byte:
		return strconv.ParseFloat(string(v), 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetFloatByName(row int, name string) (float64, error) {
	if column, err := r.NameIndex(name); err != nil {
		return 0, err
	} else {
		return r.GetFloat(row, column)
	}
}

func (r *Resultset) GetString(row, column int) (string, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return "", err
	}

	switch v := d.(type) {
	case string:
		return v, nil
	case []byte:
		return hack.String(v), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetStringByName(row int, name string) (string, error) {
	if column, err := r.NameIndex(name); err != nil {
		return "", err
	} else {
		return r.GetString(row, column)
	}
}
