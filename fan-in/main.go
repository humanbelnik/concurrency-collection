package main

import (
	"fmt"
	"sync"
)

/*
Read from each source in a separate goroutine.

When closing out channel?
- When all sources will be closed
*/
func fanin(sources ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(sources))

	for _, source := range sources {
		go func() {
			defer wg.Done()
			for v := range source {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1, ch2, ch3 := make(chan int, 10), make(chan int, 10), make(chan int, 10)
	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()
		for i := range 10 {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()

	counter := 0
	for range fanin(ch1, ch2, ch3) {
		counter++
	}
	fmt.Println(counter)
}
