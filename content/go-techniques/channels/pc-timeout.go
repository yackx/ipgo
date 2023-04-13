package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(items chan<- int) {
	i := 0
	for {
		// Pretend to produce an item
		sleepTime := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(sleepTime)

		// Send the item on the channel
		fmt.Println("Produced", i)
		items <- i

		i++
	}
}

func consumer(items <-chan int) {
	for item := range items {
		// Pretend to consume the item
		sleepTime := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(sleepTime)
		fmt.Println("Consumed", item)
	}
}

func main() {
	items := make(chan int)

	// Launch the producer and consumer goroutines
	go producer(items)
	go consumer(items)

	// Wait for 5 seconds
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout reached, exit")
	}
}
