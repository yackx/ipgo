## Variables and constants

A **variable** is a storage place in the computer memory used by a program. It has a name that identifies it. For instance, we could perform the following sequence of instructions in pseudo-code:

```
h <- 'hello'
a <- 1
b <- 2
score <- a + b + 4
```

Which would result in:

* The variable `h` contains the string of characters `'hello'`
* The variable `a` contains the value `1`
* The variable `b` contains the value `2`
* The variable `score` is the sum of `a+b+4`, that is `7`.

Or visually:

\usetikzlibrary{positioning}
\begin{center}
\begin{tikzpicture}[node distance=1cm]
\node (h_label) at (0,0) {$h$};
\node (h) [draw, rectangle, right of=h_label] {'hello'};
\node (a_label) [below=0.5cm of h_label] {$a$};
\node (a) [draw, rectangle, right of=a_label] {1};
\node (b_label) [right=2cm of h_label] {$b$};
\node (b) [draw, rectangle, right of=b_label] {2};
\node (score_label) [right=2cm of a_label] {$score$};
\node (score) [draw, rectangle, right of=score_label] {7};
\end{tikzpicture}
\end{center}


The equivalent in Go is to declare the variables, and to assign them a value.

```go
func main() {
    h := "hello"
    a := 1
    b := 2
    score := a + b + 4
}
```

Notice the `:=` operator to assign a value to a variable.

Assigning variables does not do much. If you were to enter this code snippet in the program main method, the program would execute but nothing would be displayed. Go ahead and try it in the [Go Playground](https://play.golang.org/).

Experiment with this small program and try to add `"hello"` and `2`. You will get an error message, as one can of course only perform additions on numbers.

Of course, we can print the results. We can modify the program so that it displays the sum of three numbers.

```go
fmt.Println(sum)
```

The modified program would unsurprisingly produce the following output:

```
7.5
```

As its name implies, a variable can *vary*. Or more precisely, the value it holds can vary.

```go
a := 7
a = a + 2
```

What is going on?


* After the first instruction, `a` holds the value `7`
* After the second instruction, `a` holds the value `a+2`, that is `7+2`, which evaluates to `9`.

As opposed to variables, we can define **constants** whose values cannot change once they are set.

```go
const a := 5
a = 6   // won't compile, already set
```
