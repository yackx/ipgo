## Sets

A set is an abstract data type that can store **unique** values, **without any particular order**.

Unlike many high level languages, Go does not have a data type for set. It has however another data type that can serve the same purpose: a map. Each key can appear at most once in a map, so the first property of a set (unique keys) can be met. There is no order either so the second property is met as well.

The trick is to ignore the values in the map.

The idiomatic way (or one of the idiomatic ways) to implement a set in Go is the following, assuming a set of strings:

```go
set := make(map[string]bool)
```

Suppose we want to store the folling values in a set: ${orange, apple, raspberry}$. Each fruit would be a key, with `true` as the associated value.

```go
fruits := map[string]bool{
    "apple":  true,
    "orange": true,
    "pear":   true,
}
```

All the concepts we have seen about map (add, delete, iterate) apply here.

The boolean value is a dummy value, serving no purpose other than to "make the map work". It is slightly inefficient (a bit of memory is used to hold a fake value) and arguably inelegant, so you may run into other implementations using a `struct{}` as the value.

```go
type void struct{}
var member void
fruits := map[string]void{
    "apple":  void,
    "orange": void,
    "pear":   void,
}
```

The first version is more straighforward if you need to declare a set "on the spot". In practice, and in this chapter, we will prefer a dedicated structure that **encapsulate** (hide the underlying details) the data type _set_ and its operations.

So instead of manipulating a map acting as a set, we will be manipulating a new data type Set. We will know that underneath, there is a map, but it won't appear in our interactions with the set _from the outside_.

First, let us create a set of strings.

```go
package myset

type StringSet struct {
    set map[string]bool
}
```

The `set` lower case is not accessible from outside the package `myset`. This is as desired, because want to deal with a `StringSet` rather than with the underlying map. We need in return to implement some basic functionalities, otherwise there would be no way to interact with the set. Let's start with a constructor.

```go
func NewStringSet() StringSet {
    return &StringSet{make(map[string]bool)}
}
```

`NewStringSet` builds a new `StringSet` structure by means of `StringSet{...}`. The unique parameter `make(map[string]bool)` creates the underlying map called `set` (because functionally it acts as a set, even if it is technically a map) in the `StringSet` structure.

If this does not make sense, consider a more verbose version:

```go
func NewStringSet() StringSet {
    // new structure, empty shell
    newStringSet := StringSet{}
    // actual set (yes, it's a map, don't tell anyone)
    actualSet := make(map[string]bool)
    // assign it to `set` inside the struct
    newStringSet.set = actualSet
    // return the new structure, set included
    return newStringSet
}
```

This second version is equivalent to the first one, but a trained programmer would prefer the first one. Try to figure out how they match.

We can create a `StringSet` with `set := NewStringSet()` but we must implement a few more operations before we can do something usefull with it.

| Operation  | Description
| ---------- | --------------------------------
| Add        | Add an element to the set
| Remove     | Remove an element from the set
| Contains   | Check if the set contains a value

We already know how to add or delete an element from a map.

```go
func (ss StringSet) Add(value string)  {
	ss.set[value] = true
}

func (ss StringSet) Remove(value string) {
    delete(ss.set, value)
}

func (ss StringSet) Contains(value string) bool {
	_, found := ss.set[value]
	return found
}
```

All these functions are self-explanatory. After all, our structure is merely a map in disguise. Notice the repeating `(ss StringSet)`. It means the function is applicable to a `StringSet` using a dot, like so:

```go
fruits := NewStringSet()
fruits.Add("blackberry")
fruits.Contains("orange")  // false
fruits.Remove("blackberry")
```

> Tip: in real life, you should consider using a package implementing a set, and avoid re-inventing the whell. This advise is applicable to any data type or structure.

Finally, let us add two usual operations.

| Operation    | Description
| ------------ | --------------------------------
| Intersection | Return a new set containing the elements common to this set and another set
| Union        | Return a new set containing the elements of this set and another set

```go
func (ss StringSet) Union(otherSet StringSet) StringSet {
	union := NewStringSet()
	for k, _ := range ss.set {
		union.set[k] = true
	}
	for k, _ := range otherSet.set {
		union.set[k] = true
	}
	return union
}
```

We can say that `Union` is a function on a `StringSet` that accepts one argument of type `StringSet` and returns a `StringSet`. There are 3 sets at play: the "target" `ss`, the `otherSet` and the returned set. The original sets `ss` and `otherSet` are left unchanged. In action:

```go
	vegetables := NewStringSet()
	vegetables.Add("eggplant")
	vegetables.Add("salad")
	union := fruits.Union(vegetables)
```

For `Intersection`, we will start from the target set. For each key in `ss`, we search for the same key in `otherSet`.

```go
func (ss StringSet) Intersection(otherSet StringSet) StringSet {
	intersection := NewStringSet()
	for k, _ := range ss.set {
        if _, found := otherSet.set[k]; found {
    		intersection.set[k] = true
        }
	}
	return intersection
}
```

We have used an idiomatic construct in the `if` statement: we assign and check `found` on the same line. It saves one line and express the intent that we want to check if the key is present and discard `found` afterwards.

To illustrate the function, I know of one fruit that can be considered as a vegetable.

```go
	fruits.Add("tomato")
	vegetables.Add("tomato")
	interestion := fruits.Intersection(vegetables)  // tomato
```

> Exercise: implement the following operations.

| Operation    | Description                              
| ------------ | ---------------------------------------------------------
| Pick         | Return an arbitrary element (any element) from this set
| Pop          | Return an arbitrary element (any element) from this set, deleting it from this set
| Count        | Count the number of elements in the set                     
| Difference   | Return the difference of this set and another set
| IsSubset     | Test whether this set is a subset of another set
