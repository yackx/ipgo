## Loops

A **loop** repeats an action or a set of actions *while* a condition is met.

Sometimes, it can be a negative condition, e.g. "while countdown *is not* 0". A program often combines loops and conditional statements, like the coffee example we just saw.

In Go, all loops are performed using the `for` instruction. The following program count downs from 10 to 0.

```go
n := 10
for n >= 0 {
	fmt.Println(n)
	n = n - 1
}
```

On each step (called **iteration**), the program checks if the condition `n >= 0` (*is n greater or equal to 0?*) holds. If it does, it executes the instructions in the “for” block, that is it prints the current value of the counter and decrements it.

A general form of a `for` statement is the following:

```
"for" [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] Block
```

where

* `InitStmt` is an initialization statement, performed once before starting the loop;
* `Condition` is the loop condition that must hold true to continue to iterate, for instance `i < 10`;
* `PostStmt` is a statement executed after each execution of the loop. Typically we increment a counter.

For example, the following program will compute the sum of all integers from 0 to 10 excluded, and print it.

```go
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
fmt.Println(sum)
```

* Init statetment is `i := 0`. It is executed once;
* End condition is `i < 10`, aka we will perform the loop while `i` is less than 10;
* After each execution of the block, `i` is increased by 1 (**incremented**)  with the instruction `i++`, a shortcut for `i = i + 1`.
