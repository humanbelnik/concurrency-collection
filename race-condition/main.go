package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := range 10 {
		wg.Go(func() {
			fmt.Println(i)
		})
	}
	wg.Wait()
}
