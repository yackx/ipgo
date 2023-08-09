## Functions

Suppose we want to display the sum of the squares of any two numbers.

$$f(a, b) \rightarrow a^2 + b^2$$

We could do the following:

```go
func main() {
    a := 2
    b := 3
    fmt.Println(a*a + b*b)
    a := 5
    b := 4
    fmt.Println(a*a + b*b)
}
```

Observe that we are repeating ourselves. Instead of writing the same formula `a*a + b*b` over and over again, we could write a reusable **function**. We have already used functions, namely the one called `main`. It is easy to create our own. We shall call it `sumOfSquares`:

```go
func sumOfSquares(a, b int) int {
    return a*a + b*b
}
```

We say `sumOfSquares` is a function that take takes two **parameters** `a` and `b`. They are of **type** `int`, which stands for *integer*, and it returns another *integer*. Indeed, as our **arguments** `2` and `3` are integer values. The function **returns** another integer, which the value that was computed (the sum of squares). We can invoke it and print its output:

```go
func sumOfSquares(a, b int) int {
	return a*a + b*b
}

func main() {
	fmt.Println(sumOfSquares(2, 3))
	fmt.Println(sumOfSquares(5, 4))
}
```

With a predictable result:

```
13
41
```

Our function `sumOfSquares` is reusable. Our `main` methods neately calls it twice, hiding the calculation details, promoting reusability and providing a higher level of abstraction.
