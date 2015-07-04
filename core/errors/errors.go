package errors

import (
	"errors"
)

var (
	ErrNoMasterConn  = errors.New("no master connection")
	ErrNoSlaveConn   = errors.New("no slave connection")
	ErrNoDefaultNode = errors.New("no default node")
)
