## Prime numbers

A **prime number** (or simply a **prime**) is a natural number greater than 1 that is not the product of two smaller natural numbers[^prime].

[^prime]https://en.wikipedia.org/wiki/Prime_number

### Trial division

To determine if a number is prime, a simple but slow method consists of testing if $n$ is divisible by any integer $2 < d < n$. Actually, the check can stop at $\sqrt(n)$ as they are new candidates beyond that value.

```go
func IsPrime(n int) bool {
	for i := 2; i * i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
```
