## Conditional statements

### If-else

Programs perform different computations or actions depending on whether a condition evaluates to true or false via a mechanism called **conditional statement**. We can express this in pseudo-code:

```
if condition:
    action1
    action2
else:
    action3
    action4
```

Notice the indent, as it is not accidental. The `action1` and `action2` will only be executed `if` the condition is met, a.k.a if it evaluates to `true`. Or `else`, `action3` and `action4` will be executed.

What would our `condition` actually be? We could for instance check that “a is geater than 10”, expressed as `if a > 10`.

There can be any number of instructions in the “if” and in the “else” block. Let’s see an example in Go:

```go
a := 10
b := 5
if isDivisible(a, b) {
	fmt.Printf("%d is divisible by %d\n", a, b)
} else {
	fmt.Printf("%d is not divisible by %d\n", a, b)
}
```

Provided the `isDivisible()` function is defined, the program would display “10 is divisible by 5” if 10 is actually divisible by 5 (which it is). The part of the program displaying “10 is not divisible by 5” would never be executed.

Go uses curly braces `{` and `}` to define a block of actions. In this case, each block contains a single statement that prints a message.

The *else* part is optional. We could simply do nothing if the “if” condition was not met.

### Example - coffee

Here is a more practical example, in pseudo-code:

```
if no coffee:
    exit house
    buy coffee
    enter house
pour coffee
drink coffee
```

In this example, we will pour the cofee and drink it whether we have some left in the first place. The condition `if no coffee` applies to the block represented by the 3 instructions below it, **indented**. So we will exit house, buy coffee and enter the house only _if_ there is no coffee left. The two last actions *pour coffee* and *drink coffee* are outside the “if” block.

Coffee is easy. Let’s combine it with bread.

```
if no bread or no coffee:
    exit house
    if no bread:
        buy bread
    if no cofee:
        buy cofee
    enter house
slice bread
pour coffee
eat bread
drink coffee
```

Several blocks can be constructed. Copy the following code in the [Go Playground](https://play.golang.org/) and run it with different values for `a` to get different output.

```go
a := 10
if a > 10 {
	fmt.Printf("%d is greater than 10\n", a)
} else if a < 0 {
	fmt.Printf("%d is less than 0\n", a)
} else {
	fmt.Printf("%d is between 0 and 10\n", a)
}
```

Blocks can also be nested:

```go
a := 1000
if a < 0 {
	fmt.Printf("%d is negative\n", a)
} else {
	if a > 100 {
		fmt.Printf("%d is a large number\n", a)
	}
}
```

Again, experiment with different values of `a`.

### Example - bakery

Let’s see a variation of the coffee examples.

```
if no bread:
    exit house
    timer <- 10 minutes
    while timer > 0
        search bakery
    if barkery found:
        buy bread
    enter house
if bread:
    slide bread
    eat bread
```

In this fictive example, we have two “main” chuncks: one where there is no bread, and one where there is.

In the first part, if there is no bread, we exit the house and start searching for a bakery. We have put a limit of 10 minutes on the “search bakery” action, expressed with a “*while*” loop, so that we do not end up roaming the streets endlessly -- because we may not have found a bakery after the time is up. Therefore, we put a condition “bakery found” on the “buy bread” action.

If we omitted this condition, and if there was no bakery found after 10 minutes, our program would attempt to buy bread out of thin air, and would *terminate abnormally*, a.k.a. it would crash to say things colloquially.

In any case, bakery found and bread found or not, we return to the house afterwards.

The second part handles the case where there is bread, in which case we slice it and eat it. This can happen in two cases:

* Either we have bread from the beginning. We did not have to step out of the house, search and buy bread.
* Or we had no bread upfront, we went out, found the bakery and bought the bread.

There is one case where we don’t slice and eat the bread: if we did not have any bread to start with, went out, did not find a bakery, and returned home without bread.

That’s it! We have covered all the cases, and hopefully we don’t run the streets forever, we don’t try to buy bread out of non-existing bakery and we don’t try to slice bread we don’t have. Failing to cover any of these cases would result in a **program failure**.

Note: we intentionally left the case where there is no bread in the bakery. It is a valid assumption for this illustration, as we are free to state that a bakery will always have bread in store. In real-life, you would have to take care of that case too.
