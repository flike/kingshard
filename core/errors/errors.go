package errors

import (
	"errors"
)

var (
	ErrNoMasterConn  = errors.New("no master connection")
	ErrNoSlaveConn   = errors.New("no slave connection")
	ErrNoDefaultNode = errors.New("no default node")
	ErrNoMasterDb    = errors.New("no master database")
	ErrNoSlaveDb     = errors.New("no slave database")
	ErrNoDatabase    = errors.New("no database")

	ErrMasterDown = errors.New("master is down")
	ErrSlaveDown  = errors.New("slave is down")

	ErrAddressNull     = errors.New("address is nil")
	ErrInvalidArgument = errors.New("argument is invalid")
	ErrCmdUnsupport    = errors.New("command unsupport")

	ErrLocationsCount = errors.New("locations count are not equal")
	ErrNoCriteria     = errors.New("plan have no criteria")
	ErrNoRouteNode    = errors.New("no route node")
	ErrSelectInInsert = errors.New("select in insert not allowed")
	ErrInsertInMulti  = errors.New("insert in multi node")
	ErrUpdateInMulti  = errors.New("update in multi node")
	ErrDeleteInMulti  = errors.New("delete in multi node")
	ErrReplaceInMulti = errors.New("replace in multi node")
	ErrExecInMulti    = errors.New("exec in multi node")
	ErrTransInMulti   = errors.New("transaction in multi node")

	ErrNoPlan       = errors.New("statement have no plan")
	ErrUpdateKey    = errors.New("routing key in update expression")
	ErrStmtConvert  = errors.New("statement fail to convert")
	ErrExprConvert  = errors.New("expr fail to convert")
	ErrConnNotEqual = errors.New("the length of conns not equal sqls")
)
