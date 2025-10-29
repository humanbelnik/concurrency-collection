package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type Mutex struct {
	// 1 occupied
	// 0 free
	occupied int32
}

func (m *Mutex) Lock() {
	// CAS(ptr, old, new):
	//	if *ptr != old: return false
	//\ *ptr = new
	// 	return true

	// If m.occupied == 1 (mutex is locked), CAS will return false, !CAS will be true and spin-lock will be executed.
	// Otherwise m.occupied will be set to 1 end spin-lock will be skipped -> lock acquired
	for !atomic.CompareAndSwapInt32(&m.occupied, 0, 1) {
		runtime.Gosched()
	}
}

func (m *Mutex) Unlock() {
	atomic.StoreInt32(&m.occupied, 0)
}

func main() {
	m := &Mutex{}
	wg := sync.WaitGroup{}
	i := 0
	for range 1000 {
		wg.Go(func() {
			m.Lock()
			defer m.Unlock()
			i++
		})
	}
	wg.Wait()
	fmt.Println(i)
}
