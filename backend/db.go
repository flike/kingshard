package backend

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
	"github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/core/golog"
)

const (
	Up = iota
	Down
	Unknown
)

type DB struct {
	sync.Mutex

	addr         string
	user         string
	password     string
	db           string
	wait_timeout int

	maxIdleConns int
	maxConns     int32
	state        int32

    conns chan *Conn
    initialConns int
    numConn int32
}

func Open(addr string, user string, password string, dbName string,wait_timeout int,initialConns,maxConns int) (*DB ,error){
	db := new(DB)

	db.addr = addr
	db.user = user
	db.password = password
	db.db = dbName
	db.wait_timeout = wait_timeout

	atomic.StoreInt32(&(db.state), Unknown)

    db.initialConns = initialConns
    db.maxConns = int32(maxConns)
    db.numConn = 0
    db.conns = make(chan *Conn, db.maxConns)

    // create initial connections, if something goes wrong,
    // just close the pool error out.
    for i := 0; i < db.initialConns; i++ {
        conn, err := db.newConn()
        if err != nil {
            db.Close()
            return nil, fmt.Errorf("not able to fill the pool: %s", err)
        }
        db.conns <- conn
    }

	return db,nil
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
    return len(db.getConns())
}

func (db *DB) Close() {
    db.Lock()
    conns := db.conns
    db.conns = nil
    db.Unlock()

    if conns == nil {
        return
    }

    close(conns)
    for conn := range conns {
        db.closeConn(conn)
    }
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

func (db *DB) getConns() chan *Conn {
    db.Lock()
    conns := db.conns
    db.Unlock()
    return conns
}

func (db *DB) GetIdleConnNum() int {
    return len(db.getConns())
}

func (db *DB) newConn() (*Conn, error) {
	co := new(Conn)

	if err := co.Connect(db.addr, db.user, db.password, db.db,db.wait_timeout); err != nil {
		return nil, err
	}
	atomic.AddInt32(&(db.numConn),1)
	return co, nil
}

func (db *DB) closeConn(co *Conn) error {
    atomic.AddInt32(&(db.numConn),-1)
	return co.Close()
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

func (db *DB) PopConn() (co *Conn, err error) {
    golog.Error("begin","","",0,"db.numConn",db.numConn)
    conns := db.getConns()
    if conns == nil {
        return nil,fmt.Errorf("database connection pool closed")
    }

    select {
    case co := <-conns:
        if co == nil {
            return nil,fmt.Errorf("database connection pool closed")
        }else{
            if err := co.Ping(); err == nil {
                if err := db.tryReuse(co); err == nil {
                    //connection may alive
                    return co, nil
                }
            }
            db.closeConn(co)
            return nil,fmt.Errorf("connection can't use , please get again")
        }
    case <- time.After(time.Millisecond * 500):
        db.Lock()
        if db.numConn >= db.maxConns {
            db.Unlock()
            return nil,fmt.Errorf("exceed pool size %d",db.maxConns)
        }
        golog.Error("create","","",0,"db.numConn",db.numConn)
        co,err := db.newConn()
        db.Unlock()
        if err != nil {
            return nil,err
        }
        return co,err
    }
}

func (db *DB) PushConn(co *Conn, err error) {
    db.Lock()
    defer db.Unlock()

    if db.conns == nil {
        db.closeConn(co)
        return
    }

    // put the resource back into the pool. If the pool is full, this will
    // block and the default case will be executed.
    select {
    case db.conns <- co:
        return
    default:
        // pool is full, close passed connection
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
        }else {
            p.db.PushConn(p.Conn, p.Conn.pkgErr)
            p.Conn = nil
        }
    }
}

func (db *DB) GetConn() (*BackendConn, error) {
	c, err := db.PopConn()
	backend_conn := db.wrapConn(c)
	return backend_conn, err
}

func (db *DB) wrapConn(conn *Conn) *BackendConn {
    backend_conn := &BackendConn{conn, db}
    return backend_conn
}
