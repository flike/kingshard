package server

import (
	"encoding/binary"
	"fmt"
	. "github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	. "github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/sqlparser"
	"math"
	"strconv"
	"strings"
)

var paramFieldData []byte
var columnFieldData []byte

func init() {
	var p = &Field{Name: []byte("?")}
	var c = &Field{}

	paramFieldData = p.Dump()
	columnFieldData = c.Dump()
}

type Stmt struct {
	id uint32

	params  int
	columns int

	args []interface{}

	s sqlparser.Statement

	sql string
}

func (s *Stmt) ResetParams() {
	s.args = make([]interface{}, s.params)
}

func (c *ClientConn) handleStmtPrepare(sql string) error {
	if c.schema == nil {
		return NewDefaultError(ER_NO_DB_ERROR)
	}

	s := new(Stmt)

	sql = strings.TrimRight(sql, ";")

	var err error
	s.s, err = sqlparser.Parse(sql)
	if err != nil {
		return fmt.Errorf(`parse sql "%s" error`, sql)
	}

	s.sql = sql

	//	var tableName string
	//	switch s := s.s.(type) {
	//	case *sqlparser.Select:
	//		tableName = nstring(s.From)
	//	case *sqlparser.Insert:
	//		tableName = nstring(s.Table)
	//	case *sqlparser.Update:
	//		tableName = nstring(s.Table)
	//	case *sqlparser.Delete:
	//		tableName = nstring(s.Table)
	//	case *sqlparser.Replace:
	//		tableName = nstring(s.Table)
	//	default:
	//		return fmt.Errorf(`unsupport prepare sql "%s"`, sql)
	//	}

	defaultRule := c.schema.rule.DefaultRule

	n := c.proxy.GetNode(defaultRule.Nodes[0])

	if co, err := n.GetMasterConn(); err != nil {
		return fmt.Errorf("prepare error %s", err)
	} else {
		defer co.Close()

		if err = co.UseDB(c.schema.db); err != nil {
			return fmt.Errorf("parepre error %s", err)
		}

		if t, err := co.Prepare(sql); err != nil {
			return fmt.Errorf("parepre error %s", err)
		} else {

			s.params = t.ParamNum()
			s.columns = t.ColumnNum()
		}
	}

	s.id = c.stmtId
	c.stmtId++

	if err = c.writePrepare(s); err != nil {
		return err
	}

	s.ResetParams()

	c.stmts[s.id] = s

	return nil
}

func (c *ClientConn) writePrepare(s *Stmt) error {
	data := make([]byte, 4, 128)

	//status ok
	data = append(data, 0)
	//stmt id
	data = append(data, Uint32ToBytes(s.id)...)
	//number columns
	data = append(data, Uint16ToBytes(uint16(s.columns))...)
	//number params
	data = append(data, Uint16ToBytes(uint16(s.params))...)
	//filter [00]
	data = append(data, 0)
	//warning count
	data = append(data, 0, 0)

	if err := c.writePacket(data); err != nil {
		return err
	}

	if s.params > 0 {
		for i := 0; i < s.params; i++ {
			data = data[0:4]
			data = append(data, []byte(paramFieldData)...)

			if err := c.writePacket(data); err != nil {
				return err
			}
		}

		if err := c.writeEOF(c.status); err != nil {
			return err
		}
	}

	if s.columns > 0 {
		for i := 0; i < s.columns; i++ {
			data = data[0:4]
			data = append(data, []byte(columnFieldData)...)

			if err := c.writePacket(data); err != nil {
				return err
			}
		}

		if err := c.writeEOF(c.status); err != nil {
			return err
		}

	}
	return nil
}

func (c *ClientConn) handleStmtExecute(data []byte) error {
	if len(data) < 9 {
		return ErrMalformPacket
	}

	pos := 0
	id := binary.LittleEndian.Uint32(data[0:4])
	pos += 4

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_execute")
	}

	flag := data[pos]
	pos++
	//now we only support CURSOR_TYPE_NO_CURSOR flag
	if flag != 0 {
		return NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("unsupported flag %d", flag))
	}

	//skip iteration-count, always 1
	pos += 4

	var nullBitmaps []byte
	var paramTypes []byte
	var paramValues []byte

	paramNum := s.params

	if paramNum > 0 {
		nullBitmapLen := (s.params + 7) >> 3
		if len(data) < (pos + nullBitmapLen + 1) {
			return ErrMalformPacket
		}
		nullBitmaps = data[pos : pos+nullBitmapLen]
		pos += nullBitmapLen

		//new param bound flag
		if data[pos] == 1 {
			pos++
			if len(data) < (pos + (paramNum << 1)) {
				return ErrMalformPacket
			}

			paramTypes = data[pos : pos+(paramNum<<1)]
			pos += (paramNum << 1)

			paramValues = data[pos:]
		}

		if err := c.bindStmtArgs(s, nullBitmaps, paramTypes, paramValues); err != nil {
			return err
		}
	}

	var err error

	switch stmt := s.s.(type) {
	case *sqlparser.Select:
		err = c.handlePrepareSelect(stmt, s.sql, s.args)
	case *sqlparser.Insert:
		err = c.handlePrepareExec(s.s, s.sql, s.args)
	case *sqlparser.Update:
		err = c.handlePrepareExec(s.s, s.sql, s.args)
	case *sqlparser.Delete:
		err = c.handlePrepareExec(s.s, s.sql, s.args)
	case *sqlparser.Replace:
		err = c.handlePrepareExec(s.s, s.sql, s.args)
	default:
		err = fmt.Errorf("command %T not supported now", stmt)
	}

	s.ResetParams()

	return err
}

func (c *ClientConn) handlePrepareSelect(stmt *sqlparser.Select, sql string, args []interface{}) error {
	defaultRule := c.schema.rule.DefaultRule
	if len(defaultRule.Nodes) == 0 {
		return ErrNoDefaultNode
	}
	defaultNode := c.proxy.GetNode(defaultRule.Nodes[0])

	//execute in Master DB
	conn, err := c.getBackendConn(defaultNode, false)
	if err != nil {
		return err
	}

	if conn == nil {
		r := c.newEmptyResultset(stmt)
		return c.writeResultset(c.status, r)
	}

	var rs []*Result
	rs, err = c.executeInNode(conn, sql, args)

	c.closeConn(conn, false)
	if err != nil {
		golog.Error("ClientConn", "handlePrepareSelect", err.Error(), c.connectionId)
		return err
	}

	err = c.mergeSelectResult(rs, stmt)
	if err != nil {
		golog.Error("ClientConn", "handlePrepareSelect", err.Error(), c.connectionId)
	}

	return err
}

func (c *ClientConn) handlePrepareExec(stmt sqlparser.Statement, sql string, args []interface{}) error {
	defaultRule := c.schema.rule.DefaultRule
	if len(defaultRule.Nodes) == 0 {
		return ErrNoDefaultNode
	}
	defaultNode := c.proxy.GetNode(defaultRule.Nodes[0])

	//execute in Master DB
	conn, err := c.getBackendConn(defaultNode, false)
	if err != nil {
		return err
	}

	if conn == nil {
		return c.writeOK(nil)
	}

	var rs []*Result
	rs, err = c.executeInNode(conn, sql, args)
	c.closeConn(conn, err != nil)

	if err == nil {
		err = c.mergeExecResult(rs)
	}

	return err
}

func (c *ClientConn) bindStmtArgs(s *Stmt, nullBitmap, paramTypes, paramValues []byte) error {
	args := s.args

	pos := 0

	var v []byte
	var n int = 0
	var isNull bool
	var err error

	for i := 0; i < s.params; i++ {
		if nullBitmap[i>>3]&(1<<(uint(i)%8)) > 0 {
			args[i] = nil
			continue
		}

		tp := paramTypes[i<<1]
		isUnsigned := (paramTypes[(i<<1)+1] & 0x80) > 0

		switch tp {
		case MYSQL_TYPE_NULL:
			args[i] = nil
			continue

		case MYSQL_TYPE_TINY:
			if len(paramValues) < (pos + 1) {
				return ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint8(paramValues[pos])
			} else {
				args[i] = int8(paramValues[pos])
			}

			pos++
			continue

		case MYSQL_TYPE_SHORT, MYSQL_TYPE_YEAR:
			if len(paramValues) < (pos + 2) {
				return ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint16(binary.LittleEndian.Uint16(paramValues[pos : pos+2]))
			} else {
				args[i] = int16((binary.LittleEndian.Uint16(paramValues[pos : pos+2])))
			}
			pos += 2
			continue

		case MYSQL_TYPE_INT24, MYSQL_TYPE_LONG:
			if len(paramValues) < (pos + 4) {
				return ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint32(binary.LittleEndian.Uint32(paramValues[pos : pos+4]))
			} else {
				args[i] = int32(binary.LittleEndian.Uint32(paramValues[pos : pos+4]))
			}
			pos += 4
			continue

		case MYSQL_TYPE_LONGLONG:
			if len(paramValues) < (pos + 8) {
				return ErrMalformPacket
			}

			if isUnsigned {
				args[i] = binary.LittleEndian.Uint64(paramValues[pos : pos+8])
			} else {
				args[i] = int64(binary.LittleEndian.Uint64(paramValues[pos : pos+8]))
			}
			pos += 8
			continue

		case MYSQL_TYPE_FLOAT:
			if len(paramValues) < (pos + 4) {
				return ErrMalformPacket
			}

			args[i] = float32(math.Float32frombits(binary.LittleEndian.Uint32(paramValues[pos : pos+4])))
			pos += 4
			continue

		case MYSQL_TYPE_DOUBLE:
			if len(paramValues) < (pos + 8) {
				return ErrMalformPacket
			}

			args[i] = math.Float64frombits(binary.LittleEndian.Uint64(paramValues[pos : pos+8]))
			pos += 8
			continue

		case MYSQL_TYPE_DECIMAL, MYSQL_TYPE_NEWDECIMAL, MYSQL_TYPE_VARCHAR,
			MYSQL_TYPE_BIT, MYSQL_TYPE_ENUM, MYSQL_TYPE_SET, MYSQL_TYPE_TINY_BLOB,
			MYSQL_TYPE_MEDIUM_BLOB, MYSQL_TYPE_LONG_BLOB, MYSQL_TYPE_BLOB,
			MYSQL_TYPE_VAR_STRING, MYSQL_TYPE_STRING, MYSQL_TYPE_GEOMETRY,
			MYSQL_TYPE_DATE, MYSQL_TYPE_NEWDATE,
			MYSQL_TYPE_TIMESTAMP, MYSQL_TYPE_DATETIME, MYSQL_TYPE_TIME:
			if len(paramValues) < (pos + 1) {
				return ErrMalformPacket
			}

			v, isNull, n, err = LengthEnodedString(paramValues[pos:])
			pos += n
			if err != nil {
				return err
			}

			if !isNull {
				args[i] = v
				continue
			} else {
				args[i] = nil
				continue
			}
		default:
			return fmt.Errorf("Stmt Unknown FieldType %d", tp)
		}
	}
	return nil
}

func (c *ClientConn) handleStmtSendLongData(data []byte) error {
	if len(data) < 6 {
		return ErrMalformPacket
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_send_longdata")
	}

	paramId := binary.LittleEndian.Uint16(data[4:6])
	if paramId >= uint16(s.params) {
		return NewDefaultError(ER_WRONG_ARGUMENTS, "stmt_send_longdata")
	}

	if s.args[paramId] == nil {
		s.args[paramId] = data[6:]
	} else {
		if b, ok := s.args[paramId].([]byte); ok {
			b = append(b, data[6:]...)
			s.args[paramId] = b
		} else {
			return fmt.Errorf("invalid param long data type %T", s.args[paramId])
		}
	}

	return nil
}

func (c *ClientConn) handleStmtReset(data []byte) error {
	if len(data) < 4 {
		return ErrMalformPacket
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_reset")
	}

	s.ResetParams()

	return c.writeOK(nil)
}

func (c *ClientConn) handleStmtClose(data []byte) error {
	if len(data) < 4 {
		return nil
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	delete(c.stmts, id)

	return nil
}
