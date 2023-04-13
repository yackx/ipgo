## Numeric

Numeric types include _integer_, _floating-point_ and _complex_ types.

Consider the following program:

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

It computes and prints the sum of the squares of two _integers_.

$$a^2 +b^2, a = 3, b = 23$$
$$3^2 + 2^2 = 3.3 + 2.2 = 13$$

The function `sumOfSquares` deals with the `int` type to compute the sum of the squares of to integers. It makes sense. You can compute the square of `3` but not the square of an orange. `sumOfSquares` does not accept types other than `int`. Try to replace the main call by:

```go
sumOfSquares("hello", 2)
```

If we try to call the function with `"hello"` as a first argument, Go won’t even run our program. The compiler will analyze it and rightfully complain:

```
cannot use "hello" (untyped string constant) as int value in argument to sumOfSquares
```

It will not go any better if we use a floating point number:

```go
sumOfSquares(2.5, 2)
```

```
cannot use 2.5 (untyped float constant) as int value in argument to sumOfSquares
```

Next to the `int` type, we have discovered the `string` (more on that later) and the  `float` data types simply by looking at the compiler error message. There are actually [many more predeclared types](https://golang.org/ref/spec#Types)[^numeric-1] in Go. There is a granularity to integer types, ranging from 8 to 64 bits representation. An 8 bit integer can hold $2^8$ or 256 different values, as we have already seen.

[^numeric-1]: https://golang.org/ref/spec#Types

```
         the set of all...
uint8    unsigned  8-bit integers (0 to 255)
uint16   unsigned 16-bit integers (0 to 65535)
uint32   unsigned 32-bit integers (0 to 4294967295)
uint64   unsigned 64-bit integers
           (0 to 18446744073709551615)
````

Next to these types that can hold **unsigned** integers, there are corresponding **signed** integers. They have the same number of bytes, can hold as many different values, only the range differ. For instance, while the value of an unsigned integer `uint8` range from $0$ to $255$, a signed integer `int8` hold values ranging from $-128$ to $+127$.

```
         the set of all...
int8     signed  8-bit integers (-128 to 127)
int16    signed 16-bit integers (-32768 to 32767)
int32    signed 32-bit integers (-2147483648 to 2147483647)
int64    signed 64-bit integers
           (-9223372036854775808 to 9223372036854775807)
```

Type `int` alone is an odd duck, as it will be either 32 or 64 bits, depending on your platform.

The situation is slightly simpler with "non integers non complex" numbers, as there are only two types `float32` and `float64` IEEE-754 32-bit floating-point representations. You can represent very large and very small numbers (in absolute value) but… with a finite precision. This is not a Go-specific limitation as floating-point aritmethic is very common among programming languages.

**WARNING**: Never use a floating point values to manipulate financial amounts, whatever the language.
