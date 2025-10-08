package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(workerName string, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("Worker %s processing %d\n", workerName, n)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup

	dispatchFromChan := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go dispatchFromChan(c)
	}

	// Close the output channel once all input channels are drained
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := generator(2, 3, 4, 5, 6)

	// TODO: fan out square stage to run two instances.
	ch1 := square("Worker 1", in)
	ch2 := square("Worker 2", in)

	// TODO: fan in the results of square stages.
	for n := range merge(ch1, ch2) {
		fmt.Println(n)
	}
}
