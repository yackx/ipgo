## First program

To get acquainted with a programming language, it is customary to write a program that displays a [friendly message](https://en.wikipedia.org/wiki/%22Hello,_World!%22_program/)[^hello-world]. For the first chapters, you do not necessarily need to have Go installed on your computer. You can simply enter an execute your program in the [Go Playground](https://go.dev/play/)[^go-playground]:

[^hello-world]: "Hello, World!". See https://en.wikipedia.org/wiki/%22Hello,_World!%22_program/

[^go-playground]: See https://go.dev/play/

```go
package main

import "fmt"

func main() {
        fmt.Println("Hello, world!")
}
```

This simple program already raises several questions. What is a `package`, why is there an `import` with `"fmt"` in it? One thing at the time. For now, remember that `Println` prints a line on the screen. The term “print” and its variations are found in many programming languages, more often than “display” and comes from the ancient times where the primary interface with the computer was a printer rather than a screen. `main` is the program entry point. It is where the first instruction will be executed.

This program is found by default when you open the playground, so you don’t even have to type it. Mark this page, as this program will be the skeleton for many exercises we will solve in the next chapters.
