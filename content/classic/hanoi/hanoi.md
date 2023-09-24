## Tower of Hanoi

**Learn recursion with this classical problem.**

### Presentation

The Tower of Hanoi[^hanoi] is a mathematical game or puzzle consisting of three _rods_ (or _sticks_) and a number of _disks_ (or _pegs_) of different diameters, which can slide onto any rod. The puzzle starts with the disks stacked on one rod in order of decreasing size, the smallest at the top.

[^hanoi]: https://en.wikipedia.org/wiki/Tower_of_Hanoi

For instance, let us label the 3 rods `A`, `B` and `C` and use 3 disks `{3, 2, 1}`, `3` being the largest:

```
A = {3, 2, 1}
B = {}
C = {}
```

![A model set of the Tower of Hanoi with 8 disks](content/classic/hanoi/Tower_of_Hanoi.jpeg)

### Rules

The objective is to move the entire stack to the last rod, obeying the following simple rules:

1. Only one disk may be moved at a time.
2. Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack or on an empty rod.
3. No disk may be placed on top of a disk that is smaller than it.

The minimal number of moves required to solve a Tower of Hanoi puzzle is $2^n − 1$, where $n$ is the number of disks.

### Recursive method

If we only display the moves and not the rods successive states, we do not need a data structure to represent the rods and the disks. This solution is straightforward, at the cost of testability.

To $solve(n, A, B, C)$ for $n$ disks using pegs $A$ (**source**), $B$ (**auxiliary**) and $C$ (**destination**) with $n>0$:

1. $solve(n-1, A, C, B)$
2. $move(A, C)$
3. $solve(n-1, B, A, C)$

In the chapter on recursion, we saw that we need a terminal condition. When $n=0$, there is nothing left to do.

```go
func solve(n int, source, auxiliary, destination string) {
	if n == 0 {
		return
	}
	solve(n-1, source, destination, auxiliary)
	fmt.Printf("Move from %s to %s\n", source, destination)
	solve(n-1, auxiliary, source, destination)
}

func main() {
	solve(4, "A", "B", "C")
}
```

And that's it! Problem solved, no sweat.

### Iterative method

The iterative approach is more involved than the recursive one. We declare a type `Tower`. Each rod is an `[]int` labelled by a string.

```go
type Tower struct {
	n int
	rods map[string][]int
}

func (tower *Tower) String() string {
	return fmt.Sprintf("%v", tower.rods)
}
```

Provide a new tower constructor based on the number of disks:

```go
func NewTower(n int) *Tower {
	var disks []int
	for i := n; i > 0; i-- {
		disks = append(disks, i)
	}
	tower := &Tower{n,make(map[string][]int)}
	tower.rods["A"] = disks  // descending [3, 2, 1]
	tower.rods["B"] = make([]int, 0)
	tower.rods["C"] = make([]int, 0)
	return tower
}
```

We will need to `swap` disks "smartly", depending on the rod's configuration. Swapping $x$ and $y$ implies to move either $x$ to $y$ or $y$ to $x$. Only one move is legal as one disk is always larger than the other.

```go
func (tower *Tower) swap(x, y string) {
	topDiskX := tower.topDisk(x)
	topDiskY := tower.topDisk(y)
	if topDiskX == 0 && topDiskY == 0 {
		return
	}
	if (topDiskX < topDiskY && topDiskX != 0) || topDiskY == 0 {
		tower.move(x, y)
	} else {
		tower.move(y, x)
	}
}
```

Return the top disk of the given rod:

```go
func (tower *Tower) topDisk(rod string) int {
	count := len(tower.rods[rod])
	if count == 0 {
		return 0
	}
	return tower.rods[rod][count-1]
}
```

We need to be able to move a disk from one rod to another:

```go
func (tower *Tower) move(from, to string) {
	fmt.Println(tower)
	fmt.Printf("Move from %s to %s\n", from, to)
	countFrom := len(tower.rods[from])
	topFrom := tower.rods[from][countFrom - 1]
	tower.rods[from] = tower.rods[from][:countFrom - 1]
	tower.rods[to] = append(tower.rods[to], topFrom)
}
```

We know we are done when the destination disk `C` contains all the disks, that is, its size is the size `n` of the tower.

```go
func (tower *Tower) isDone() bool {
	return len(tower.rods["C"]) == tower.n
}
```

At last, to put it all together, we have two sets of rules:

- Even number of disks
- Odd number of disks

For an even number of disks:

1. Make the legal move between rods $A$ and $C$ (in either direction).
2. Make the legal move between rods $A$ and $B$ (in either direction).
3. Make the legal move between rods $B$ and $C$ (in either direction).

For an odd number of disks:

1. make the legal move between rods $A$ and $C$ (in either direction).
2. make the legal move between rods $A$ and $B$ (in either direction).
3. make the legal move between rods $B$ and $C$ (in either direction).

Repeat the steps until complete.

The part "in either direction" explains why we designed a generic function `swap()` that is able to `move(x, y)` or `move(y, x)`.

We can now solve according to our algorithm.

```go
func (tower *Tower) solve() {
	var steps [][]string

	if tower.n % 2 == 0 {
		steps = [][]string{{"A", "B"}, {"A", "C"}, {"B", "C"}}
	} else {
		steps = [][]string{{"A", "C"}, {"A", "B"}, {"B", "C"}}
	}
	for {
		for _, step := range steps {
			tower.swap(step[0], step[1])
			if tower.isDone() {
				return
			}
		}
	}
}
```

This `main()` function will make our algorithm come to life.

```go
func main() {
	tower := NewTower(4)
	tower.solve()
	fmt.Println(tower)
}
```
