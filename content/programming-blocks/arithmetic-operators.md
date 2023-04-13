## Arithmetic operators

### Standard operators

[Arithmetic operators](https://golang.org/ref/spec#Arithmetic_operators) apply to numeric values and yield a result of the same type as the first operand. The four standard arithmetic operators (+, -, \*, /) apply to integers and floating-points.

---- -------------- -----------------------------------------
`+`  sum            integers, floats, complex values, strings
`-`  difference     integers, floats, complex values
`*`  product        integers, floats, complex values
`/`  quotient       integers, floats, complex values
`%`  remainder      integers
---- -------------- -----------------------------------------

You are familiar with the standard arithmetic operators. However the remainder operator may be new to you. If we divide 7 by 2, the result will be 3, and the remainder will be 1, because `7 == 2*3 + 1`.

```
7 / 2
3

7 % 2
1
```

It gets tricky with negative numbers; refer to the Go documentation and experiment if you want to know more.

### Concatenation

The `+` operator can also be applied to two strings: `"Hello, " + "world!"`. In this case, rather than being *added*, the strings are *concatenated*.

### Bitwise and shift

In a later chapter, we will present _bitwise logical operators_ and _shifts_, after we cover binary representation.
