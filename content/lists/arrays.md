## Arrays

Suppose you want to manipulate several strings of characters. For instance, a shopping list, where each string is an item. So far we have used variables to contain numeric (integer or floating point) and string values. They were single values. You can declare several variables to hold several values, but this will quickly become cumbersome if the number of items on our list is not known in advance, or if there are a large amount of items to purchase.

All high-level languages offer some data structure to that effect. In many instances, the basic building block is called an **array**, defined as a *collection of elements* (values or variables), each identified by an *index* (plurals *indices*).

The following expression declares a variable `fruits` as an array of 3 strings:

```go
var fruits [3]string
```

There are no fruits in the array yet, or more precisely, each fruit is the empty string.

\begin{center}
\begin{tikzpicture}
  % Define the style for the array elements
  \tikzset{arr/.style={draw, minimum width=2cm, minimum height=0.75cm}}

  % Draw the array elements
  \node[arr] (fruit0) at (0,0) {};
  \node[arr] (fruit1) at (2,0) {};
  \node[arr] (fruit2) at (4,0) {};

  % Draw the indices
  \node at (0,-0.75) {0};
  \node at (2,-0.75) {1};
  \node at (4,-0.75) {2};

  % Add the array name
  \node [anchor=east] at (-1.5,0) {fruits};
\end{tikzpicture}
\end{center}

You can assign actual values to an index:

```go
fruits[0] = "apple"
fruits[2] = "banana"
fruits[1] = "orange"
```

Which can be represented as:

\begin{center}
\begin{tikzpicture}
  % Define the style for the array elements
  \tikzset{arr/.style={draw, minimum width=2cm, minimum height=0.75cm}}

  % Draw the array elements
  \node[arr] (fruit0) at (0,0) {apple};
  \node[arr] (fruit1) at (2,0) {orange};
  \node[arr] (fruit2) at (4,0) {banana};

  % Draw the indices
  \node at (0,-0.75) {0};
  \node at (2,-0.75) {1};
  \node at (4,-0.75) {2};

  % Add the array name
  \node [anchor=east] at (-1.5,0) {fruits};
\end{tikzpicture}
\end{center}

The first element is at position $0$. Using $0$ rather than $1$ as the index of the first element is a widely spread convention in computer programming. There is no obligation to fill all the slots, nor to fill them in any particular order, as shown in the example above. You can also declare the array with its elements:

```go
fruits := [3]string{"apple", "banana", "orange"}
```

The compiler can count the elements for you, eliminating the need to explicitly declare how many of them are present by using `...`, like so:

```go
fruits := [...]string{"apple", "banana", "organge"}
```

The array has a length, accessible with `len`

```go
fmt.Println(len(fruits))
```

```
3
```

An array cannot be resized. Not to worry, in Go there is a more potent data structure at our disposal called *slice*.
