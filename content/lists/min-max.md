## Min and max

In order to find the greatest element of a slice, we first initialize `max` as the first element of the slice, as it is our sole candidate when we start.

```
12   34   6   13
^^
max == 12
```

Then, we iterate the slice. We compare each element we encounter with `max`. If the element is greater, it becomes the new `max`. Otherwise, we keep the current max… by doing nothing.

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

Noteworthy: Go 1.21 introduces `min` and `max` built-ins.

```go
fmt.Println(max(12, 34, 6, 13))
```

Unfortunately, they are not variadic.

```go
a := []int{12, 34, 6, 13}
fmt.Println(max(a))
// invalid argument: a (variable of type []int) cannot be ordered
```
