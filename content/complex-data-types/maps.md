## Maps

### Definition

A **map** is a data type composed of a collection of **key-value pairs**. Each key can appear at most once.

With a map, we can for instance associate fruits with their energy value, or countries with their capital cities. Conceptually, it looks like this:

\begin{center}
\usetikzlibrary{positioning, arrows.meta}
\begin{tikzpicture}[node distance=1.5cm, auto, auto,->=stealth]
    \node[draw, rectangle] (A) {fruit};
    \node[draw, rectangle, right=of A] (B) {quantity};
    \draw[->, shorten >=0.1cm, shorten <=0.1cm] (A) -- (B);
\end{tikzpicture}
\end{center}

With some content:

\begin{center}
\usetikzlibrary{positioning, arrows.meta}
\begin{tikzpicture}[node distance=0.2cm and 1.5cm, auto,->=stealth]
    \node[draw, rectangle] (A) {apple};
    \node[draw, rectangle, right=of A] (B) {2};
    \draw[->, shorten >=0.1cm, shorten <=0.1cm] (A) -- (B);
    \node[draw, rectangle, below=of A] (C) {orange};
    \node[draw, rectangle, right=of C] (D) {3};
    \draw[->, shorten >=0.1cm, shorten <=0.1cm] (C) -- (D);
    \node[draw, rectangle, below=of C] (E) {pear};
    \node[draw, rectangle, right=of E] (F) {1};
    \draw[->, shorten >=0.1cm, shorten <=0.1cm] (E) -- (F);
\end{tikzpicture}
\end{center}

### Operations

A map offers operations to:

- add a new key-pair
- retrieve the value associated to a key
- check the existence of the key and therefore of the value
- delete a key-value pair
- iterate through its elements.

In Go, the corresponding type is aptly named `Map`.

```go
map[KeyType]ValueType
```

We can declare a variable `fruits` as a map in which keys are `string` (the name of the fruit) and the values are `int` (the quantity). It is up to the programmer to give meaning to the keys and values. Here, we intended to count fruits.

```go
var fruits map[string]int
```

This map is not initialized. It will behave like an empty map if you attempt to read from it, but any attempt to modify it will cause a runtime panic. Instead, we can declare the same map and initialize it with the built-in `make`.

```go
fruits = make(map[string]int)
// fruits == map[]
```

Now that the map is initialized, we can write to it.

```go
fruits["apple"] = 2
// fruits == map[apple: 2]
```

Alternatively, the map can be initialized with a _map literal_.

```go
fruits := map[string]int{
    "apple":  2,
    "orange": 3,
    "pear":   1,
}
// fruits == map[apple:2 orange:3 pear:1]
```

The _map literal_ can also be empty, using the same syntax, only with an empty literal `{}`. In this case, the empty map will be readable, just like an uninitialized map, but it will also be writable (add, modify or delete entries).

```go
fruits = map[string]int{}
```

A value can be retrieved and assigned to a variable. Suppose we set the apple count to `2`:

```go
appleCount = fruits["apple"]
// appleCount == 2
```

The number of elements (keys) in the map can be obtained with the built-in `len`

```go
typesOfFruits := len(fruits)
// typesOfFruits == 3
```

To add a new key-pair, assign the value to the corresponding key like so:

```go
fruits["banana"] = 5
```

There cannot be any duplicate key. If you assign a value an existing key, the existing value will be replaced by the new one. Duplicate values are not allowed.

```go
fruits["banana"] = 5
// map[apple:2 orange:6 pear:1, banana:5]
fruits["banana"] = 1
// "banana" value replaced, no new key created
// map[apple:2 orange:6 pear:1, banana:1]
```

To remove an entry from the map, use `delete`. It is safe to invoke even if the value is not present in the map.

```go
// fruits == {"apple": 2, "orange": 3, "pear": 1}
delete(fruits, "apple")
// fruits == {"orange": 3, "pear": 1}
delete(fruits, "foo")
// fruits == {"orange": 3, "pear": 1}
```

In a typical Go way, we can perform a two-values assignement that tests for the presence of the key, and retrieve it in one operation. Starting from our original map, `orangeCount` will contain `3` and `present` will be `true` after this assignment.

```go
orangeCount, present = fruits["orange"]
// orangeCount == 3
// present == true
```

Using a similar construct, we can test a key for existence by discarding the value with `_` (underscore).

```go
_, present = fruits["orange"]
// present == true
```

### Iteration

Iteration is achieved with the well-known `range`. Each iteration returns a key and a value, so in our case, a fruit and its count.

```go
for fruit, count := range fruits {
    fmt.Println("Fruit:", fruit, "Count:", count)
}
```

```
Fruit: apple Count: 2
Fruit: orange Count: 3
Fruit: pear Count: 1
```

> _Note_: The map is not ordered.

### Exercices

1. Declare an uninitialized map holding name-age pairs. Attempt to read the key `john`, which is not present. What happens? Attempt to write the pair `steve`: `32`. Observe the behavior.

2. Map countries to their capitals, using a map litteral with two entries: `Belgium`: `Brussels` and `Spain`: `Madrid`. Read the values of the following countries: `Spain`, `Italy`.

3. In the same map, add the entry `Chile`: `Santiago`.

4. Iterate the map to display all countries and their capitals, one per line, in the following format: "The capital of Belgium is Brussels".

5. Delete a country of your choice.

6. Display how many countries you have in your map (should be 2).
