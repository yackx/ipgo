package main

import (
	"fmt"
	"time"
)

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func main() {
	go foo()
	fmt.Println("Terminated")
	time.Sleep(time.Millisecond * 500)
}
