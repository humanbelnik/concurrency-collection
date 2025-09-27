package main

import "fmt"

func fanout(source <-chan int, nsplits int) []chan string {
	outs := make([]chan string, nsplits)
	for i := range nsplits {
		outs[i] = make(chan string)
	}

	go func() {
		for v := range source {
			for n := range nsplits {
				outs[n] <- fmt.Sprintf("value %d goes into stream %d", v, n)
			}

		}
		for _, out := range outs {
			close(out)
		}
	}()

	return outs
}

func main() {
	source := make(chan int, 10)
	go func() {
		for i := range 120 {
			source <- i
		}
		close(source)
	}()

	outs := fanout(source, 3)
	for {
		select {
		case v, ok := <-outs[0]:
			if !ok {
				outs[0] = nil
				continue
			}
			fmt.Println(v)
		case v, ok := <-outs[1]:
			if !ok {
				outs[1] = nil
				continue
			}
			fmt.Println(v)
		case v, ok := <-outs[2]:
			if !ok {
				outs[2] = nil
				continue
			}
			fmt.Println(v)
		default:
			if outs[0] == nil &&
				outs[1] == nil &&
				outs[2] == nil {
				return
			}
		}

	}
}
