package main

import (
	"runtime"
	"testing"
	"time"
)

func BenchmarkAllocationEveryMs(b *testing.B) {
	// need permanent allocation to clear see when the heap double its size
	var s *[]int
	tmp := make([]int, 1100000, 1100000)
	s = &tmp

	var a *[]int
	for i := 0; i < b.N; i++ {
		tmp := make([]int, 10000, 10000)
		a = &tmp

		time.Sleep(time.Millisecond)
	}
	_ = a
	runtime.KeepAlive(s)
}
