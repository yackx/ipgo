package primes

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	data := []struct {
		n      int
		wanted bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{97, true},
		{100, false},
	}

	for _, d := range data {
		actual := IsPrime(d.n)
		if actual != d.wanted {
			t.Errorf("IsPrime(%d) incorrect (%v)\n", d.n, actual)
		}
	}
}

func TestIsPrime6k1(t *testing.T) {
	data := []struct {
		n      int
		wanted bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{97, true},
		{100, false},
	}

	for _, d := range data {
		actual := IsPrime6k1(d.n)
		if actual != d.wanted {
			t.Errorf("IsPrime(%d) incorrect (%v)\n", d.n, actual)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	_ = IsPrime(2880655109)
}

func BenchmarkIsPrime6k1(b *testing.B) {
	_ = IsPrime6k1(2880655109)
}

func TestEratosthenes(t *testing.T) {
	wanted := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	actual := Eratosthenes(30)

	if !reflect.DeepEqual(wanted, actual) {
		t.Errorf("Eratosthenes wanted=%v actual=%v\n", wanted, actual)
	}
}