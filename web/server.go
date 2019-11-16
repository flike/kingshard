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
package web

import (
	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/proxy/server"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	//"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
)

type ApiServer struct {
	cfg         *config.Config
	proxy       *server.Server
	webAddr     string
	webUser     string
	webPassword string
	web *echo.Echo
}

func NewApiServer(cfg *config.Config, srv *server.Server) (*ApiServer, error) {
	s := new(ApiServer)
	s.cfg = cfg
	s.proxy = srv
	s.webAddr = cfg.WebAddr
	s.webUser = cfg.WebUser
	s.webPassword = cfg.WebPassword
	s.web = echo.New()
	golog.Info("web", "NewApiServer", "Api Server running", 0,
		"netProto",
		"http",
		"address",
		s.webAddr)
	return s, nil
}

func (s *ApiServer) Run() error {
	s.web.HideBanner = true
	s.web.HidePort = true

	s.RegisterMiddleware()
	s.RegisterURL()
	//std := standard.New(s.webAddr)
	//std.SetHandler(s)
	//graceful.ListenAndServe(std.Server, 5*time.Second)
	//return nil
	err := s.web.Start(s.webAddr)
	if err != nil {
		log.Errorf("AdminServer.Start:web server start error,err:%s", err)
	}
	return nil
}

func (s *ApiServer) RegisterMiddleware() {
	//s.Use(mw.Logger())
	s.web.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		Output: golog.GlobalSqlLogger,
	}))
	s.web.Use(mw.Recover())
	s.web.Use(mw.BasicAuth(s.CheckAuth))
}

func (s *ApiServer) RegisterURL() {
	s.web.GET("/api/v1/nodes/status", s.GetNodesStatus)

	s.web.POST("/api/v1/nodes/slaves", s.AddOneSlave)
	s.web.DELETE("/api/v1/nodes/slaves", s.DeleteOneSlave)
	s.web.PUT("/api/v1/nodes/slaves/status", s.ChangeSlaveStatus)

	s.web.PUT("/api/v1/nodes/masters/status", s.ChangeMasterStatus)

	s.web.GET("/api/v1/proxy/status", s.GetProxyStatus)
	s.web.PUT("/api/v1/proxy/status", s.ChangeProxyStatus)

	s.web.GET("/api/v1/proxy/schema", s.GetProxySchema)

	s.web.GET("/api/v1/proxy/allow_ips", s.GetAllowIps)
	s.web.POST("/api/v1/proxy/allow_ips", s.AddAllowIps)
	s.web.DELETE("/api/v1/proxy/allow_ips", s.DelAllowIps)

	s.web.GET("/api/v1/proxy/black_sqls", s.GetAllBlackSQL)
	s.web.POST("/api/v1/proxy/black_sqls", s.AddOneBlackSQL)
	s.web.DELETE("/api/v1/proxy/black_sqls", s.DelOneBlackSQL)

	s.web.GET("/api/v1/proxy/slow_sql/time", s.GetSlowLogTime)
	s.web.PUT("/api/v1/proxy/slow_sql/status", s.SwitchSlowSQL)
	s.web.PUT("/api/v1/proxy/slow_sql/time", s.SetSlowLogTime)

	s.web.PUT("/api/v1/proxy/config/save", s.SaveProxyConfig)
}

func (s *ApiServer) CheckAuth(username, password string,ctx echo.Context) (bool,error) {
	if username == s.webUser && password == s.webPassword {
		return true,nil
	}
	return false,nil
}
