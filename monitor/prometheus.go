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
	data := p.svr.GetMonitorData()

	for addr, data := range data {
		label := make(map[string]string)

		label["addr"] = addr
		label["type"] = data["type"]

		idleConn 	:= data["idleConn"]
		maxConn 	:= data["maxConn"]

		idleConnGauge := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "idleConn",
				Help: "the db idle connection",
				ConstLabels: label,
			},
		)
		prometheus.MustRegister(idleConnGauge)

		maxConnGauge := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "maxConn",
				Help: "the max connection config",
				ConstLabels: label,
			},
		)
		prometheus.MustRegister(maxConnGauge)

		go func() {
			for {
				idleConnValue, _ := strconv.ParseFloat(idleConn, 10)
				idleConnGauge.Set(idleConnValue)
				time.Sleep(5 * time.Second)
			}
		}()

		go func() {
			for {
				maxConnValue, _ := strconv.ParseFloat(maxConn, 10)
				maxConnGauge.Set(maxConnValue)
				time.Sleep(5 * time.Second)
			}
		}()
	}

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(p.addr, mux)

	if err != nil {
		golog.Error("prometheus", "Run", err.Error(), 0)
	}
}