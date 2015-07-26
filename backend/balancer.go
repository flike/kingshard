package backend

import (
	. "github.com/flike/kingshard/core/errors"
	"math/rand"
	"time"
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
		return nil, ErrNoDatabase
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
