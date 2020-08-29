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
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/proxy/server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Prometheus struct {
	addr      string
	svr       *server.Server
	nodesData sync.Map
}

//新建prometheus实例
func NewPrometheus(addr string, svr *server.Server) (*Prometheus, error) {
	prom := new(Prometheus)
	prom.addr = addr
	prom.svr = svr

	golog.Info("prometheus", "Run", "Prometheus running", 0,
		"address",
		addr)

	return prom, nil
}

//启动prometheus的http监控
func (p *Prometheus) Run() {
	//设置标签及注册
	data := p.svr.GetMonitorData()

	for addr, item := range data {
		p.addNodeData(addr, item)
	}

	go p.flush()

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(p.addr, mux)

	if err != nil {
		golog.Error("prometheus", "Run", err.Error(), 0)
	}
}

func (p *Prometheus) addNodeData(addr string, data map[string]string) {
	var (
		nodeData = make(map[string]prometheus.Gauge)

		label = map[string]string{
			"addr": addr,
			"type": data["type"],
		}
	)

	nodeData["idleConn"] = p.addGauge("idleConn", "the db idle connection", label)
	nodeData["cacheConns"] = p.addGauge("cacheConns", "the db cache connection", label)
	nodeData["pushConnCount"] = p.addGauge("pushConnCount", "the db pushConnCount", label)
	nodeData["popConnCount"] = p.addGauge("popConnCount", "the db popConnCount", label)
	nodeData["maxConn"] = p.addGauge("maxConn", "the max connection config", label)

	p.nodesData.Store(addr, nodeData)
}

func (p *Prometheus) addGauge(name string, help string, label map[string]string) prometheus.Gauge {
	gauge := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:        name,
			Help:        help,
			ConstLabels: label,
		},
	)

	prometheus.MustRegister(gauge)

	return gauge
}

func (p *Prometheus) flush() {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		data := p.svr.GetMonitorData()

		for addr, item := range data {
			tmp, ok := p.nodesData.Load(addr)
			if !ok {
				p.addNodeData(addr, item)
				continue
			}

			nodeData, ok := tmp.(map[string]prometheus.Gauge)
			if !ok {
				continue
			}

			for name, gauge := range nodeData {
				if _, ok := item[name]; !ok {
					continue
				}

				val, _ := strconv.ParseFloat(item[name], 10)
				gauge.Set(val)
			}
		}
	}
}
