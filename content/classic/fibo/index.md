## Fibonacci numbers

The Fibonacci numbers ($F_{n}$) form a sequence called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from $0$ and $1$.

$$
    n > 1:
    F_{0} = 0,
    F_{1} = 1,
    F_{n} = F_{n-1} + F_{n-2}
$$

The sequence starts with

$$0 , 1 , 1 , 2 , 3 , 5 , 8 , 13 , 21 , 34 , 55 , 89 , 144$$

### Recursive method

It is possible to translate the mathematical definition into Go code without much effort.

```go
func Fib(n int) int {
    if n <= 1 {
        return n
    }
    return Fib(n - 1) + Fib(n - 2)
}
```

This implementation is trivial, albeit very inefficient. Let's visualise the recursive calls for $Fib(4)$.

$$\begin{array}{rccccl}
    \; & \; & \; & Fib(4) \; & \; & \\
    \; & \; & \swarrow & \; & \searrow & \; \\
    \; & \; & Fib(3) & \; & Fib(2) & \; \\
    \; & \swarrow & \downarrow & \; & \; \downarrow & \searrow \\
    \; & Fib(2) & Fib(1) & \; \; & Fib(1) & Fib(0) \\
    \swarrow & \downarrow & \; & \; & \; & \; \\
    Fib(1) & Fib(0) & \; & \; & \; & \; \\
\end{array}$$

Even for a simple case where $n$ is low, there is a lot of repetition. The same $Fib(n)$ is evaluated several times. $Fib(1)$ is computed 3 times. It only gets worse as $n$ becomes bigger. A program cannot "remember" the previously calculated values if it has not been written to do so, and will therefore performs the recursive calls as many time as needed.

> Exercise: draw the top of the recursive calls tree for $Fib(5)$.

### Memoization

Memoization is an **optimization technique** used primarily to speed up programs by **storing the results of expensive computation** and returning the cached result when the same inputs occur again[^memoize].

[^memoize]: https://en.wikipedia.org/wiki/Memoization

In order to _memoize_ the `Fib()` function, our cache will be of type `map[int]int`, where each key is a number, and each value the corresponding Fibonacci number. So as $Fib(6) = 8$, the cache will be populated with `{6: 8}` _the first time_ this number is calculated.

We start with the trivial cases `{0: 0, 1: 1}` to populate our cache.

We want to keep the function signature tidy. Therefore, we do not want any concept of memoization to appear to the calling program. `FibMemoize()` will initialize the cache, and call the auxiliary function `fibMemoizeRecurse()`, which does the actual recursive work, passing our cache along.

When we enter `fibMemoizeRecurse()`, we first check if we have already computed the number by accessing our cache. In case of a hit, we return the cached value directly. Otherwise, we perform the recursive calls for $F(n-1)$ and $F(n-2)$, passing the cache along.

Before we return the computed number, we store it in the cache, so other callers can benefit from it.

```go
func FibMemoize(n int) int {
	cache := map[int]int{0: 0, 1: 1}
	return fibMemoizeRecurse(n, cache)
}

func fibMemoizeRecurse(n int, cache map[int]int) int {
	cached, hit := cache[n]
	if hit {
		return cached
	}
	n1 := fibMemoizeRecurse(n-1, cache)
	n2 := fibMemoizeRecurse(n-2, cache)
	f := n1 + n2
	cache[n] = f
	return f
}
```

The recursive call tree is greatly simplified. In the following diagram, the cache hits are denoted by a $*$. It might not look much on a small tree like $Fib(4)$, but with larger numbers, entire branches of the tree are pruned.

$$\begin{array}{rccccl}
    \; & \; & \; & Fib(4) \; & \; & \\
    \; & \; & \swarrow & \; & \searrow & \; \\
    \; & \; & Fib(3) & \; & Fib(2)* & \; \\
    \; & \swarrow & \downarrow & \; & \; & \; \\
    \; & Fib(2) & Fib(1)* & \; & \; & \; \\
    \swarrow & \downarrow & \; & \; & \; & \; \\
    Fib(1) & Fib(0) & \; & \; & \; & \; \\
\end{array}$$

> Exercise: extend the recursive calls tree above for $Fib(5)$ with memoization.

### Iterative method

The Fibonacci sequence can be comptuted with an iterative approach. We loop over integers if $n$ is greater than $2$. We retain only the two previous numbers $F(n-1)$ and $F(n-2)$ as we progress.

Notice the use of an interesting Go property: `a, b = b, a` will swap the values of the variables `a` and `b`. This is _not_ equivalent to `a = b; b = a` as we would loose the value of `a` after the first assignment.

```go
func Fibo(n int) int {
    if n <= 1 {
        return n
    }

    n1 := 1
    n2 := 0
    for i := 2; i <= n; i++ {
        n2, n1 = n1, n1 + n2
    }

    return n1
}
```

There is no growing stack usage as $n$ increases.

### Related

Fibonacci numbers are used in the Fibonacci search technique, and can form a tiling.

![Tiling pattern using Fibonacci numbers](content/classic/fibo/FibonacciBlocks.png)
