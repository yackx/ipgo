## Filtering

A common operation on lists is to filter values from a slice that match a certain criteria. For instance, we have a list of scores, ranging from 0 to 20. We want to keep all scores equal or above 12. Putting together what we already know about loops, slices, and conditional statements, we can write the following program:

```go
scores := []int{12, 14, 20, 3, 10, 16}
success := make([]int, 0)
for _, score := range scores {
	if score >= 12 {
		success = append(success, score)
	}
}
fmt.Println(success)
```

```
[12 14 20 16]
```

Back to the kitchen. Let’s say we want to skip every other fruits, and store the result in another slice called `skimmed`. To skip fruits in our iteration, instead of incrementing with `i++`, we will increase `i` by `2`. For the sake of simplicity, we will do something inefficient by declaring `skimmed` with a length of `0` and relying only on `append` to expand the slice.

```go
fruits := []string{"apple", "banana", "orange", "grapefruit"}
skimmed := make([]string, 0)
for i := 0; i < len(fruits); i = i + 2 {
	skimmed = append(skimmed, fruits[i])
}
fmt.Printf("%d fruits in %v\n", len(skimmed), skimmed)
```

A more efficient technique would be to declare `skimmed` with the proper length, that is, half of `fruits` length. On top of that, we would need a second index to remember where we stand in `skimmed`, and that index would be different than the index in `fruits`. We could also have leveraged the slice **capacity**, but we have left out that characterstic of slices in order to focus on algorithms. Hopefully our naive and simple approach does the trick.

```
2 fruits in [apple orange]
```
