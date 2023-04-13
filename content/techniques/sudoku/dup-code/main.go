package main

import "fmt"

type Board [][]int

func (b Board) check() bool {
	// Check rows
	for row := 0; row < 9; row++ {
		used := make(map[int]bool)
		for col := 0; col < 9; col++ {
			if b[row][col] == 0 {
				continue
			}
			if used[b[row][col]] {
				return false
			}
			used[b[row][col]] = true
		}
	}

	// Check columns
	for col := 0; col < 9; col++ {
		used := make(map[int]bool)
		for row := 0; row < 9; row++ {
			if b[row][col] == 0 {
				continue
			}
			if used[b[row][col]] {
				return false
			}
			used[b[row][col]] = true
		}
	}

	// Check 3x3 squares
	for squareRow := 0; squareRow < 3; squareRow++ {
		for squareCol := 0; squareCol < 3; squareCol++ {
			used := make(map[int]bool)
			for row := squareRow * 3; row < (squareRow+1)*3; row++ {
				for col := squareCol * 3; col < (squareCol+1)*3; col++ {
					if b[row][col] == 0 {
						continue
					}
					if used[b[row][col]] {
						return false
					}
					used[b[row][col]] = true
				}
			}
		}
	}

	return true
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
