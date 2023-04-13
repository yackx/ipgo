package fibo

import "testing"

func TestFiboRecursive(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	for i, w := range want {
		actual := FiboRecursive(i)
		if w != actual {
			t.Errorf("FiboRecursive(%d) want=%d actual=%d\n", i, w, actual)
		}
	}
}

func TestFiboIterative(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	for i, w := range want {
		actual := FiboIterative(i)
		if w != actual {
			t.Errorf("FiboIterative(%d) want=%d actual=%d\n", i, w, actual)
		}
	}
}

func TestFibo(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	functions := []func(int) int{FiboIterative, FiboRecursive, FiboMemoize}
	for _, f := range functions {
		for i, w := range want {
			actual := f(i)
			if w != actual {
				t.Errorf("(%d) want=%d actual=%d\n", i, w, actual)
			}
		}
	}
}
