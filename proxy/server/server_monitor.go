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
	"net"
	"runtime"
	"time"

	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/golog"
)

type ServerMonitor struct {
	addr     string
	listener net.Listener
	running  bool
}

func NewMonitorServer(cfg *config.Config) (*ServerMonitor, error) {
	s := new(ServerMonitor)
	s.addr = cfg.MonitorAddr

	var err error
	netProto := "tcp"

	s.listener, err = net.Listen(netProto, s.addr)

	if err != nil {
		return nil, err
	}

	golog.Info("server", "ServerMonitor", "Server running", 0,
		"netProto",
		netProto,
		"address",
		s.addr)
	return s, nil
}

func (s *ServerMonitor) reply(c net.Conn) {

	defer c.Close()
	c.SetDeadline(time.Now().Add(time.Second * 5))

	replyData := "OKOKOK"
	_, err := c.Write([]byte(replyData))
	if err != nil {
		golog.Error("ServerMonitor", "reply", "error", 0,
			"error info: ", err.Error())
		return
	}

	// clear network buffer
	rb := make([]byte, 1024)
	for {
		n, err := c.Read(rb)
		if err != nil || n <= 0 {
			return
		} else {
			golog.Info("ServerMonitor", "reply", "receive", 0, "msg: ", rb)
		}
	}
}

func (s *ServerMonitor) onConn(c net.Conn) {

	defer func() {
		err := recover()
		if err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)] //获得当前goroutine的stacktrace
			golog.Error("ServerMonitor", "onConn", "error", 0,
				"remoteAddr", c.RemoteAddr().String(),
				"stack", string(buf),
			)
		}

		c.Close()
	}()

	s.reply(c)
}

func (s *ServerMonitor) Run() error {
	s.running = true
	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			golog.Error("ServerMonitor", "Run", err.Error(), 0)
			continue
		}

		go s.onConn(conn)
	}

	return nil
}

func (s *ServerMonitor) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
}
