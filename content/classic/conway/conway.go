package conway

/*
A simple Conway's Game of Life in golang.
http://en.wikipedia.org/wiki/Conway's_Game_of_Life

No cycle detection. Infinite grid.

Copyright (C) 2021 Youri Ackx under GNU General Public License.
See the LICENSE file and [http://www.gnu.org/licenses/].
*/

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

// Next moves a grid to its next state
func Next(grid *CellSet) *CellSet {
	newGrid := NewCellSet()

	// All neighbours of all living cells. They are all candidate newborns
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