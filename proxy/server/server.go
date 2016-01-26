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

package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/flike/kingshard/mysql"

	"github.com/flike/kingshard/backend"
	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/proxy/router"
)

type Schema struct {
	db string

	nodes map[string]*backend.Node

	rule *router.Router
}

type BlacklistSqls struct {
	sqls    map[string]string
	sqlsLen int
}

type Server struct {
	cfg           *config.Config
	addr          string
	user          string
	password      string
	db            string
	blacklistSqls *BlacklistSqls
	allowips      []net.IP
	counter       *Counter
	nodes         map[string]*backend.Node
	schema        *Schema

	listener net.Listener
	running  bool
}

func (s *Server) parseAllowIps() error {
	cfg := s.cfg
	if len(cfg.AllowIps) == 0 {
		return nil
	}
	ipVec := strings.Split(cfg.AllowIps, ",")
	s.allowips = make([]net.IP, 0, 10)
	for _, ip := range ipVec {
		s.allowips = append(s.allowips, net.ParseIP(strings.TrimSpace(ip)))
	}
	return nil
}

//parse the blacklist sql file
func (s *Server) parseBlackListSqls() error {
	bs := new(BlacklistSqls)
	bs.sqls = make(map[string]string)
	if len(s.cfg.BlsFile) != 0 {
		file, err := os.Open(s.cfg.BlsFile)
		if err != nil {
			return err
		}

		defer file.Close()
		rd := bufio.NewReader(file)
		for {
			line, err := rd.ReadString('\n')
			//end of file
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			line = strings.TrimSpace(line)
			if len(line) != 0 {
				fingerPrint := mysql.GetFingerprint(line)
				md5 := mysql.GetMd5(fingerPrint)
				bs.sqls[md5] = fingerPrint
			}
		}
	}
	bs.sqlsLen = len(bs.sqls)
	s.blacklistSqls = bs

	return nil
}

func (s *Server) parseNode(cfg config.NodeConfig) (*backend.Node, error) {
	var err error
	n := new(backend.Node)
	n.Cfg = cfg

	n.DownAfterNoAlive = time.Duration(cfg.DownAfterNoAlive) * time.Second
	err = n.ParseMaster(cfg.Master)
	if err != nil {
		return nil, err
	}
	err = n.ParseSlave(cfg.Slave)
	if err != nil {
		return nil, err
	}

	go n.CheckNode()

	return n, nil
}

func (s *Server) parseNodes() error {
	cfg := s.cfg
	s.nodes = make(map[string]*backend.Node, len(cfg.Nodes))

	for _, v := range cfg.Nodes {
		if _, ok := s.nodes[v.Name]; ok {
			return fmt.Errorf("duplicate node [%s].", v.Name)
		}

		n, err := s.parseNode(v)
		if err != nil {
			return err
		}

		s.nodes[v.Name] = n
	}

	return nil
}

func (s *Server) parseSchema() error {
	schemaCfg := s.cfg.Schema
	if len(schemaCfg.Nodes) == 0 {
		return fmt.Errorf("schema [%s] must have a node.", schemaCfg.DB)
	}

	nodes := make(map[string]*backend.Node)
	for _, n := range schemaCfg.Nodes {
		if s.GetNode(n) == nil {
			return fmt.Errorf("schema [%s] node [%s] config is not exists.", schemaCfg.DB, n)
		}

		if _, ok := nodes[n]; ok {
			return fmt.Errorf("schema [%s] node [%s] duplicate.", schemaCfg.DB, n)
		}

		nodes[n] = s.GetNode(n)
	}

	rule, err := router.NewRouter(&schemaCfg)
	if err != nil {
		return err
	}

	s.schema = &Schema{
		db:    schemaCfg.DB,
		nodes: nodes,
		rule:  rule,
	}
	s.db = schemaCfg.DB

	return nil
}

func NewServer(cfg *config.Config) (*Server, error) {
	s := new(Server)

	s.cfg = cfg
	s.counter = new(Counter)

	s.addr = cfg.Addr
	s.user = cfg.User
	s.password = cfg.Password

	if err := s.parseBlackListSqls(); err != nil {
		return nil, err
	}

	if err := s.parseAllowIps(); err != nil {
		return nil, err
	}

	if err := s.parseNodes(); err != nil {
		return nil, err
	}

	if err := s.parseSchema(); err != nil {
		return nil, err
	}

	var err error
	netProto := "tcp"

	s.listener, err = net.Listen(netProto, s.addr)

	if err != nil {
		return nil, err
	}

	golog.Info("server", "NewServer", "Server running", 0,
		"netProto",
		netProto,
		"address",
		s.addr)
	return s, nil
}

func (s *Server) flushCounter() {
	for {
		s.counter.FlushCounter()
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) newClientConn(co net.Conn) *ClientConn {
	c := new(ClientConn)
	tcpConn := co.(*net.TCPConn)

	//SetNoDelay controls whether the operating system should delay packet transmission
	// in hopes of sending fewer packets (Nagle's algorithm).
	// The default is true (no delay),
	// meaning that data is sent as soon as possible after a Write.
	//I set this option false.
	tcpConn.SetNoDelay(false)
	c.c = tcpConn

	c.schema = s.GetSchema()

	c.pkg = mysql.NewPacketIO(tcpConn)
	c.proxy = s

	c.pkg.Sequence = 0

	c.connectionId = atomic.AddUint32(&baseConnId, 1)

	c.status = mysql.SERVER_STATUS_AUTOCOMMIT

	c.salt, _ = mysql.RandomBuf(20)

	c.txConns = make(map[*backend.Node]*backend.BackendConn)

	c.closed = false

	c.collation = mysql.DEFAULT_COLLATION_ID
	c.charset = mysql.DEFAULT_CHARSET

	c.stmtId = 0
	c.stmts = make(map[uint32]*Stmt)

	return c
}

func (s *Server) onConn(c net.Conn) {
	s.counter.IncrClientConns()
	conn := s.newClientConn(c) //新建一个conn

	defer func() {
		err := recover()
		if err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)] //获得当前goroutine的stacktrace
			golog.Error("server", "onConn", "error", 0,
				"remoteAddr", c.RemoteAddr().String(),
				"stack", string(buf),
			)
		}

		conn.Close()
		s.counter.DecrClientConns()
	}()

	if allowConnect := conn.IsAllowConnect(); allowConnect == false {
		err := mysql.NewError(mysql.ER_ACCESS_DENIED_ERROR, "ip address access denied by kingshard.")
		conn.writeError(err)
		conn.Close()
		return
	}
	if err := conn.Handshake(); err != nil {
		golog.Error("server", "onConn", err.Error(), 0)
		c.Close()
		return
	}

	conn.Run()
}

func (s *Server) Run() error {
	s.running = true

	// flush counter
	go s.flushCounter()

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			golog.Error("server", "Run", err.Error(), 0)
			continue
		}

		go s.onConn(conn)
	}

	return nil
}

func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *Server) DeleteSlave(node string, addr string) error {
	n := s.GetNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}

	return n.DeleteSlave(addr)
}

func (s *Server) AddSlave(node string, addr string) error {
	n := s.GetNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}

	return n.AddSlave(addr)
}

func (s *Server) UpMaster(node string, addr string) error {
	n := s.GetNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}

	return n.UpMaster(addr)
}

func (s *Server) UpSlave(node string, addr string) error {
	n := s.GetNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}

	return n.UpSlave(addr)
}

func (s *Server) DownMaster(node, masterAddr string) error {
	n := s.GetNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}
	return n.DownMaster(masterAddr, backend.ManualDown)
}

func (s *Server) DownSlave(node, slaveAddr string) error {
	n := s.GetNode(node)
	if n == nil {
		return fmt.Errorf("invalid node [%s].", node)
	}
	return n.DownSlave(slaveAddr, backend.ManualDown)
}

func (s *Server) GetNode(name string) *backend.Node {
	return s.nodes[name]
}

func (s *Server) GetSchema() *Schema {
	return s.schema
}
