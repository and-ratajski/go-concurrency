package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int) {
	for i := 0; i < 18; i++ {
		fmt.Println("Sending:", i)
		ch <- i
	}
}

func receiver(ch <-chan int) {
	for v := range ch {
		fmt.Println("Received:", v)
	}
}

func main() {
	ch := make(chan int, 3)

	go sender(ch)
	go receiver(ch)

	time.Sleep(time.Second * 1)
}
