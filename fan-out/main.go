package main

import "fmt"

func fanout(source <-chan int, predicate func(int) bool) [2]chan int {
	outs := [2]chan int{
		make(chan int),
		make(chan int),
	}

	go func() {
		defer func() {
			for _, out := range outs {
				close(out)
			}
		}()
		for x := range source {
			if predicate(x) {
				outs[0] <- x
			} else {
				outs[1] <- x
			}
		}
	}()

	return outs
}

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range 100 {
			ch <- i
		}
	}()

	chans := fanout(ch, func(v int) bool {
		if v%2 == 0 {
			return true
		}
		return false
	})

	i := 0
	for chans[0] != nil || chans[1] != nil {
		select {
		case x, ok := <-chans[0]:
			if !ok {
				chans[0] = nil
				continue
			}
			i++
			fmt.Println("from 0:", x)
		case x, ok := <-chans[1]:
			if !ok {
				chans[1] = nil
				continue
			}
			i++
			fmt.Println("from 1:", x)
		}
	}
	fmt.Println("i", i)
}
