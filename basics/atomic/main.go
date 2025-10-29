package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	var a, b int32 = 0, 0

	go func() {
		a = 2
		atomic.StoreInt32(&b, 1)
	}()

	for {
		if n := atomic.LoadInt32(&b); n == 1 {
			fmt.Println(a)
			break
		}
		runtime.Gosched()
	}
}
