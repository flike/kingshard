// Copyright 2015 The kingshard Authors. All rights reserved.
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

package backend

import (
	"math/rand"
	"time"

	"github.com/flike/kingshard/core/errors"
)

func Gcd(ary []int) int {
	var i int
	min := ary[0]
	length := len(ary)
	for i = 0; i < length; i++ {
		if ary[i] < min {
			min = ary[i]
		}
	}

	for {
		isCommon := true
		for i = 0; i < length; i++ {
			if ary[i]%min != 0 {
				isCommon = false
				break
			}
		}
		if isCommon {
			break
		}
		min--
		if min < 1 {
			break
		}
	}
	return min
}

func (n *Node) InitBalancer() {
	var sum int
	n.LastSlaveIndex = 0
	gcd := Gcd(n.SlaveWeights)

	for _, weight := range n.SlaveWeights {
		sum += weight / gcd
	}

	n.RoundRobinQ = make([]int, 0, sum)
	for index, weight := range n.SlaveWeights {
		for j := 0; j < weight/gcd; j++ {
			n.RoundRobinQ = append(n.RoundRobinQ, index)
		}
	}

	//random order
	if 1 < len(n.SlaveWeights) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < sum; i++ {
			x := r.Intn(sum)
			temp := n.RoundRobinQ[x]
			other := sum % (x + 1)
			n.RoundRobinQ[x] = n.RoundRobinQ[other]
			n.RoundRobinQ[other] = temp
		}
	}
}

func (n *Node) GetNextSlave() (*DB, error) {
	var index int
	if len(n.RoundRobinQ) == 0 {
		return nil, errors.ErrNoDatabase
	}
	if len(n.RoundRobinQ) == 1 {
		index = n.RoundRobinQ[0]
		return n.Slave[index], nil
	}

	queueLen := len(n.RoundRobinQ)
	index = n.RoundRobinQ[n.LastSlaveIndex]
	db := n.Slave[index]
	n.LastSlaveIndex++
	if queueLen <= n.LastSlaveIndex {
		n.LastSlaveIndex = 0
	}
	return db, nil
}
