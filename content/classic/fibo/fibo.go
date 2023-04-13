package fibo

func FiboRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FiboRecursive(n-1) + FiboRecursive(n-2)
}

func FiboIterative(n int) int {
	if n <= 1 {
		return n
	}

	n1 := 1
	n2 := 0
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n1
}

func FiboMemoize(n int) int {
	cache := map[int]int{0: 0, 1: 1}
	return fiboMemoizeRecurse(n, cache)
}

func fiboMemoizeRecurse(n int, cache map[int]int) int {
	cached, hit := cache[n]
	if hit {
		return cached
	}
	n1 := fiboMemoizeRecurse(n-1, cache)
	n2 := fiboMemoizeRecurse(n-2, cache)
	f := n1 + n2
	cache[n] = f
	return f
}
