package server

import (
	"github.com/flike/kingshard/backend"
	"github.com/flike/kingshard/mysql"
)

func (c *ClientConn) isInTransaction() bool {
	return c.status&mysql.SERVER_STATUS_IN_TRANS > 0 ||
		!c.isAutoCommit()
}

func (c *ClientConn) isAutoCommit() bool {
	return c.status&mysql.SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *ClientConn) handleBegin() error {
	c.status |= mysql.SERVER_STATUS_IN_TRANS
	return c.writeOK(nil)
}

func (c *ClientConn) handleCommit() (err error) {
	if err := c.commit(); err != nil {
		return err
	} else {
		return c.writeOK(nil)
	}
}

func (c *ClientConn) handleRollback() (err error) {
	if err := c.rollback(); err != nil {
		return err
	} else {
		return c.writeOK(nil)
	}
}

func (c *ClientConn) commit() (err error) {
	c.status &= ^mysql.SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Commit(); e != nil {
			err = e
		}
		co.Close()
	}
	c.Lock()
	c.txConns = make(map[*backend.Node]*backend.BackendConn)
	c.Unlock()
	return
}

func (c *ClientConn) rollback() (err error) {
	c.status &= ^mysql.SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Rollback(); e != nil {
			err = e
		}
		co.Close()
	}
	c.Lock()
	c.txConns = make(map[*backend.Node]*backend.BackendConn)
	c.Unlock()
	return
}
