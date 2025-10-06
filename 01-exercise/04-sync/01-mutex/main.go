package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int32
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int32) {
		//mu.Lock()
		//balance += amount
		//mu.Unlock()
		atomic.AddInt32(&balance, amount)
	}

	withdrawal := func(amount int32) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount

	}

	// make 100 deposits of $1
	// and 100 withdrawal of $1 concurrently.
	// run the program and check result.

	// TODO: fix the issue for consistent output.

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
