package main

import (
	"fmt"
	"sync"
)

func exec(id int, source chan int, apply func(int) int, out chan int) {
	for x := range source {
		fmt.Printf("exec %d got %d\n", id, x)
		out <- apply(x)
	}
}

func main() {
	source := make(chan int)
	out := make(chan int)
	go func() {
		defer close(source)
		for i := range 10 {
			source <- i
		}
	}()

	n := 3
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			exec(i, source, func(x int) int { return x * x }, out)
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for x := range out {
		fmt.Println(x)
	}
}
