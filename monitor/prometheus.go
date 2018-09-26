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

package monitor

import (
	"sync"
	"net/http"
	"time"
	"strconv"

	"github.com/flike/kingshard/proxy/server"
	"github.com/flike/kingshard/core/golog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Prometheus struct {
	addr string
	svr  *server.Server
	data sync.Map
}

//新建prometheus实例
func NewPrometheus(addr string, svr *server.Server) (*Prometheus, error) {
	prometheus := new(Prometheus)
	prometheus.addr = addr
	prometheus.svr  = svr

	golog.Info("prometheus", "Run", "Prometheus running", 0,
		"address",
		addr)

	return prometheus, nil 
}

//启动prometheus的http监控
func (p *Prometheus) Run() {
	// 开始每秒钟获取一次数据
	go func() {
		for {
			data := p.svr.GetMonitorData()
		
			for _, data := range data {
				p.data.Store("idleConn", data["idleConn"])
				p.data.Store("maxConn", data["maxConn"])
				p.data.Store("cacheConns", data["cacheConns"])
				p.data.Store("pushConnCount", data["pushConnCount"])
				p.data.Store("popConnCount", data["popConnCount"])
			}

			time.Sleep(1 * time.Second)
		}
	}()

	//设置标签及注册
	data := p.svr.GetMonitorData()

	label := make(map[string]string)

	for addr, data := range data {
		label["addr"] = addr
		label["type"] = data["type"]

		p.addGauge("idleConn", "the db idle connection", label)
		p.addGauge("cacheConns", "the db cache connection", label)
		p.addGauge("pushConnCount", "the db pushConnCount", label)
		p.addGauge("popConnCount", "the db popConnCount", label)
		p.addGauge("maxConn", "the max connection config", label)
	}

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(p.addr, mux)

	if err != nil {
		golog.Error("prometheus", "Run", err.Error(), 0)
	}
}

func (p *Prometheus) addGauge(name string, help string, label map[string]string) {
	gauge := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: name,
			Help: help,
			ConstLabels: label,
		},
	)

	prometheus.MustRegister(gauge)

	go func() {
		for {
			pValueInterface, _ 	:= p.data.Load(name)
			pValueString, ok 	:= pValueInterface.(string)
			
			if ok {
				floatValue, _ := strconv.ParseFloat(pValueString, 10)
				gauge.Set(floatValue)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}