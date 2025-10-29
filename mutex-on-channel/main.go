package main

import (
	"sync"
)

type Mutex struct {
	ch chan struct{}
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) Unlock() {
	m.ch <- struct{}{}
}
func main() {
	// m := &Mutex{}
	// wg := sync.WaitGroup{}
	// i := 0
	// for range 1000 {
	// 	wg.Go(func() {
	// 		m.Lock()
	// 		defer m.Unlock()
	// 		i++
	// 	})
	// }
	// wg.Wait()
	// fmt.Println(i)

	mu := sync.Mutex{}
	mu.Lock()
	mu.Lock()
}
