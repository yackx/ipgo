## Concurrency (goroutines)

Concurrency is property of a program, when two or more tasks can start, run, and complete in overlapping time periods.

For instance, a web crawler that scans pages on the internet can be made of a scanner that downloads pages, and a parser that interprets them. They may run at the same time, but it is not mandatory to achieve concurrency. The parser could run out of work if the scanner is stuck because of a broken internet connection.

Go has outstanding built-in support for concurrency via **goroutines**.

A goroutine is a function that can run concurrently with other functions. It is a plain old function invoked with the `go` keyword.

```go
package main

import "fmt"

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func main() {
	go foo()
	fmt.Println("Terminated")
}
```

If you run this program, the result may not be what you expect. The program displays "Terminated" and... terminates. It does not print number from 0 to 99. This is because the goroutine had no chance to make any progress before the end of the program.

Let's modify this program so that waits half a second after printing the end message, but before it terminates.

```go
import (
	"fmt"
	"time"
)

func main() {
	go foo()
	fmt.Println("Terminated")
	time.Sleep(time.Millisecond * 500)
}
```

This time, the numbers are printed... after the message. We gave time to the goroutine to complete, but there is **no guarantee on the execution order**.

Besides, adding `time.Sleep` is a bad practice (how do you even decide on the "right" sleep duration? How much time does it need to display the number on different machines?). We will need another approach in the next section.
