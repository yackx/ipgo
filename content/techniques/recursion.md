## Recursion

Recursion is the technique of **making a function call itself**. It is a way to break down a problem into smaller and easier problems, until the problem becomes trivial. A function (or a procedure) that goes through recursion is said to be _recursive_.

The canonical use case for teaching recursion is the calculation of factorial.

$$n!=\prod _{i}^{n}i=1\times 2\times 3\times \ldots \times (n-1)\times n\quad n\in \mathbb {N}$$

By convention, factorial of $0$ is $1$ ; $0! = 1$.

For example, $5! = 120$.

$$5! = 5\times 4\times 3\times 2\times 1 = 120$$

The factorial of a number can be calculated with an **iterative** version and a loop construct you already know.

```go
func factorial(n uint) uint {
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}
```

Recursion can be used to solve this problem. Let's see how it works with a methaphore. Suppose we are asked to compute $fact(3)$. But that is too much work. We can multiply 2 numbers however. We also know Alice can calculate $fact(2)$. So we ask her to give us the value of $fact(2)$, and we multiply the result by $3$.

    3! = 3 * 2!  // ask Alice for 2!

How does Alice calculate $fact(2)$? She too can multiply 2 numbers. But she can't bothered to put too much thoughts into $fact(1)$. So she calls Bob. He knows how to calculate $fact(1)$. She will simply multiply by 2 the result given by Bob, $2 \times fact(1)$.

    2! = 2 * 1!  // Alice asks Bob for 1!

So far we have:

    3! = 3 * (2 * 1!)

Bob does not need to call anyone. He knows $fact(1) = 1$. He deals with the **terminal condition** or **base case**, upon which the solution stops recurring. We end up with the following stack of calls:

    3 * 2!
    3 * (2 * 1!)
    3 * (2 * (1 * 1))
    3 * (2 * 1)
    3 * 2
    6

In a computer program, we, Alice and Bob roles will be played by a single function. A recursive function. More formally, the recursive definition of factorial is:

    0! = 1.
    n > 0, n! = (n – 1)! × n.

The definition can be translated to Go:

```go
func factorial(n uint) uint {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

The general form of recursion is:

$$f(x) = e_{0} \text{ if } x \in D_{0}$$
$$f(x) = F(f(e_{1}), f(e_{2}), ..., f(e_{k}) \text{ if } x \in D_{v}$$

where:

- $f$ is the name of the function
- $e_{i}(i=0..k)$ are expressions that depend solely on $x$
- $F(y_{1}, ..., y_{k})$ is an expression that depends solely on $x$ and $(i=1..k)$
- $D_{0}$ and $D_{v}$ are disjoint sets

There is a cost to this approach. When we decompose the chain of calls made by Alice and friends, we realise how each recursive call cannot be evaluated before the whole chain completes. Each call maintains a record on the program stack. 

$$f(5) \rightarrow f(4) \rightarrow f(3) \rightarrow f(2) \rightarrow f(1) \rightarrow f(0)$$

The stack space needed for the recursive calls is proportional to the number $n$. We say the space complexity is $\mathcal{O}(n)$. Compare this to the iterative version, for which the space complexity is constant in $\mathcal{O}(1)$: no matter how big $n$ is, the program will not require to memorize the previous function calls. For small numbers, keeping calls on the stack has limited adverse effect. For larger numbers however, the stack may **overflow**.

A stack overflow can be avoided by using **tail recursion**, in which the recursive call is the _last call_. We use an `accumulator` to carry on the intermediate result. An **auxiliary** function is used to perform the actual recursion. The auxiliary function is not meant to be called directly.

```go
// Tail recursion. Not optimized in Go.
func factorial_rec(n, accumulator uint) uint {
    if n == 0 {
        return accumulator
    } else {
        return factorial_rec(n-1, n*accumulator)
    }
}

func factorial(n uint) uint {
    return factorial_rec(n, 1)
}
```

Unfortunately, like many imperative languages, Go does not optimize tail recursion, unlike most functional languages (Haskell, Scheme etc). But we wanted to briefly demontrate the technique nonetheless.

Arguably, this last version is less readable than the initial, iterative version. Although factorial is a traditional example to explain recursion, it may not be the preferred approach in practice. Recursion is however a key technique well suited to solve many problems, like Fibonacci numbers or the Tower of Hanoi. Some problems are way easier to solve with recursion.
