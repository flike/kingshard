package backend

import (
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
)

const (
	Master      = "master"
	Slave       = "slave"
	SlaveSplit  = ","
	WeightSplit = "@"
)

type Node struct {
	Cfg config.NodeConfig
	sync.Mutex
	Master *DB

	Slave          []*DB
	LastSlaveIndex int
	RoundRobinQ    []int
	SlaveWeights   []int

	DownAfterNoAlive time.Duration

	LastMasterPing int64
	LastSlavePing  int64
}

func (n *Node) Run() {
	//to do
	//1 check connection alive
	//2 check remove mysql server alive

	n.checkMaster()
	n.checkSlave()

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

func (n *Node) FormatSlave() string {
	s := make([]byte, 0, 16)
	for _, v := range n.Slave {
		s = append(s, []byte(v.addr)...)
		s = append(s, []byte("\n")...)
	}
	return string(s)
}

func (n *Node) GetMasterConn() (*BackendConn, error) {
	db := n.Master
	if db == nil {
		return nil, errors.ErrNoMasterConn
	}
	if atomic.LoadInt32(&(db.state)) == Down {
		return nil, errors.ErrMasterDown
	}

	return db.GetConn()
}

func (n *Node) GetSlaveConn() (*BackendConn, error) {
	n.Lock()
	db, err := n.GetNextSlave()
	n.Unlock()
	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.ErrNoSlaveDb
	}
	if atomic.LoadInt32(&(db.state)) == Down {
		return nil, errors.ErrSlaveDown
	}

	return db.GetConn()
}

func (n *Node) checkMaster() {
	db := n.Master
	if db == nil {
		golog.Error("Node", "checkMaster", "Master is no alive", 0)
		return
	}
	if atomic.LoadInt32(&(db.state)) == Down {
		return
	}

	if err := db.Ping(); err != nil {
		golog.Error("Node", "checkMaster", "Ping", 0, "db.Addr", db.Addr(), "error", err.Error())
	} else {
		n.LastMasterPing = time.Now().Unix()
	}

	if int64(n.DownAfterNoAlive) > 0 && time.Now().Unix()-n.LastMasterPing > int64(n.DownAfterNoAlive) {
		golog.Info("Node", "checkMaster", "Master down", 0,
			"db.Addr", db.Addr(),
			"Master_down_time", int64(n.DownAfterNoAlive/time.Second))
		n.DownMaster(db.addr)
	}
}

func (n *Node) checkSlave() {
	n.Lock()
	if n.Slave == nil {
		n.Unlock()
		return
	}
	slaves := make([]*DB, len(n.Slave))
	copy(slaves, n.Slave)
	n.Unlock()

	for i := 0; i < len(slaves); i++ {
		if atomic.LoadInt32(&(slaves[i].state)) == Down {
			continue
		}
		if err := slaves[i].Ping(); err != nil {
			golog.Error("Node", "checkSlave", "Ping", 0, "db.Addr", slaves[i].Addr(), "error", err.Error())
		} else {
			n.LastSlavePing = time.Now().Unix()
		}

		if int64(n.DownAfterNoAlive) > 0 && time.Now().Unix()-n.LastSlavePing > int64(n.DownAfterNoAlive) {
			golog.Info("Node", "checkMaster", "Master down", 0,
				"db.Addr", slaves[i].Addr(),
				"slave_down_time", int64(n.DownAfterNoAlive/time.Second))
			//If can't ping slave after DownAfterNoAlive set slave Down
			n.DownSlave(slaves[i].addr)
		}
	}

}

func (n *Node) OpenDB(addr string) *DB {
	db := Open(addr, n.Cfg.User, n.Cfg.Password, "")

	db.SetMaxIdleConnNum(n.Cfg.IdleConns)
	return db
}

func (n *Node) checkUpDB(addr string) (*DB, error) {
	db := n.OpenDB(addr)

	if err := db.Ping(); err != nil {
		db.Close()
		atomic.StoreInt32(&(db.state), Down)
		return nil, err
	}
	atomic.StoreInt32(&(db.state), Up)
	return db, nil
}

func (n *Node) UpMaster(addr string) error {
	db, err := n.checkUpDB(addr)
	if err != nil {
		golog.Error("Node", "UpMaster", err.Error(), 0)
	}
	n.Master = db
	return err
}

func (n *Node) UpSlave(addr string) error {
	db, err := n.checkUpDB(addr)
	if err != nil {
		golog.Error("Node", "UpSlave", err.Error(), 0)
	}

	n.Lock()
	for k, slave := range n.Slave {
		if slave.addr == addr {
			n.Slave[k] = db
			n.Unlock()
			return nil
		}
	}
	n.Slave = append(n.Slave, db)
	n.Unlock()

	return err
}

func (n *Node) DownMaster(addr string) error {
	db := n.Master
	if db == nil || db.addr != addr {
		return errors.ErrNoMasterDb
	}
	db.Close()
	atomic.StoreInt32(&(db.state), Down)
	return nil
}

func (n *Node) DownSlave(addr string) error {
	n.Lock()
	if n.Slave == nil {
		n.Unlock()
		return errors.ErrNoSlaveDb
	}
	slaves := make([]*DB, len(n.Slave))
	copy(slaves, n.Slave)
	n.Unlock()

	//slave is *DB
	for _, slave := range slaves {
		if slave.addr == addr {
			slave.Close()
			atomic.StoreInt32(&(slave.state), Down)
			return nil
		}
	}
	return nil
}

func (n *Node) ParseMaster(masterStr string) error {
	if len(masterStr) == 0 {
		return errors.ErrNoMasterDb
	}

	n.Master = n.OpenDB(masterStr)
	return nil
}

//slaveStr(127.0.0.1:3306@2)
func (n *Node) ParseSlave(slaveStr string) error {
	var db *DB
	var weight int
	var err error

	if len(slaveStr) == 0 {
		return nil
	}
	slaveStr = strings.Trim(slaveStr, SlaveSplit)
	slaveArray := strings.Split(slaveStr, SlaveSplit)
	count := len(slaveArray)
	n.Slave = make([]*DB, 0, count)
	n.SlaveWeights = make([]int, 0, count)

	//parse addr and weight
	for i := 0; i < count; i++ {
		addrAndWeight := strings.Split(slaveArray[i], WeightSplit)
		if len(addrAndWeight) == 2 {
			weight, err = strconv.Atoi(addrAndWeight[1])
			if err != nil {
				return err
			}
		} else {
			weight = 1
		}
		n.SlaveWeights = append(n.SlaveWeights, weight)
		db = n.OpenDB(addrAndWeight[0])
		n.Slave = append(n.Slave, db)
	}
	n.InitBalancer()
	return nil
}
