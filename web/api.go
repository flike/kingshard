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
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	ksError "github.com/flike/kingshard/core/errors"
	"github.com/flike/kingshard/core/golog"
	"github.com/labstack/echo"
)

func (s *ApiServer) GetAllowIps(c echo.Context) error {
	allowIps := s.proxy.GetAllowIps()
	return c.JSON(http.StatusOK, allowIps)
}

//add one or multi ips
func (s *ApiServer) AddAllowIps(c echo.Context) error {
	args := struct {
		AllowIPs []string `json:"allow_ips"`
	}{}
	err := c.Bind(&args)
	if err != nil {
		return err
	}
	for _, v := range args.AllowIPs {
		err = s.proxy.AddAllowIP(strings.TrimSpace(v))
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, "ok")
}

//delete one or multi ips
func (s *ApiServer) DelAllowIps(c echo.Context) error {
	args := struct {
		AllowIPs []string `json:"allow_ips"`
	}{}
	err := c.Bind(&args)
	if err != nil {
		return err
	}
	for _, v := range args.AllowIPs {
		err = s.proxy.DelAllowIP(strings.TrimSpace(v))
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, "ok")
}

type DBStatus struct {
	Node      		string `json:"node"`
	Address   		string `json:"address"`
	Type      		string `json:"type"`
	Status    		string `json:"status"`
	LastPing  		string `json:"laste_ping"`
	MaxConn   		int    `json:"max_conn"`
	IdleConn  		int    `json:"idle_conn"`
	CacheConn 		int    `json:"cache_conn"`
	PushConnCount  	int64  `json:"push_conn_count"`
	PopConnCount   	int64  `json:"pop_conn_count"`
}

//get nodes status
func (s *ApiServer) GetNodesStatus(c echo.Context) error {
	var masterStatus, slaveStatus DBStatus

	dbStatus := make([]DBStatus, 0, 1)
	nodes := s.proxy.GetAllNodes()

	for nodeName, node := range nodes {
		//get master counter
		idleConns,cacheConns,pushConnCount,popConnCount := node.Master.ConnCount()

		//get master status
		masterStatus.Node = nodeName
		masterStatus.Address = node.Master.Addr()
		masterStatus.Type = "master"
		masterStatus.Status = node.Master.State()
		masterStatus.LastPing = fmt.Sprintf("%v", time.Unix(node.Master.GetLastPing(), 0))
		masterStatus.MaxConn = node.Cfg.MaxConnNum
		masterStatus.IdleConn = idleConns
		masterStatus.CacheConn = cacheConns
		masterStatus.PushConnCount = pushConnCount
		masterStatus.PopConnCount = popConnCount
		dbStatus = append(dbStatus, masterStatus)

		//get slaves status
		for _, slave := range node.Slave {
			//get slave counter
			idleConns,cacheConns,pushConnCount,popConnCount := slave.ConnCount()

			slaveStatus.Node = nodeName
			slaveStatus.Address = slave.Addr()
			slaveStatus.Type = "slave"
			slaveStatus.Status = slave.State()
			slaveStatus.LastPing = fmt.Sprintf("%v", time.Unix(slave.GetLastPing(), 0))
			slaveStatus.MaxConn = node.Cfg.MaxConnNum
			slaveStatus.IdleConn = idleConns
			slaveStatus.CacheConn = cacheConns
			slaveStatus.PushConnCount = pushConnCount
			slaveStatus.PopConnCount = popConnCount
			dbStatus = append(dbStatus, slaveStatus)
		}
	}
	return c.JSON(http.StatusOK, dbStatus)
}

func (s *ApiServer) AddOneSlave(c echo.Context) error {
	args := struct {
		Node string `json:"node"`
		Addr string `json:"addr"`
	}{}
	err := c.Bind(&args)
	if err != nil {
		return err
	}
	err = s.proxy.AddSlave(args.Node, args.Addr)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func (s *ApiServer) DeleteOneSlave(c echo.Context) error {
	args := struct {
		Node string `json:"node"`
		Addr string `json:"addr"`
	}{}
	err := c.Bind(&args)
	if err != nil {
		return err
	}
	err = s.proxy.DeleteSlave(args.Node, args.Addr)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func (s *ApiServer) ChangeSlaveStatus(c echo.Context) error {
	args := struct {
		Opt  string `json:"opt"`
		Node string `json:"node"`
		Addr string `json:"addr"`
	}{}

	err := c.Bind(&args)
	if err != nil {
		return err
	}

	args.Opt = strings.ToLower(args.Opt)
	if args.Opt != "up" && args.Opt != "down" {
		return errors.New("opt only can be up or down")
	}
	if args.Opt == "down" {
		err = s.proxy.DownSlave(args.Node, args.Addr)
	} else {
		err = s.proxy.UpSlave(args.Node, args.Addr)
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func (s *ApiServer) ChangeMasterStatus(c echo.Context) error {
	args := struct {
		Opt  string `json:"opt"`
		Node string `json:"node"`
		Addr string `json:"addr"`
	}{}

	err := c.Bind(&args)
	if err != nil {
		return err
	}

	args.Opt = strings.ToLower(args.Opt)
	if args.Opt != "up" && args.Opt != "down" {
		return errors.New("opt only can be up or down")
	}
	if args.Opt == "down" {
		err = s.proxy.DownMaster(args.Node, args.Addr)
	} else {
		err = s.proxy.UpMaster(args.Node, args.Addr)
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func (s *ApiServer) GetProxyStatus(c echo.Context) error {
	status := s.proxy.Status()
	return c.JSON(http.StatusOK, status)
}

func (s *ApiServer) ChangeProxyStatus(c echo.Context) error {
	args := struct {
		Opt string `json:"opt"`
	}{}

	err := c.Bind(&args)
	if err != nil {
		return err
	}
	args.Opt = strings.ToLower(args.Opt)
	if args.Opt != "online" && args.Opt != "offline" {
		return errors.New("opt only can be online or offline")
	}

	err = s.proxy.ChangeProxy(args.Opt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

//range,hash or date
type ShardConfig struct {
	User          string   `json:"user"`
	DB            string   `json:"db"`
	Table         string   `yaml:"table"`
	Key           string   `yaml:"key"`
	Nodes         []string `yaml:"nodes"`
	Locations     []int    `yaml:"locations"`
	Type          string   `yaml:"type"`
	TableRowLimit int      `yaml:"table_row_limit"`
	DateRange     []string `yaml:"date_range"`
}

func (s *ApiServer) GetProxySchema(c echo.Context) error {
	shardConfig := make([]ShardConfig, 0, 10)
	for _, schema := range s.cfg.SchemaList{
		//append default rule
		shardConfig = append(shardConfig,
			ShardConfig{
				User:  schema.User,
				Type:  "default",
				Nodes: schema.Nodes,
			})
		for _, r := range schema.ShardRule {
			shardConfig = append(shardConfig,
				ShardConfig{
					User:		   schema.User,
					DB:            r.DB,
					Table:         r.Table,
					Key:           r.Key,
					Nodes:         r.Nodes,
					Locations:     r.Locations,
					Type:          r.Type,
					TableRowLimit: r.TableRowLimit,
					DateRange:     r.DateRange,
				})
		}

	}
	return c.JSON(http.StatusOK, shardConfig)
}

func (s *ApiServer) GetAllBlackSQL(c echo.Context) error {
	sqls := s.proxy.GetAllBlackSqls()
	return c.JSON(http.StatusOK, sqls)
}

func (s *ApiServer) AddOneBlackSQL(c echo.Context) error {
	args := struct {
		SQL string `json:"sql"`
	}{}

	err := c.Bind(&args)
	if err != nil {
		return err
	}
	err = s.proxy.AddBlackSql(args.SQL)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func (s *ApiServer) DelOneBlackSQL(c echo.Context) error {
	args := struct {
		SQL string `json:"sql"`
	}{}

	err := c.Bind(&args)
	if err != nil {
		return err
	}
	err = s.proxy.DelBlackSql(args.SQL)
	if err != nil {
		if err == ksError.ErrBlackSqlNotExist {
			errMsg := fmt.Sprintf("`%s` isn't exist in black sql", args.SQL)
			return c.JSON(http.StatusNotFound, errMsg)
		}
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func (s *ApiServer) SwitchSlowSQL(c echo.Context) error {
	args := struct {
		Opt string `json:"opt"`
	}{}

	err := c.Bind(&args)
	if err != nil {
		return err
	}
	args.Opt = strings.ToLower(args.Opt)
	if args.Opt != golog.LogSqlOn && args.Opt != golog.LogSqlOff {
		return errors.New("opt only can be on or off")
	}

	err = s.proxy.ChangeLogSql(args.Opt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func (s *ApiServer) SetSlowLogTime(c echo.Context) error {
	args := struct {
		SlowTime int64 `json:"slow_time"`
	}{}

	err := c.Bind(&args)
	if err != nil {
		return err
	}
	slowTimeStr := strconv.FormatInt(args.SlowTime, 10)
	err = s.proxy.ChangeSlowLogTime(slowTimeStr)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func (s *ApiServer) GetSlowLogTime(c echo.Context) error {
	time := s.proxy.GetSlowLogTime()
	return c.JSON(http.StatusOK, time)
}

func (s *ApiServer) SaveProxyConfig(c echo.Context) error {
	err := s.proxy.SaveProxyConfig()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}
