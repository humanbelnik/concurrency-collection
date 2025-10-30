package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 0
	wg := sync.WaitGroup{}
	for range 1000 {
		wg.Go(func() {
			n++
		})
	}
	wg.Wait()
	fmt.Println(n)
}
