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

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/flike/kingshard/core/golog"
)

type Prometheus struct {
	addr string
}

//新建prometheus实例
func NewPrometheus(addr string) (*Prometheus, error) {
	prometheus := new(Prometheus)
	prometheus.addr = addr

	golog.Info("prometheus", "Run", "Prometheus running", 0,
		"address",
		addr)

	return prometheus, nil
}

//启动prometheus的http监控
func (p *Prometheus) Run() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(p.addr, mux)

	if err != nil {
		golog.Error("prometheus", "Run", err.Error(), 0)
	}
}