package main

import "fmt"

func exec(source chan int, apply func(int) int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for x := range source {
			out <- apply(x)
		}
	}()

	return out
}

func main() {
	source := make(chan int)
	go func() {
		defer close(source)
		for i := range 10 {
			source <- i
		}
	}()

	for x := range exec(source, func(x int) int { return x * x }) {
		fmt.Println(x)
	}
}
