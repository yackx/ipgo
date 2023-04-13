package primes

func IsPrime(n int) bool {
	for i := 2; i * i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime6k1(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n % 2 == 0 || n % 3 == 0 {
		return false
	}
	for i := 5; i * i <= n; i += 6 {
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
	}
	return true
}

func Eratosthenes(n int) []int {
	// init sieve, all values to true
	sieve := make([]bool, n + 1)
	for i := range sieve {
		sieve[i] = true
	}

	// iteratively mark "not prime" the multiples of each prime
	for i := 2; i * i <= n; i++ {
		if sieve[i] {
			for j := i * i; j <= n;  j = j + i {
				sieve[j] = false
			}
		}
	}

	// extract prime numbers from the sieve
	primes := make([]int, 0)
	for p := 2; p < len(sieve); p++ {
		if sieve[p] {
			primes = append(primes, p)
		}
	}

	return primes
}
