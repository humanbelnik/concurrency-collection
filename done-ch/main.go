package main

import (
	"fmt"
	"time"
)

func worker(done <-chan struct{}, res *int) (finishCh chan struct{}) {
	finishCh = make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("res on done:", *res)
				func() {
					defer close(finishCh)
					for *res < 15 {
						(*res)++
					}
				}()
			default:
				time.Sleep(time.Second)
				(*res)++
			}
		}
	}()

	return finishCh
}

func main() {
	done := make(chan struct{})
	res := 0

	finishCh := worker(done, &res)

	time.AfterFunc(time.Second*2, func() {
		done <- struct{}{}

	})

	/*
		Unblocks after close op.
		Will be closed after gracefull actions are done.
	*/
	<-finishCh

	fmt.Println("finish", res)

}
