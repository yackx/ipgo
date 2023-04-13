## Conway's Game of Life

**Cellular automaton.**

### Presentation

The Game of Life[^game-of-life] is a cellular automaton devised by John Horton Conway[^conway] in 1970. It is a zero-player game, meaning that its evolution is determined by its initial state (*seed*), requiring no further input. One interacts with the game by creating an initial configuration and observing how it evolves. It is Turing complete[^turing-complete].

[^game-of-life]: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
[^conway]: https://en.wikipedia.org/wiki/John_Horton_Conway
[^turing-complete]: https://en.wikipedia.org/wiki/Turing_completeness

### Rules

The simulation evolves on an infinite, two-dimensional grid of cells, each of which is either live or dead. Every cell interacts with its eight neighbours. At each step in time, the following transitions occur:

- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
- Any live cell with two or three live neighbours lives on to the next generation.
- Any live cell with more than three live neighbours dies, as if by overpopulation.
- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

### Example of patterns

This simple set of rules can lead to remarkable patterns. Common patterns include:

- **Oscillators** -- return to their initial state after a finite number of generations called **periods** (_blinker_, _toad_, _beacon_, _pulsar_) ;
- **Still lifes** -- remain unchanged from one generation to the next (_block_, _beehive_, _loaf_, _boat_, _tub_) ;
- **Spaceships** -- travel accross the grid (_glider_ and other spaceships).

![Beehive, one form of still life](content/classic/conway/beehive.png)

Decades after Conway released his game, a self-replicating pattern called _Gemini_ was discovered. It creates a copy of itself while destroying its parent. Other astonishing patterns include _gilder guns_.

### Data structure

Some implementations use a two-dimensional array-like structures, for instace ``[][]bool`` to represent the grid. This is however inefficient and not practical. It cannot represent an infinite grid and can be a huge waste of memory on large grids. Instead, we represent a **cell** as pair of integers, and the **grid** as a set of cells. We shall use a `map[Cell]bool` to represent the grid.

```go
// Cell represents a cell, living or dead
type Cell struct {
	X, Y int
}

// CellSet represents a set of cells
// used as a grid or a set of neighbors
type CellSet struct {
	cells map[Cell]bool
}
```

Let's add:

- A `String()` method to display a cell in a human-readable format ;
- An empty grid constructor ;
- A cell-based grid constructor.

```go
func (c *Cell) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func NewCellSet() *CellSet {
	return &CellSet{make(map[Cell]bool)}
}

func CellSetFrom(cells ...Cell) *CellSet {
	cs := NewCellSet()
	for _, cell := range cells {
		cs.Add(cell)
	}
	return cs
}

func (cs *CellSet) String() string {
	s := "["
	for cell := range cs.cells {
		s = s + cell.String() + ", "
	}
	s = strings.TrimRight(s, ", ")
	s = s + "]"
	return s
}
```

### Implementation

Go does not offer a built-in type for set, so we need to bring our own helper functions like ``Add``, ``Contains`` or ``Intersect`` to the table.

```go
func (cs *CellSet) Add(cell Cell) {
	cs.cells[cell] = true
}

func (cs *CellSet) Contains(cell Cell) bool {
	_, found := cs.cells[cell]
	return found
}

func (cs *CellSet) Intersect(other *CellSet) *CellSet {
	intersect := NewCellSet()
	for _, cell := range cs.Cells() {
		for _, otherCell := range other.Cells() {
			if cell == otherCell {
				intersect.Add(cell)
			}
		}
	}
	return intersect
}

func (cs *CellSet) Cells() []Cell {
	var cells []Cell
	for cell := range cs.cells {
		cells = append(cells, Cell{cell.X, cell.Y})
	}
	return cells
}

func (cs *CellSet) Len() int {
	return len(cs.Cells())
}
```

With this setup, we can focus on the simulation itself. We will need a function that returns the neighbours of a cell.

```go
// Returns all the neighbours of the given cell, living or dead
func neighbours(cell *Cell) *CellSet {
	n := NewCellSet()
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				n.Add(Cell{i + cell.X, j + cell.Y})
			}
		}
	}
	return n
}
```

Lastly, `Next()` allows to transition from one state to the next. We first check if a cell survive to the next generation according to the rules. Then we examine all neighbor cells to determine if a new cell will appear (it must have 3 neighbors so the possible newborns always come next to living cell).

If the function returns an empty grid, we have reached a terminal state.

```go
// Next moves a grid to its next state
func Next(grid *CellSet) *CellSet {
	newGrid := NewCellSet()

	// All neighbours of all living cells.
    // They are all candidate newborns
	candidates := NewCellSet()

	// Survivors and current neighbours
	for _, cell := range grid.Cells() {
		neighbours := neighbours(&cell)
		countLivingNeighbours := grid.Intersect(neighbours).Len()
		if countLivingNeighbours == 2 || countLivingNeighbours == 3 {
			newGrid.Add(cell)
		}
		// Neighbors of this cell are newborn candidates
		for _, neighbour := range neighbours.Cells() {
			candidates.Add(neighbour)
		}
	}

	// Add eligible newborns
	for _, candidate := range candidates.Cells() {
		if !grid.Contains(candidate) {  // It's a empty cell
			neighbours := neighbours(&candidate)
			countLivingNeighbours := grid.Intersect(neighbours).Len()
			if countLivingNeighbours == 3 {
				newGrid.Add(candidate)
			}
		}
	}

	return newGrid
}
```

Create a `main()` method to run the simulation with different input, for instance

- Trival case: an empty grid that remains empty.
- A single cell that dies after the first generation.
- A beehive (still life) that remains unchanged, no matter how many times you call the next generation.
- A blinker of period 2.

For example:

```
NewCellSet()
CellSetFrom(Cell{0, 0})
CellSetFrom(
    Cell{2, 1}, Cell{3, 1}, Cell{1, 2},
    Cell{4, 2}, Cell{2, 3}, Cell{3, 3})
CellSetFrom(Cell{1, 0}, Cell{1, 1}, Cell{1, 2})
```

### Exercise

The `String()` grid and cell representations we provided are accurate, although a bit dull. As an exercise, display the automaton in a grid-like manner.
