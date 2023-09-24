## Eight Queens


**Place eight queens on a chessboard without any threat.**

### Statement

The eight queens puzzle requires to place 8 queens on a 8 by 8 chessboard so that no queen threatens another. In other words, no two queens can be found on the same row, column (called _file_ in chess) or diagonal.

The puzzle admits 92 solutions, of which 12 *fundamental* that differ other than by symetry, reflection or rotation.

The problem can be extended to $n$ queens on a $n \times n$ chessboard.

![One solution to the puzzle](content/classic/8queens/8queens.png)

### Data structure

An array of integers will suffice, with each index corresponding to a column and each value of the array representing a row. There is no need to declare a two-dimensional array.

The diagram above is therefore represented by `[]{6, 4, 2, 0, 5, 7, 1, 3}`.

### Backtracking

The puzzle serves as a good exercise to illustrate the **backtracking depth-first search** algorithm. Backtracking incrementally builds candidates to the solutions, and discards them (_backtracks_) as soon as the candidate cannot lead to valid solution. This approach is much more effecient than brute-force, where we would try to build all combinations on the board, even if they are doomed from the start, for instance by placing two queens on the first row and attempting to place a third one elsewhere, even though this can never lead to a valid position because the two first queens are in check.

To solve an $n$ queens board, we start by creating a slice of $n$ integers. Then we solve boards of increasing size $k$, starting at $k=0$ and finishing at $k=n$.

```go
func SolveNQueens(n int) {
	board := make([]int, n)
	solve(board, 0, n)
}
```

An auxiliary `solve()` functions does the actual work. When $k=n$, it means we have managed to put $n$ queens on the board, and we have found a valid solution that we can print. Otherwise, we first test if the board with the new queen is safe. In the affirmative, we try all possibles rows for column $k$.

```go
func solve(board []int, k, n int) {
	if k == n {
		fmt.Println(board)
	} else {
		for i := 0; i < n; i++ {
			if isSafe(board, k, i) {
				board[k] = i
				solve(board, k+1, n)
			}
		}
	}
}
```

Checking if a new queen is valid on the board is done in `isSafe()`.

```go
func isSafe(board []int, column int, row int) bool {
	for i := column - 1; i >= 0; i-- {
		d := column - i
		if board[i] == row || board[i] == row-d || board[i] == row+d {
			return false
		}
	}
	return true
}
```

Now we can solve a board of size `4` for instance:

```go
func main() {
	solveNQueens(4)
}
```

With the solutions:

```
[1 3 0 2]
[2 0 3 1]
```

### Caveats

- Our solver **prints** a solution when it founds one, rather than returning a list of all solutions. This can be considered as a side-effect of `solve()`. It limits the ability of the caller to print the solutions the way is they see fit, for instance in chess notation. It brings simplicity however, as a list of solutions does not have to be "carried around" by the solving function.

- For the sake of simplicity, we did not declare a custom type `Board` to abstract the low level `[]int`. Again, this is fine in the context of a simple exercise, and we may argue that introducing a custom type would be overengineering the solution.

### Other methods

- **Iterative** -- an iterative solver would require more code and would be less intuitive.
- **Channels** -- Go has a built-in mechansim called *channels*. It is used in the context of concurrent programming but it can act as an elegant communication mechanism between the solver and the caller, passing solutions from the former to the latter.
- **Representation** -- The solutions could be displayed in more natural chessboard coordinates.
