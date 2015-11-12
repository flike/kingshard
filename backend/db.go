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

	InitConnCount     = 16
	DefaultMaxConnNum = 1024
)

type DB struct {
	sync.Mutex

	addr     string
	user     string
	password string
	db       string
	state    int32

	maxConnNum  int
	InitConnNum int
	idleConns   chan *Conn
	cacheConns  chan *Conn
	checkConn   *Conn
}

func Open(addr string, user string, password string, dbName string, maxConnNum int) (*DB, error) {
	var err error
	db := new(DB)
	db.addr = addr
	db.user = user
	db.password = password
	db.db = dbName

	if 0 < maxConnNum {
		db.maxConnNum = maxConnNum
		if db.maxConnNum < 16 {
			db.InitConnNum = db.maxConnNum
		} else {
			db.InitConnNum = db.maxConnNum / 4
		}
	} else {
		db.maxConnNum = DefaultMaxConnNum
		db.InitConnNum = InitConnCount
	}
	//check connection
	db.checkConn, err = db.newConn()
	if err != nil {
		db.Close()
		return nil, errors.ErrDatabaseClose
	}

	db.idleConns = make(chan *Conn, db.maxConnNum)
	db.cacheConns = make(chan *Conn, db.maxConnNum)
	atomic.StoreInt32(&(db.state), Unknown)

	for i := 0; i < db.maxConnNum; i++ {
		if i < db.InitConnNum {
			conn, err := db.newConn()
			if err != nil {
				db.Close()
				return nil, errors.ErrDBPoolInit
			}
			db.cacheConns <- conn
		} else {
			conn := new(Conn)
			db.idleConns <- conn
		}
	}

	return db, nil
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
	return len(db.cacheConns)
}

func (db *DB) Close() error {
	db.Lock()
	idleChannel := db.idleConns
	cacheChannel := db.cacheConns
	db.cacheConns = nil
	db.idleConns = nil
	db.Unlock()
	if cacheChannel == nil || idleChannel == nil {
		return nil
	}

	close(cacheChannel)
	for conn := range cacheChannel {
		db.closeConn(conn)
	}
	close(idleChannel)
	return nil
}

func (db *DB) getConns() (chan *Conn, chan *Conn) {
	db.Lock()
	cacheConns := db.cacheConns
	idleConns := db.idleConns
	db.Unlock()
	return cacheConns, idleConns
}

func (db *DB) getCacheConns() chan *Conn {
	db.Lock()
	conns := db.cacheConns
	db.Unlock()
	return conns
}

func (db *DB) getIdleConns() chan *Conn {
	db.Lock()
	conns := db.idleConns
	db.Unlock()
	return conns
}

func (db *DB) Ping() error {
	var err error
	if db.checkConn == nil {
		db.checkConn, err = db.newConn()
		if err != nil {
			return err
		}
	}
	err = db.checkConn.Ping()
	return err
}

func (db *DB) newConn() (*Conn, error) {
	co := new(Conn)

	if err := co.Connect(db.addr, db.user, db.password, db.db); err != nil {
		return nil, err
	}

	return co, nil
}

func (db *DB) closeConn(co *Conn) error {
	if co != nil {
		co.Close()
		conns := db.getIdleConns()
		conns <- co
	}
	return nil
}

func (db *DB) tryReuse(co *Conn) error {
	var err error
	//reuse Connection
	if co.IsInTransaction() {
		//we can not reuse a connection in transaction status
		err = co.Rollback()
		if err != nil {
			return err
		}
	}

	if !co.IsAutoCommit() {
		//we can not  reuse a connection not in autocomit
		_, err = co.exec("set autocommit = 1")
		if err != nil {
			return err
		}
	}

	//connection may be set names early
	//we must use default utf8
	if co.GetCharset() != mysql.DEFAULT_CHARSET {
		err = co.SetCharset(mysql.DEFAULT_CHARSET)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) PopConn() (*Conn, error) {
	var co *Conn
	var err error

	cacheConns, idleConns := db.getConns()
	if cacheConns == nil || idleConns == nil {
		return nil, errors.ErrDatabaseClose
	}

	if 0 < len(cacheConns) {
		co = <-cacheConns
	} else {
		select {
		case co = <-idleConns:
			err = co.Connect(db.addr, db.user, db.password, db.db)
			if err != nil {
				db.closeConn(co)
				return nil, err
			}
			return co, nil
		case co = <-cacheConns:
			break
		}
	}

	if co == nil {
		return nil, errors.ErrConnIsNil
	}
	err = db.tryReuse(co)
	if err != nil {
		db.closeConn(co)
		return nil, err
	}

	return co, nil
}

func (db *DB) PushConn(co *Conn, err error) {
	if co == nil {
		return
	}
	conns := db.getCacheConns()
	if conns == nil {
		co.Close()
		return
	}
	if err != nil {
		db.closeConn(co)
		return
	}

	select {
	case conns <- co:
		return
	default:
		db.closeConn(co)
		return
	}
}

type BackendConn struct {
	*Conn

	db *DB
}

func (p *BackendConn) Close() {
	if p.Conn != nil {
		if p.Conn.pkgErr != nil {
			p.db.closeConn(p.Conn)
		} else {
			p.db.PushConn(p.Conn, nil)
		}
	}
}

func (db *DB) GetConn() (*BackendConn, error) {
	c, err := db.PopConn()
	if err != nil {
		return nil, err
	}
	return &BackendConn{c, db}, nil
}
