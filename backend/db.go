package backend

import (
	"sync"
	"sync/atomic"

	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/mysql"
)

const (
	Up = iota
	Down
	Unknown

	DefaultMaxConnNum = 64
)

type DB struct {
	sync.Mutex

	addr       string
	user       string
	password   string
	db         string
	maxConnNum int
	state      int32

	idleConns chan *Conn

	connNum int
}

func Open(addr string, user string, password string, dbName string, maxConnNum int) *DB {
	db := new(DB)

	db.addr = addr
	db.user = user
	db.password = password
	db.db = dbName
	if 0 < maxConnNum {
		db.maxConnNum = maxConnNum
	} else {
		db.maxConnNum = DefaultMaxConnNum
	}

	db.idleConns = make(chan *Conn, maxConnNum)
	db.connNum = 0
	atomic.StoreInt32(&(db.state), Unknown)

	return db
}

func (db *DB) Addr() string {
	return db.addr
}

func (db *DB) State() string {
	var state string
	switch db.state {
	case Up:
		state = "up"
	case Down:
		state = "down"
	case Unknown:
		state = "unknow"
	}
	return state
}

func (db *DB) IdleConnCount() int {
	db.Lock()
	defer db.Unlock()
	return len(db.idleConns)
}

func (db *DB) Close() error {
	db.Lock()
	connChannel := db.idleConns
	db.idleConns = nil
	db.connNum = 0
	db.Unlock()
	if connChannel == nil {
		return nil
	}
	close(connChannel)
	for conn := range connChannel {
		conn.Close()
	}

	return nil
}

func (db *DB) Ping() error {
	c, err := db.PopConn()
	if err != nil {
		return err
	}

	err = c.Ping()
	db.PushConn(c, err)
	return err
}

func (db *DB) newConn() (*Conn, error) {
	co := new(Conn)

	if err := co.Connect(db.addr, db.user, db.password, db.db); err != nil {
		return nil, err
	}

	return co, nil
}

func (db *DB) tryReuse(co *Conn) error {
	if co.IsInTransaction() {
		//we can not reuse a connection in transaction status
		if err := co.Rollback(); err != nil {
			return err
		}
	}

	if !co.IsAutoCommit() {
		//we can not  reuse a connection not in autocomit
		if _, err := co.exec("set autocommit = 1"); err != nil {
			return err
		}
	}

	//connection may be set names early
	//we must use default utf8
	if co.GetCharset() != mysql.DEFAULT_CHARSET {
		if err := co.SetCharset(mysql.DEFAULT_CHARSET); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) PopConn() (*Conn, error) {
	var conn *Conn = nil
	var needNewConn bool

	db.Lock()
	connChannel := db.idleConns
	db.Unlock()
	if connChannel == nil {
		return nil, errors.ErrDatabaseClose
	}
	if 0 < len(connChannel) {
		conn = <-connChannel
	}

	if conn != nil {
		if err := conn.Ping(); err == nil {
			if err := db.tryReuse(conn); err == nil {
				//connection may alive
				return conn, nil
			}
		}
		conn.Close()
	}

	db.Lock()
	if db.connNum < db.maxConnNum {
		db.connNum++
		needNewConn = true
	}
	db.Unlock()

	if needNewConn {
		return db.newConn()
	}

	conn = <-connChannel
	if conn != nil {
		if err := conn.Ping(); err == nil {
			if err := db.tryReuse(conn); err == nil {
				//connection may alive
				return conn, nil
			}
		}
		conn.Close()
	}

	//conn is nil
	return nil, errors.ErrDatabaseClose
}

func (db *DB) PushConn(co *Conn, err error) {
	db.Lock()
	defer db.Unlock()

	if err != nil || db.idleConns == nil {
		co.Close()
		db.connNum--
		return
	}

	select {
	case db.idleConns <- co:
		break
	default:
		db.connNum--
		co.Close()
	}
}

type BackendConn struct {
	*Conn

	db *DB
}

func (p *BackendConn) Close() {
	if p.Conn != nil {
		p.db.PushConn(p.Conn, p.Conn.pkgErr)
		p.Conn = nil
	}
}

func (db *DB) GetConn() (*BackendConn, error) {
	c, err := db.PopConn()
	return &BackendConn{c, db}, err
}
