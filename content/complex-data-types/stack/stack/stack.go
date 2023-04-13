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
