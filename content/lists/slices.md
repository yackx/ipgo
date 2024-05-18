## Slices

Arrays are a key construct but they offer limited flexibility. In Go, *slices* are usually preferred. They are built on top of arrays, leveraging their performance, while adding convenience.

A slice can be declared like an array, without counting the elements. Notice the subtle difference: we use `[]` for slices when we had `[...]` or `[3]` for arrays. So the following is a *slice*, not an *array*:

```go
fruits := []string{"apple", "banana", "orange"}
```

If you do not want to declare the elements yet, a slice can be created with the `make` function.

```go
fruits := make([]string, 3)
```

This declares a slice of 3 elements, each initialized to the empty string. Other than that, you can access and modify the slice elements just like we did with arrays. In addition, you can **add** an element at the end of a slice thanks to aptly-named the built-in `append` function.

```go
fruits = append(fruits, "cherry")
// len(fruits) == 4
// fruits[3] == "cherry"
```

Resulting in a slice of 4 elements, of which the three first remain empty.

\begin{center}
\begin{tikzpicture}
  % Define the style for the array elements
  \tikzset{arr/.style={draw, minimum width=2cm, minimum height=0.75cm}}

  % Draw the array elements
  \node[arr] (fruit0) at (0,0) {};
  \node[arr] (fruit1) at (2,0) {};
  \node[arr] (fruit2) at (4,0) {};
  \node[arr] (fruit3) at (6,0) {cherry};

  % Draw the indices
  \node at (0,-0.75) {0};
  \node at (2,-0.75) {1};
  \node at (4,-0.75) {2};
  \node at (6,-0.75) {3};

  % Add the array name
  \node [anchor=east] at (-1.5,0) {fruits};
\end{tikzpicture}
\end{center}

Notice that we assigned the result of the `append` expression back to `fruits`. The `append` statement alone does not modify the slice of fruits ; instead, it returns a modified slice. On top of that, the Go compliler rightfully prevents any attempt to execute a statement if it is not used. Therefore, any attempt to compile the following code would fail, because the result of the `append` call would be left unused:

```go
append(fruits, "cherry")

./prog.go:9:8: append(fruits, "cherry") evaluated but not used
```

However. Go allows a statement without effect if the result is explicitly discarded. To discard the result, assign it to `_` (underscore).

```go
_ = append(fruits, "cherry")
```

Of course, this last expression does probably not make sense in a program, but we are still free to shoot ourselves in the foot. The compiler will detect many errors made in “good faith”, but it does not protect us against all programming errors.

A slice can be appended to another as follows:

```go
a := []string{"apple", "pear"}
b := []string{"grapefruit", "orange"}
a = append(a, b...)
// a == []string{"apple", "pear", "grapefruit", "orange"}
```

\begin{center}
\begin{tikzpicture}
  % Define the style for the array elements
  \tikzset{arr/.style={draw, minimum width=2cm, minimum height=0.75cm}}

  % Draw the array elements
  \node[arr] (fruit0) at (0,0) {apple};
  \node[arr] (fruit1) at (2,0) {pear};
  \node[arr] (fruit2) at (4,0) {grapefruit};
  \node[arr] (fruit3) at (6,0) {orange};

  % Draw the indices
  \node at (0,-0.75) {0};
  \node at (2,-0.75) {1};
  \node at (4,-0.75) {2};
  \node at (6,-0.75) {3};

  % Add the array name
  \node [anchor=east] at (-1.5,0) {fruits};
\end{tikzpicture}
\end{center}

A slice can be created by “**slicing**” another slice (or array), that is to say, by taking a portion of it. Slicing is achieved by specifying a half-open range $[low,high)$ where $low$ is the index of the first element (included) and $high$ is the index of the last element (excluded). For example, `fruits[1:3]` will create a new slice including the fruits at indices from `1` included to `3` excluded --- that is, fruits at positions 1 and 2.

The newly created slice will have indices `0` and `1`, not the indices `1` and `2` it had in the original slice.

```go
fruits := []string{"apple", "banana", "orange", "grapefruit"}
preferred := fruits[1:2]
// preferred == []string{"banana", "orange"}
// preferred[0] == "banana"
// fruits[1] == "banana"
```

\begin{center}
\begin{tikzpicture}
  % Define the style for the array elements
  \tikzset{arr/.style={draw, minimum width=2cm, minimum height=0.75cm}}

  % Draw the array elements
  \node[arr] (fruit0) at (0,0) {apple};
  \node[arr] (fruit1) at (2,0) {banana};
  \node[arr] (fruit2) at (4,0) {orange};
  \node[arr] (fruit3) at (6,0) {grapefruit};

  % Draw the indices
  \node at (0,-0.75) {0};
  \node at (2,-0.75) {1};
  \node at (4,-0.75) {2};
  \node at (6,-0.75) {3};

  % Add the array name
  \node [anchor=east] at (-1.5,0) {fruits};
\end{tikzpicture}
\end{center}

\begin{center}
\begin{tikzpicture}
  % Define the style for the array elements
  \tikzset{arr/.style={draw, minimum width=2cm, minimum height=0.75cm}}

  % Draw the array elements
  \node[arr] (fruit0) at (0,0) {banana};
  \node[arr] (fruit1) at (2,0) {orange};

  % Draw the indices
  \node at (0,-0.75) {0};
  \node at (2,-0.75) {1};

  % Add the array name
  \node [anchor=east] at (-1.5,0) {preferred};
\end{tikzpicture}
\end{center}

Start and end indices can be omitted. In `s[low:high]`, the default value are `0` for `low`  and `len(s)` for `high`:

```go
// fruits[2:] == []string{"orange", "grapefruit"}
// fruits[:2] == []string{"apple", "banana"}
// fruits[:] == fruits
```

**WARNING**: Beware that slicing **does not copy data**. This makes the slicing operation efficient, but the outcome may not be what was expected.

Take the following example. Notice how, by modifying `letters`, we also modified `twoFirst`:

```go
letters := []byte{'a', 'b', 'c', 'd', 'e'}
twoFirst := letters[:2]
fmt.Printf("%c%c\n", twoFirst[0], twoFirst[1])  // ab
letters[0] = 'z'
fmt.Printf("%c%c\n", twoFirst[0], twoFirst[1])  // zb
```

To actually copy a slice, the built-in `copy`, as its name suggests, **copies** a slice from source to destination. It returns the number of elements copied. Watch out for the order of the arguments. The `destination` comes first!

```go
func copy(dst, src []T) int
```

The length of the destination should be equal or greater than the length of the source, otherwise only the smaller number of elements will be copied. Example:

```go
a := []string{"apple", "banana", "orange", "grapefruit"}
b := make([]string, len(a))
howMany := copy(b, a)
fmt.Println(b)
fmt.Println(howMany)
tooShort := make([]string, 2)
notEnough := copy(tooShort, a)
fmt.Println(tooShort)
fmt.Println(howMany)
```

Will produce:

```
[apple banana orange grapefruit]
4
[apple banana]
2
```

### Iterating

Remember the chapter on loops. Back then, we wrote a simple sum function.

```go
func sumTen() int {
  sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}
```

The same construct can be used to iterate over an array or a slice. At each iteration, instead of displaying the value of `i` (our index), we display the fruit at index `i`. We keep iterating while our index is strictly less than the number of fruits, so `i < len(fruits)`. The index will therefore vary from $0$ to $3$, and never reach $4$, otherwise we would attempt to access the fruit at index $4$, which would cause an error.

```go
fruits := []string{"apple", "banana", "orange", "grapefruit"}
for i := 0; i < len(fruits); i++ {
	fmt.Println(fruits[i])
}
```

This loop construct is generic but somewhat verbose, as we need to explicitly declare an index, declare the starting point at $0$, the end condition and the post-statement. Whenever possible, simpler forms of loops using `range` are preferred.

```go
fruits := []string{"apple", "banana", "orange", "grapefruit"}
for i := range fruits {
	fmt.Println(fruits[i])
}
```

Or the alternative form that allows you to retrieve both the index and the element at the same time:

```go
fruits := []string{"apple", "banana", "orange", "grapefruit"}
for i, fruit := range fruits {
	fmt.Printf("%d -> %s\n", i, fruit)
}
```

```
0 -> apple
1 -> banana
2 -> orange
3 -> grapefruit
```

What if you are not interested in the index `i`, but still want to retrieve the fruit? The index can be discarded with underscore `_`:

```go
fruits := []string{"apple", "banana", "orange", "grapefruit"}
for _, fruit := range fruits {
	fmt.Println(fruit)
}
```

The generic, more verbose construct has its place. If we need to skip some elements, or loop over them backward for instance, it can accomodate a wider range of scenarios. Let’s display the list of fruits in reverse order:

```go
fruits := []string{"apple", "banana", "orange", "grapefruit"}
for i := len(fruits) - 1; i >= 0; i-- {
	fmt.Println(fruits[i])
}
```

Notice that we had to start at `len(fruits) - 1` with $-1$ because the length is $4$, so the highest possible index is $3$.

```
grapefruit
orange
banana
apple
```

### Reference

[Slice intro on golang blog](https://blog.golang.org/slices-intro)[^slice-references]

[^slice-references]: https://blog.golang.org/slices-intro
