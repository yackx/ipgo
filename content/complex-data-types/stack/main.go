package main

import (
	"fmt"

	"be.sugoi.ipgo/stack"
)

func main() {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for s.Len() > 0 {
		elem, ok := s.Pop()
		if !ok {
			panic("stack is empty")
		}
		fmt.Println(elem)
	}
	fmt.Println(s.Pop()) // 0 false
}
