package main

import (
	"context"
	"fmt"
	"time"
)

func transform(source <-chan int, f func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range source {
			out <- f(v)
		}
	}()

	return out
}

func filter(source <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range source {
			if pred(v) {
				out <- v
			}
		}
	}()

	return out
}

func generate(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				out <- i
				i++
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for v := range filter(transform(generate(ctx),
		func(v int) int {
			return -v
		}), func(v int) bool {
		return v%5 == 0
	}) {
		fmt.Println(v)
	}
}
