package main

import "fmt"

func AddOne(x *int) {
	*x = *x + 1
}

func main() {
	x := 10
	AddOne(&x)
	fmt.Println(x) // x is 11
}
