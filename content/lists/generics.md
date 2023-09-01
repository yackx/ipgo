## Generics

Let's revisit our `sumOfSquare` function.

```go
package main

import "fmt"

func sumOfSquares(a, b int) int {
    return a*a + b*b
}

func main() {
        fmt.Println(sumOfSquares(3, 2))
}
```

We have seen that it won't work with a type other than a `int`. It makes sense not to square another type like `string`, but we could perfectly want to compute the square of two floating-point numbers.

Go allows to define a **generic** function[^go-generics] that accepts different types. We can modify `sumOfSquares` to accept `int`, `float32` and `float64`.

[^go-generics]: https://go.dev/doc/tutorial/generics#add_generic_function

First, we declare a type `Number` with constraints:

```go
type Number interface {
	int | float32 | float64
}
```

Then, we make `sumOfSquare` generic by allowing it to accept any `Number` symbolized by `T`, and return the same type `T`.

```go
func sumOfSquares[T Number](a, b T) T {
	return a*a + b*b
}
```

We can now call it with different numeric types.

```go
fmt.Println(sumOfSquares(1, 2))
fmt.Println(sumOfSquares(1.5, 3.2))
```

Remark: we cannot call our function with mixed types, for instance `sumOfSquare(2.5, 2)`. The following error occurs: `default type int of 2 does not match inferred type float64 for T`. The second argument must be turned into a `float64` to be compatible with the first one:

```go
sumOfSquares(2.5, float64(2))
```
