package backend

import (
	"fmt"
	"github.com/flike/kingshard/config"
	. "github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"sync"
	"time"
)

const (
	Master = "master"
	Slave  = "slave"
)

type Node struct {
	sync.Mutex

	Cfg config.NodeConfig

	Master *DB
	Slave  *DB

	DownAfterNoAlive time.Duration

	LastMasterPing int64
	LastSlavePing  int64
}

func (n *Node) Run() {
	//to do
	//1 check connection alive
	//2 check remove mysql server alive

	t := time.NewTicker(3000 * time.Second)
	defer t.Stop()

	n.LastMasterPing = time.Now().Unix()
	n.LastSlavePing = n.LastMasterPing
	for {
		select {
		case <-t.C:
			n.checkMaster()
			n.checkSlave()
		}
	}
}

func (n *Node) String() string {
	return n.Cfg.Name
}

func (n *Node) GetMasterConn() (*BackendConn, error) {
	n.Lock()
	db := n.Master
	n.Unlock()

	if db == nil {
		return nil, ErrNoMasterConn
	}

	return db.GetConn()
}

func (n *Node) GetSlaveConn() (*BackendConn, error) {
	n.Lock()
	db := n.Slave
	n.Unlock()

	if db == nil {
		return nil, ErrNoSlaveConn
	}

	return db.GetConn()
}

func (n *Node) checkMaster() {
	n.Lock()
	db := n.Master
	n.Unlock()

	if db == nil {
		golog.Error("Node", "checkMaster", "Master is no alive", 0)
		return
	}

	if err := db.Ping(); err != nil {
		golog.Error("Node", "checkMaster", "Ping", 0, "db.Addr", db.Addr(), "error", err.Error())
	} else {
		n.LastMasterPing = time.Now().Unix()
		return
	}

	if int64(n.DownAfterNoAlive) > 0 && time.Now().Unix()-n.LastMasterPing > int64(n.DownAfterNoAlive) {
		golog.Info("Node", "checkMaster", "Master down", 0,
			"db.Addr", db.Addr(),
			"Master_down_time", int64(n.DownAfterNoAlive/time.Second))
		n.DownMaster()
	}
}

func (n *Node) checkSlave() {
	if n.Slave == nil {
		return
	}

	db := n.Slave
	if err := db.Ping(); err != nil {
		golog.Error("Node", "checkSlave", "Ping", 0, "db.Addr", db.Addr(), "error", err.Error())
	} else {
		n.LastSlavePing = time.Now().Unix()
	}

	if int64(n.DownAfterNoAlive) > 0 && time.Now().Unix()-n.LastSlavePing > int64(n.DownAfterNoAlive) {
		golog.Info("Node", "checkMaster", "Master down", 0,
			"db.Addr", db.Addr(),
			"slave_down_time", int64(n.DownAfterNoAlive/time.Second))

		n.DownSlave()
	}
}

func (n *Node) OpenDB(addr string) (*DB, error) {
	db, err := Open(addr, n.Cfg.User, n.Cfg.Password, "")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConnNum(n.Cfg.IdleConns)
	return db, nil
}

func (n *Node) checkUpDB(addr string) (*DB, error) {
	db, err := n.OpenDB(addr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func (n *Node) UpMaster(addr string) error {
	n.Lock()
	if n.Master != nil {
		n.Unlock()
		return fmt.Errorf("%s master must be down first", n)
	}
	n.Unlock()

	db, err := n.checkUpDB(addr)
	if err != nil {
		return err
	}

	n.Lock()
	n.Master = db
	n.Unlock()

	return nil
}

func (n *Node) UpSlave(addr string) error {
	n.Lock()
	if n.Slave != nil {
		n.Unlock()
		return fmt.Errorf("%s, slave must be down first", n)
	}
	n.Unlock()

	db, err := n.checkUpDB(addr)
	if err != nil {
		return err
	}

	n.Lock()
	n.Slave = db
	n.Unlock()

	return nil
}

func (n *Node) DownMaster() error {
	n.Lock()
	if n.Master != nil {
		n.Master = nil
	}
	return nil
}

func (n *Node) DownSlave() error {
	n.Lock()
	db := n.Slave
	n.Slave = nil
	n.Unlock()

	if db != nil {
		db.Close()
	}

	return nil
}
