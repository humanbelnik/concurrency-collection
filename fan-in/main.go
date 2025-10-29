package main

import (
	"fmt"
	"sync"
)

func fanin(sources ...<-chan int) <-chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(sources))

	/*
		Read from each source until it's closed.
	*/
	for _, source := range sources {
		go func() {
			defer wg.Done()
			for x := range source {
				out <- x
			}
		}()
	}

	/*
		Not a single writer must be able to close 'out' channel.
		'out' must be closed when ALL sources are closed.
	*/
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()
		for i := range 100 {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()

	i := 0
	for range fanin(ch1, ch2, ch3) {
		i++
	}

	fmt.Println(i)

}
