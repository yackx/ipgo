## Stack

A stack is an abstract data type that serves as a collection of elements[^stack-wikipedia], with 2 main operations: 

- **push** adds an element to the collection
- **pop** removes the most recently added element from the collection

And optional operations:

- **len** returns the number of elements in the collection
- **peek** returns the most recently added element (without removing it)

It is a **LIFO** (last in, first out) data structure.

Stacks are used for implementing function calls, storing program variables, parsing expressions, evaluating arithmetic expressions, ...

Go includes a package that implements a stack[^go-stack]. It uses an optimized `type` that contains a node, chained with pointers. It maintains a `length` to avoid navigating the chain when `Len()` is called, for performance reasons.

### A simple implementation

As an exercise, we can implement a simpler stack. We will use:

- a slice to store the values in the collection
- a generic type `T` (rather than the legacy `interface{}` in the Go library)
- and all functions will operate on a `*Stack[T]`.

This translates to:

```go
package stack

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) Pop() (elem T, ok bool) {
	if s.Len() == 0 {
		return elem, false
	}
	index := len(*s) - 1
	elem = (*s)[index]
	*s = (*s)[:index]
	return elem, true
}
```

Remarks:

- We rely on `New()` to create a `*Stack` rather than forcing the client to declare a slice.
- `Pop()` returns the element (if found) and a flag. This is a common and convenient way to return an possibly missing value in Go.

Let's add a main package to demonstrate the stack capabilities, and how to pop a value.

```go
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
```

Refer to the appendix "Modules" for instructions on how to store those files and execute the program.

### Exercise

1. Make this code run on your computer.

2. Add a `Peek()` function. Test that it does not remove the element it returns.

[^stack-wikipedia]: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
[^go-stack]: https://pkg.go.dev/github.com/golang-collections/collections/stack
