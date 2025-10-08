package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	outChan := make(chan int)
	go func() {
		for _, n := range nums {
			outChan <- n
		}
		close(outChan)
	}()
	return outChan
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(inChan <-chan int) <-chan int {
	outChan := make(chan int)
	go func() {
		for n := range inChan {
			outChan <- n * n
		}
		close(outChan)
	}()
	return outChan
}

func main() {
	// Sane return types so we can chain them like this
	// generator -> square -> print
	for v := range square(square(generator(2, 3, 4, 5, 6))) {
		fmt.Println(v)
	}
}
