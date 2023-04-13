package conway

import (
	"fmt"
	"strings"
)

// Cell represents a cell, living or dead
type Cell struct {
	X, Y int
}

func (c *Cell) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

// CellSet represents a set of cells
// used as a grid or a set of neighbors
type CellSet struct {
	cells map[Cell]bool
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

func (cs *CellSet) String() string {
	s := "["
	for cell := range cs.cells {
		s = s + cell.String() + ", "
	}
	s = strings.TrimRight(s, ", ")
	s = s + "]"
	return s
}