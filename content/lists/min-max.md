## Min and max

To find the greatest integer of a slice, we first initialize a variable `max` to the value of the first element of the slice. It is our sole candidate when we start, and the best one so far.

```
12   34   6   13
^^
max == 12
```

Then, we iterate the slice. We compare each element we encounter with `max`. If the element is greater, it becomes the new `max`. Otherwise, we keep the current `max`… by doing nothing.

```go
func maxInt(values []int) int {
	max := values[0]
	for _, i := range values {
		if i > max {
			max = i
		}
	}
	return max
}
```

At the end, the candidate is the actual maximum. Let’s call this function.

```go
func main() {
	a := []int{12, 34, 6, 13}
	fmt.Println(maxInt(a))
}
```

The output is of course `34`. The same logic can be applied to find the smallest element.

Make sure to handle the case where the slice is empty. In this case, we can `panic`.

```go
	if len(values) == 0 {
		panic("empty list")
	}
```

Go 1.18 introduces generics. We can make the function generic by adding a type parameter `T` with a constraint `constraints.Ordered`. This constraint ensures that the type `T` can be compared with the `>` operator.

```go
type Ordered interface {
	Integer | Float | ~string
}
```

Our function becomes:

```go
import "cmp"

func maxGeneric[T cmp.Ordered](values []T) T {
	max := values[0]
	for _, i := range values {
		if i > max {
			max = i
		}
	}
	return max
}
```

Lastly, Go 1.21 introduces `min` and `max` built-ins, turning our `maxGeneric` function into a pure exercise.

```go
fmt.Println(max(12, 34, 6, 13))
```

These built-ins do not work with arrays or slices however.

```go
a := []int{12, 34, 6, 13}
fmt.Println(slices.Max(a))
// ./prog.go:10:18: invalid argument: a (variable of type []int) cannot be ordered
```

To find the minimum or the maximum in an array or in a slice, use the `slices` package.

```go
import "slices"

a := []int{12, 34, 6, 13}
fmt.Println(slices.Max(a))
```
