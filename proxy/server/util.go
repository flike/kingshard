package server

import "sync/atomic"

type BoolIndex struct {
	index int32
}

func (b *BoolIndex) Set(index bool) {
	if index {
		atomic.StoreInt32(&b.index, 1)
	} else {
		atomic.StoreInt32(&b.index, 0)
	}
}

func (b *BoolIndex) Get() (int32, int32, bool) {
	index := atomic.LoadInt32(&b.index)
	if index == 1 {
		return 1, 0, true
	} else {
		return 0, 1, false
	}
}
