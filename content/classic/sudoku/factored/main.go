package main

import "fmt"

type Board [][]int

func checkUsed(used map[int]bool, value int) bool {
	if value == 0 {
		return true
	}
	if used[value] {
		return false
	}
	used[value] = true
	return true
}

func (b Board) Check() bool {
	for row := 0; row < 9; row++ {
		used := make(map[int]bool)
		for col := 0; col < 9; col++ {
			if !checkUsed(used, b[row][col]) {
				return false
			}
		}
	}

	for col := 0; col < 9; col++ {
		used := make(map[int]bool)
		for row := 0; row < 9; row++ {
			if !checkUsed(used, b[row][col]) {
				return false
			}
		}
	}

	for square := 0; square < 9; square++ {
		used := make(map[int]bool)
		row := (square / 3) * 3
		col := (square % 3) * 3
		for i := row; i < row+3; i++ {
			for j := col; j < col+3; j++ {
				if !checkUsed(used, b[i][j]) {
					return false
				}
			}
		}
	}

	return true
}

func (b Board) Solve() Board {
	// Find an empty cell
	row, col := -1, -1
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				row, col = i, j
				break
			}
		}
		if row != -1 {
			break
		}
	}

	// If there are no empty cells, the board is already solved
	if row == -1 {
		return b
	}

	// Try each possible value for the empty cell, and recursively solve the board
	for num := 1; num <= 9; num++ {
		if b.checkValid(row, col, num) {
			b[row][col] = num
			if solved := b.Solve(); solved != nil {
				return solved
			}
			b[row][col] = 0 // Backtrack
		}
	}

	// If none of the values work, return nil to backtrack
	return nil
}

func main() {
	board := Board{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	if board.check() {
		fmt.Println("Board is valid!")
	} else {
		fmt.Println("Board is invalid.")
	}

	board2 := Board{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 6, 9}, // Invalid row with two 6's
	}
	if board2.check() {
		fmt.Println("Board is valid!")
	} else {
		fmt.Println("Board is invalid.")
	}

}
