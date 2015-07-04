package server

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/flike/kingshard/backend"
	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/golog"
	. "github.com/flike/kingshard/mysql"
	"github.com/flike/kingshard/proxy/router"
)

type Schema struct {
	db string

	nodes map[string]*backend.Node

	rule *router.Router
}

type Server struct {
	cfg *config.Config

	addr     string
	user     string
	password string
	db       string

	running bool

	listener net.Listener

	allowips []int64

	nodes map[string]*backend.Node

	schemas map[string]*Schema
}

func (s *Server) parseAllowIps() error {
	cfg := s.cfg
	if len(cfg.AllowIps) == 0 {
		return nil
	}
	ipVec := strings.Split(cfg.AllowIps, ",")
	s.allowips = make([]int64, 10)
	for _, ip := range ipVec {
		var ipValue int64 = 0
		ipSeg := strings.Split(ip, ".")
		for _, seg := range ipSeg {
			k, err := strconv.ParseInt(seg, 10, 32)
			if err != nil {
				panic(err)
			}
			ipValue = ipValue + k<<4
		}
		ipVecLen := len(s.allowips)
		for i := 0; i < ipVecLen; i++ {
			if s.allowips[i] == ipValue {
				return fmt.Errorf("duplicate allow ip [%s].", ip)
			}
		}
		s.allowips = append(s.allowips, ipValue)
	}
	return nil
}

func (s *Server) parseNode(cfg config.NodeConfig) (*backend.Node, error) {
	n := new(backend.Node)
	//n.server = s
	n.Cfg = cfg

	n.DownAfterNoAlive = time.Duration(cfg.DownAfterNoAlive) * time.Second

	if len(cfg.Master) == 0 {
		return nil, fmt.Errorf("must setting master MySQL node.")
	}

	var err error
	if n.Master, err = n.OpenDB(cfg.Master); err != nil {
		return nil, err
	}

	//n.db = n.Master

	if len(cfg.Slave) > 0 {
		if n.Slave, err = n.OpenDB(cfg.Slave); err != nil {
			golog.Error("ClientConn", "handleShowProxy", err.Error(), 0)
			n.Slave = nil
		}
	}

	go n.Run()

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

func (s *Server) parseSchemas() error {
	s.schemas = make(map[string]*Schema)

	if len(s.cfg.Schemas) != 1 {
		return fmt.Errorf("must have only one schema.")
	}

	for _, schemaCfg := range s.cfg.Schemas {
		if _, ok := s.schemas[schemaCfg.DB]; ok {
			return fmt.Errorf("duplicate schema [%s].", schemaCfg.DB)
		}
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

		s.schemas[schemaCfg.DB] = &Schema{
			db:    schemaCfg.DB,
			nodes: nodes,
			rule:  rule,
		}
		s.db = schemaCfg.DB
	}

	return nil
}

func NewServer(cfg *config.Config) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	s.addr = cfg.Addr
	s.user = cfg.User
	s.password = cfg.Password
	fmt.Println(s.cfg.Password)

	if err := s.parseAllowIps(); err != nil {
		return nil, err
	}

	if err := s.parseNodes(); err != nil {
		return nil, err
	}

	if err := s.parseSchemas(); err != nil {
		return nil, err
	}

	var err error
	netProto := "tcp"
	if strings.Contains(netProto, "/") {
		netProto = "unix"
	}
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
func (s *Server) newClientConn(co net.Conn) *ClientConn {
	c := new(ClientConn)

	c.c = co
	c.schema = s.GetSchema(s.db)

	c.pkg = NewPacketIO(co)

	c.proxy = s

	c.c = co
	c.pkg.Sequence = 0

	c.connectionId = atomic.AddUint32(&baseConnId, 1)

	c.status = SERVER_STATUS_AUTOCOMMIT

	c.salt, _ = RandomBuf(20)

	c.txConns = make(map[*backend.Node]*backend.BackendConn)

	c.closed = false

	c.collation = DEFAULT_COLLATION_ID
	c.charset = DEFAULT_CHARSET

	c.stmtId = 0
	c.stmts = make(map[uint32]*Stmt)

	return c
}

func (s *Server) onConn(c net.Conn) {
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
	}()

	if allowConnect := conn.IsAllowConnect(); allowConnect == false {
		err := NewError(ER_UNKNOWN_ERROR, "ip address access denied by myshard.")
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
	/*server一直在监听，监听到conn后，启动一个协程来处理*/
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
func (s *Server) DownMaster(node string) error {
	n := s.GetNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}
	//	n.db = nil
	return n.DownMaster()
}

func (s *Server) DownSlave(node string) error {
	n := s.GetNode(node)
	if n == nil {
		return fmt.Errorf("invalid node [%s].", node)
	}
	return n.DownSlave()
}

func (s *Server) GetNode(name string) *backend.Node {
	return s.nodes[name]
}

func (s *Server) GetSchema(db string) *Schema {
	return s.schemas[db]
}
