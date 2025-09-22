package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("simple go routine")

	// goroutine with anonymous function
	go func() {
		fun("goroutine with anonymous function")
	}()

	// goroutine with function value call
	fv := fun
	go fv("goroutine with function value call")

	// wait for goroutines to end

	time.Sleep(50 * time.Millisecond) // We need to wait a bit to let goroutines start
	fmt.Println("done..")
}
