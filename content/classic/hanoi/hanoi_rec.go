package main

import "fmt"

func solve(n int, source, auxiliary, destination string) {
	if n == 0 {
		return
	}
	solve(n-1, source, destination, auxiliary)
	fmt.Printf("Move from %s to %s\n", source, destination)
	solve(n-1, auxiliary, source, destination)
}

func main() {
	solve(4, "A", "B", "C")
}
