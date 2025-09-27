package main

func fanout(source <-chan int, nsplits int) []chan<- int {
	outs := make([]chan<- int, nsplits)
	for i := range nsplits {
		outs[i] = make(chan<- int)
	}

	go func() {
		i := 0
		for v := range source {
			outs[i%nsplits] <- v
			i++
		}
		for _, out := range outs {
			close(out)
		}
	}()

	return outs
}

func main() {

}
