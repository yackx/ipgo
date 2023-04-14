\newpage
## Install and run

### Go playground

For the impatient, or when you are working from a computer where Go has not been installed, the Go Playground allows you to run code in the browser.

https://play.golang.org/

### Install

Go to the installation page and follow the instructions.

https://go.dev/doc/install

On macOS, consider using Homebrew.

https://formulae.brew.sh/formula/go

### Edit

For this book, any decent editor will do.

Visual Studio Code is recommanded on all platforms.

https://code.visualstudio.com/

### Run

```go
// cat main.go

package main

import "fmt"

func main() {
        fmt.Println("Hello, world!")
}
```

Alternatively, you can open the file from your file explorer or navigator as well.

Run with `go run`:

```bash
$ go run main.go
Hello, world!
```
### Modules

In some chapters, we make the code more modular by using modules.

A good example is found in the ch. \ref{stack} Stack. You don't need to understand the details of the code if you just began reading this book.

In this example, we have:

- a module file `go.mod`
- a stack implementation in `package stack`
- a demo in `package main`.

From a directory of your choosing, the file structure is:

```
+-- ipgo
    +-- exercise
        +-- stack
        |   +-- stack.go        
        +-- main.go
        +-- go.mod
```

First, create a module:

```bash
$ go mod init be.sugoi.ipgo
```

Go will create the following file:

```bash
$ cat go.mod
module be.sugoi.ipgo

go 1.20
```

I chose `be.sugoi.ipgo` because it matches this book's namespace, but you can choose whatever you like.

The stack:

```go
// stack.go

package stack

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

// rest of the code omitted
```

The demo:

```go
// main.go

package main

import (
	"fmt"
	"be.sugoi.ipgo/stack"
)

func main() {
	s := stack.New[int]()
	fmt.Println(s.Len()) // 0
}
```
