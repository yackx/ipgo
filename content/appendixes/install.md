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

On GNU/Linux, a recent version may be available in the software repository, depending on the distribution.

### Edit

For this book, any decent editor will do. Visual Studio Code is recommanded on all platforms. `vim` or `emacs` are find choices. Jetbrains Goland is very convenient.

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

In some chapters, Go modules are used. See for instance \ref{stack} Stack.

In the following example, we have:

- a module file `go.mod`.
- a stack implementation in `package stack`.
- a demo in `package main`.

From a directory of your choosing, the directory and file structure is:

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
$ go mod init local.myname
```

Go will create the following file:

```bash
$ cat go.mod
module local.myname

go 1.21
```

Replace `local.myname` by your own namespace.
