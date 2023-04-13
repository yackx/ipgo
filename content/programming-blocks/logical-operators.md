# Logical operators

## Conjonction (AND)

A boolean expression does not have to live in isolation. Consider we want to answer the following question:

> *Is it sunny and warm?*

This boolean expression would evaluate to *true* if and only if it is both sunny and warm at the same time. If it is sunny but cold, the result would be *false*. Same if it is rainy and warm and, of course, if it is rainy and cold.

To generalize, an expression “*p and q*” will evaluate to *true* if both *p* and *q* are *true*, and will evaluate to *false* otherwise.

Go does not use “and” as we do in English. Instead, it has an operator `&&`. This is for historical reason and it common in many programming languages, although some of them simply use `and`, where it becomes a reserved keyword.

Given “*p and q*”, written `p && q` in Go or $p \wedge q$ in logic notation, with two boolean possibilities `true` and `false` for `p` and for `q`, we can construct the following **table of truth** for the **and** logical operation (noted ⋀):

| p

 | q

 | p ⋀ q

 |
| true

     | true

                                                                                | true

               |
| true

     | false

                                                                               | false

              |
| false

    | true

                                                                                | false

              |
| false

    | false

                                                                               | false

              |
Each *p* and *q* can itself be a bit more elaborate, for instance:

> *Is 2 less than 5 and is 3 less than 10?*

The answer would be “yes”, or `true`, because indeed, 2 is less than 5, *and* 3 is less than 10 ; both are true.

That would translate into the following expression with its evaluation:

```
2 < 5 and 3 < 10
true and true
true
```

Suppose we would slightly change our initial question to become:

> *Is 2 greater than 5 and is 3 less than 10?*

You can convience yourself the answer is “no” or `false`, because although 3 is still less than 10, 2 is not greater than 5. The first condition fails, so the whole statement no longer holds. If we decompose, we would have:

```
2 > 5 and 3 < 10
false and true
false
```

## Disjunction (OR)

Replace “and” seen above by “or”, we want to test if any of two expressions is true - or if both a true at the same time. The original question would become:

> *Is it either sunny or warm?*

We want to test “p or q”, written `p || q` in Go or $p \vee q$ in logic notation. What does our table of truth become?

| p

        | q

                                                                                   | p ⋁ q

              |
| true

     | true

                                                                                | true

               |
| true

     | false

                                                                               | true

               |
| false

    | true

                                                                                | true

               |
| false

    | false

                                                                               | false

              |
There is only one case where `p || q` is `false`: when both `p` and `q` are `false`, so when it is neither sunny not warm.

## Negation (NOT)

Also called *logical complement*, it is an operation that takes a proposition *p* to another proposition “*not p*”, written $\neg p$ in logic or `!p` in Go.

Intuitively, it turns `true` into `false` and vice-versa. So if *p* is “John is 20 years old or older”, then *¬p* is “John is younger than 20 years old”.

| p

        | ¬p

                                                                                  |
| true

     | false

                                                                               |
| false

    | true

                                                                                |
## Short-circuit evaluation

In logic and in math, the “AND” and the “OR” operations are **commutative**, meaning you can swap p and q in our examples.

p \\wedge q \\equiv q \\wedge p

p \\vee q \\equiv q \\vee pIn Go, as in many programming languages, logical expressions are evaluated left to right and are tested for possible “short-circuit” evaluation using the following rules:


* `(some falsy expression) && expr` is short-circuit evaluated to the falsy expression


* `(some truthy expression) || expr` is short-circuit evaluated to the truthy expression.

Short circuit means that the `expr` in both samples above is **not evaluated**. The program acts as if it was lazy. Suppose our question was: *is it both sunny and warm?* If you look outside and the sun is nowhere to be seen, then the answer to the question is “no”. There is no need to check for the temperature, as it won’t change the final answer, whether it’s warm or not. Conversly, if our question was: *is it either sunny or warm?*, and if the sun was shining, you could answer “yes” straight away, without reading the temperature.

Therefore any **side effect** the shortcut expression on the right-hand side might have will not occur. Suppose we want to evaluate an expressions using a new function `divide`:

```go
func divide(a, b int) int {
    return a / b
}

fmt.Println((divide(5, 2) == 10) && (divide(5, 1) == 10))
```

Actually, Go will only evaluate the left part `(divide(5, 2) == 10)` to `false`. Remember our table of truth: if the left part of the `&&` expression is `false`, the resulting evaluation for the whole expression is bound to be `false` as well. So why waste CPU cycles to evaluate `(3 < 10)`, as any other outcome is impossible?

Consider this revisited example:

```go
fmt.Println((divide(5, 2) == 10) && (divide(5, 0) == 10))
```

Any attempt to divide 5 by 0 would produce an error. But this expression will simply evaluate to `false`, although there is an error waiting to happen on the right side, because the condition on the left side (5/2 == 10?) if false. The expression on the right will “benefit” from shortcut evaluation and will simply be skipped. Hence the error will not happen. Now turn the expression around:

```go
fmt.Println((divide(5, 0) == 10) && (divide(5, 2) == 10))
```

This will produce an error! This time, the program will first evaluate “5/0 == 10?” and fail.

```
panic: runtime error: integer divide by zero
```

Because of shortcut evaluation, the two expressions `p&&q` and `q&&p` are therefore **not equivalent** in Go when `p` or `q` produces a side-effect, such as generating an runtime error.

The same mechanism takes place with `||`. If the left side is `true`, Go will be satisfied and won’t bother evaluating the right side, because the end result is bound to be `true`.
