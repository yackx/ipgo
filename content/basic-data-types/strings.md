## Strings

In Go, a *string* is a defined as a *slice of bytes*. For now, consider it as a (possibly empty) **sequence of bytes**. The package `strings` contains many utility functions to manipulate them, and the `fmt` package contains numerous functions to format them.

### Formatting

Package [fmt](https://golang.org/pkg/fmt/) offers formatting functions analogous to C’s `printf` and `scanf`. The format “verbs” are derived from C. Suppose you execute the following code:

```go
fmt.Println(1.0/3)
```

The output would be:

```
0.3333333333333333
```

It is convenient to be able to round to a certain number of digits, while keeping the original arbitrary precision of the variable. Let’s say we want to display the number with 2 digits. This can be achieved as follows:

```go
fmt.Printf("%.2f", 1.0/3)
```

Which would display:

```
0.33
```

Notice we’ve used `Printf` (print formatted) instead of `Println` (print line). There is something missing however. If we were to attempt to print the value of a second expression, it would be displayed on the same line as the first. With `Printf`, we need to explicitely add a *new line* at the end, with the `\n` symbol.

```go
fmt.Printf("%.2f\n", 1.0/3)
```

Go offers a large amount of formatting options as you can [read in the `fmt` documentation](https://golang.org/pkg/fmt/).

### Length and indices

Let us declare a variable to hold a message:

```go
message := "1, 2, 3, Go"
```

The number of bytes is called the length of the string and is never negative. It can be determined by using the built-in `len` function.

```go
fmt.Println(len(message))
```

Will output:

```
11
```

The string’s bytes can be access by **indices**. Like in most programming languages, Go starts counting from 0 rather than from 1. The first character is therefore located at index `0`. If the length of the string is `n`, the last byte index is `n-1`. Our `message` indices range from `0` to `10` included -- `10` being the 11th and last byte.

Individual bytes are accessed using square bracket notation. For instance, the “G” of “Go” is the 10th character, located at index `9`:

```go
g := message[9]
fmt.Println(g)
```

```
71
```

The result is somewhat unexpected. Why would the program output `71` rather than `G`? It turns out that 71 is the ASCII code for `G`. Check out the [ASCII table](http://www.asciitable.com/)[^strings-1]. If we want the character `G` to be displayed, we need to print formatted with `Printf` and the appropriate verb `c`, rather than using `Println`.

[^strings-1]: ASCII is an old is still widely used way to encore characters. See this table: http://www.asciitable.com/

```go
fmt.Printf("%c\n", g)
```

Attempts to access an index below `0` or beyond the last byte result in an error:

```
panic: runtime error: index out of range [100] with length 11
```

We can use a loop to iterate over the bytes in the string. The *for* loops starts at index *0*, and keeps running as long as the index *i* is less than the length of the string. At each iteration, we increment our index *i*.

```go
message := "Hello"
for i := 0; i < len(message); i++ {
	fmt.Printf("%c", message[i])
}
```

There is a different form of `for` loop that does not require to specify that we start at *0* and terminate at *n* while incrementing (adding 1) at each step:

```go
message := "Hello"
for i, c := range message {
	fmt.Printf("%d -> %c\n", i, c)
}
```

This construct with the help of `range` will assign the current index to `i` and the current character to `c`. We can then display the current index and the corresponding character.

```
0 -> H
1 -> e
2 -> l
3 -> l
4 -> o
```

This shorter form is to be preferred over the former whenever suitable.

### Package strings

The package “strings” (with an *s* at the end, not to be mixed with the type *string*) [offers a plethora](https://golang.org/pkg/strings/) of functions to manipulate functions[^strings-2]. A few of them will come very handy in this book.

[^strings-2]: https://golang.org/pkg/strings/

---------------------------------------------------------------------------------------------------
Function        Description
--------------- -----------------------------------------------------------------------------------
``Contains``    Reports whether a substring is within s

``Count``       Counts the number of times a substring appears in s

``Index``       Returns the index (position) of the first instance (occurrence) of a substring in s

``Repeat``      Returns a new string consisting of count copies of the string s

``Replace``     Replaces a substring in a given string

``Split``       Splits s into substrings separated by a given separator

``ToLower``     Returns s with all letters mapped to their lower case

``ToUpper``     Returns s with all letters mapped to their upper case
---------------------------------------------------------------------------------------------------

We have slightly simplified the original definitions for readability. Refer to the official documentation for a more accurate and formal description. Let’s see them in action.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
  	fmt.Println("Contains:  ", strings.Contains("Hello", "llo"))
  	fmt.Println("Count:     ", strings.Count("Hello", "l"))
  	fmt.Println("Index:     ", strings.Index("test", "e"))
  	fmt.Println("Repeat:    ", strings.Repeat("*", 5))
  	fmt.Println("Replace:   ", strings.Replace("Hello", "l", "r", -1))
  	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
  	fmt.Println("ToLower:   ", strings.ToLower("Hello"))
  	fmt.Println("ToUpper:   ", strings.ToUpper("World"))
}
```

```
Contains:   true
Count:      2
Index:      1
Repeat:     *****
Replace:    Herro
Split:      [a b c d e]
ToLower:    hello
ToUpper:    WORLD
```

Take a moment to browse the *strings* package documentation.

### Immutability

Strings are immutable: once created, it is impossible to change the contents of a string. Therefore, the following attempt to modify the string’s first character.

```go
func main() {
    message := "hello"
    message[0] = "H"
}
```

Will result in an error:

```
./prog.go:6:13: cannot assign to message[0]
```
