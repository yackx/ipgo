## Expressions

### Examples

An intuitive way to start programming is to have a look at expressions. `"Hello, world!"` was used in our first program is an expression. `42` is another one. `3+4` a third. The basic arithmetic operations can be performed with `+`, `-`, `\*` and `/`. You can perform complex operations, with parenthesis if you need. The `*` and `/` operators take precedence over `+` and `-`.

We can go further. All of the following are expressions:

```go
sumOfSquare(2, 3)
sumOfSquare(2 * 4, 3 * 2)
sumOfSquare(sumOfSquare(2, 3), 3 * 2)
```

### Definitions

Expressions in Go are formally defined in the [language specification](https://golang.org/ref/spec#Expressions). In the **expression** `7+2`, `+` is the **operand**. `7` and `2` are its **operators**. The operator’s **arity** is 2.

Expression:

> An expression specifies the computation of a value by applying operators and functions to operands.

Operand:

> Operands denote the elementary values in an expression.

Arity:

> The number of arguments or operands that a function takes.

Binary operator:

> An operator that applies to two operands. Its **artity** is 2.

Unary operator:

> An operator that applies to one operand. The expression `-7` has an operator `-` (negate) with one operand `7`. Its **arity** is 1. In this case, the operator `-` (of arity 1) is not to be confused with the minus mathematical operator (of arity 2).

### Evaluation

Let’s take a non trivial yet simple case and evaluate `sumOfSquare`.

```
sumOfSquare(2, 3)
(2 * 2) + (3 * 3)
4 + (3 * 3)
4 + 9
13
```

**NOTE**: Notice the use of parenthesis. Although mathematically not mandatory in this case, as multiplation takes precendence over addition, they denote a group to be evaluated regardless of arithmetic precedence considerations.

In order to evaluate `sumOfSquare(2 * 4, 3 * 2)`, `2*4` and `3*2` must be evaluated first. Then we fall back on the case of `sumOfSquare` with two integers as parameters we already know.

```
sumOfSquare(2 * 4, 3 * 2)
sumOfSquare(8, 3 * 2)
sumOfSquare(8, 6)
(8 * 8) + (6 * 6)
64 + 36
100
```

### Boolean expressions

A Boolean expression is **a logical statement that is either true or false**.

The expression `3 < 5` is evaluated as `true` while `2 < 0` is evaluated as `false`. You can assign a boolean expression to a variable, for instance `a = 3 < 5`. In that case, `a` will be evaluated to `true`.

You can test two values for equality with `==` as in `3 == 3` (which of course is `true`). The double equal is used to avoid confusion with a variable assignment, as in `a = 3`. Which can be confusing in itself. Many languages use `==` to perform equality comparisons. If you want to check that two expressions are different, as in "\neq", you would use `!=`, like so: `3 != 5` which evaluates to `true` as 3 is not equal to 5.

--------------------------------------
Math       Go       Meaning
---------  -------  ------------------
>          ``>``    Greater than

≥          ``>=``   Greater or equal

<          ``<``    Less than

≤          ``<=``   Less than or equal

=          ``==``   Equal

\neq       ``!=``   Not equal 
--------------------------------------
