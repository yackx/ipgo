# Miscelaneous

## Read user input

User input can be read from the keyboard using functions from the `os` and the `bufio` packages.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello " + text)
}
```

`os.Stdin` represents the **standard input**, namely the keyboard (or equivalent device), if `stdin` has not been redirected from another source. The mechanism is to obtain a “reader” from “stdin”, and from that reader, we read a string. A string is terminated by a new line represented by `\\n` (the user presses Enter).  `ReadString` may return an error but for the sake of simplicity we simply discard it with `_`.

**WARNING**: 
1. The code above won’t work from the playground. The program will run but it will not accept any user input. You will have to run it on your own computer.


2. It is not a good practice to simply ignore possible errors. However, in the context of this book, we can agree it is acceptable.

There are other ways to read user input; the following relies on `fmt`.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Print("What is your name? ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello " + name)
}
```

The first method is more flexible and robust, the second is more straightforward and lenient. For instance, try to read an `int` but type in some letters instead and find out what happens.

This advice applies even if the logic is trivial, as it it is often the case in this chapter where we deal with the basics. In the following example, we have separated the computation logic from the rest by using a function.

```go
package main

import (
	"fmt"
)

func computeSurface(width, height int) int {
	return width * height
}

func main() {
	fmt.Print("Enter width ")
	var width int
	fmt.Scan(&width)
	fmt.Print("Enter height ")
	var height int
	fmt.Scan(&height)
	surface := computeSurface(width, height)
	fmt.Printf("Surface: %d\n", surface)
}
```

## Random

Random numbers can be generated using package [fmt](https://golang.org/pkg/math/rand). It is a **pseudo-random number generator**, because numbers it generates are not actually purely random, they come from a predetermined sequence.

Here is an example that generates a “random” number. Run this program several times. Since the seed is always the same (`42`), it is to be expected that the generated pseudo-random numbers will always be the same.

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	fmt.Println(rand.Intn(10))
}
```

This is not very useful. In practice, you will need to set the seed to something different, that changes over time. The current system time would be a perfect candidate. Replace the seed with the following (import the `time` package):

```go
rand.Seed(time.Now().UnixNano())
```

This time (pun detected), you will get different results after several executions.

What if you want to get random numbers not from *0* to *n*, but from *m* to *n*? You simply generate a random number from *0* to *m-n* then add *m* to the result.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(random(10, 20))
}
```

This little routing will come in handy to solve some of the exercises to come.

## Read from a file

TBD

## Handling errors

TBD

## Practical case - Leap year

In the Gregorian calendar (used in most parts of the world), a leap year is a year containing one additional day added to keep the calendar year synchronized with the astronomical year.

It is also referred to as an intercalary year or a bissextile year. A leap day (February 29) is added to the calendar in leap years as a corrective measure, because the Earth does not orbit the sun in precisely 365 days. The [rule](https://aa.usno.navy.mil/faq/docs/calendars.php) to determine if a given year is a leap year is the following:

> Every year that is exactly divisible by four is a leap year, except for years that are exactly divisible by 100, but these centurial years are leap years if they are exactly divisible by 400. For example, the years 1700, 1800, and 1900 are not leap years, but the years 1600 and 2000 are.

Here is a solution:


* **if** (year is not divisible by 4) then (it is a common year)


* **else** **if** (year is not divisible by 100) then (it is a leap year)


* **else** **if** (year is not divisible by 400) then (it is a common year)


* **else** (it is a leap year)

Now think about a few relevant years, and convince yourself that this algorithm yields the correct output for each of your input, by evaluating successively all applicable conditions. What years would you pick for your test? Testing with `2021` for the “not divisible by 4” case is a fine example. Testing again with `2023` would not bring anything more as you would be testing the same case (“not divisible by 4”). A good test set covers all cases exactly once.

For instance, a function `isLeapYear(year: int) bool` could be tested with the following values:

| Branch to test

 | Example

 | Expected output

 |
| (year is not divisible by 4)

 | 2021

                                                                                | false

              |
| (year is not divisible by 100)

 | 2020

                                                                                | true

               |
| (year is not divisible by 400)

 | 2100

                                                                                | false

              |
| (it is a leap year) (end)

      | 2000

                                                                                | true

               |
The algorithm could be translated to the following Go function:

```go
func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%100 != 0 {
		return true
	} else if year%400 != 0 {
		return false
	} else {
		return true	
	}
}
```

Experiment with the test values proposed above, or with your own.

Rather than using conditional statements, we could rely on the concepts explained in the Logical operators to write a boolean expression. The rules to determine a leap year can be rephrased as follows:

> A year is a leap year if it is a multiple of 4 and not multiple of 100,
> or if it is a multiple of 400.

Notice the punctuation and translate it correctly using the logical operators `AND` and `OR`:

> For the year to be leap, it is either (a multiple of 4 and not multiple of 100) or it is (a multiple of 400).

That can be translated into Go:

```go
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
```

Extra parenthesis were used to highlight the scope of the logical operators.

Which version should you prefer? Probably the most readable one. Which one is the most readable? It is debatable. For such short code, a seasoned programmer may find the second version more compact and easy to read. Beyond a few conditions, it may become difficult or counterproductive to avoid the conditional if/else statements just for the sake of it. As in many things about programming, you will learn that there are often no clear-cut answers when it comes to design and style questions. In any case, try to remain consistant and to follow the language idioms when they apply.
