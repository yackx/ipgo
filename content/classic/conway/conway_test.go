package conway

import (
	"reflect"
	"testing"
)

func TestNext(t *testing.T) {
	tests := []struct {
		name string
		grid *CellSet
		want *CellSet
	}{
		{"empty", NewCellSet(), NewCellSet()},
		{"single(dies)", CellSetFrom(Cell{0, 0}), NewCellSet()},
		{"block(still)", CellSetFrom(Cell{0, 0}, Cell{0, 1}, Cell{1, 0}, Cell{1, 1}), CellSetFrom(Cell{0, 0}, Cell{0, 1}, Cell{1, 0}, Cell{1, 1})},
		{"beehive(still)", CellSetFrom(Cell{2, 1}, Cell{3, 1}, Cell{1, 2}, Cell{4, 2}, Cell{2, 3}, Cell{3, 3}), CellSetFrom(Cell{2, 1}, Cell{3, 1}, Cell{1, 2}, Cell{4, 2}, Cell{2, 3}, Cell{3, 3})},
		{"tub(still)", CellSetFrom(Cell{1, 0}, Cell{0, 1}, Cell{2, 1}, Cell{1, 2}), CellSetFrom(Cell{1, 0}, Cell{0, 1}, Cell{2, 1}, Cell{1, 2})},
		{"blinker(2)", CellSetFrom(Cell{1, 0}, Cell{1, 1}, Cell{1, 2}), CellSetFrom(Cell{0, 1}, Cell{1, 1}, Cell{2, 1})},
	}

	for _, test := range tests {
		got := Next(test.grid)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s Next(%s)=%s, want=%s", test.name, test.grid, got, test.want)
		}
	}
}
func TestNeighbours(t *testing.T) {
	var tests = []struct {
		cell Cell
		want *CellSet
	}{
		{
			Cell{2, 4}, CellSetFrom(
			Cell{1, 3}, Cell{1, 4}, Cell{1, 5},
			Cell{2, 3}, Cell{2, 5},
			Cell{3, 3}, Cell{3, 4}, Cell{3, 5}),
		},
	}

	for _, c := range tests {
		got := neighbours(&c.cell)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("neighbours(%q) == %v, want %v", c.cell, got, c.want)
		}
	}
}
