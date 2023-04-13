package main

import (
	"fmt"
)

func solveNQueens(n int) {
	board := make([]int, n)
	solve(board, 0, n)
}

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

func isSafe(board []int, column int, row int) bool {
	for i := column - 1; i >= 0; i-- {
		d := column - i
		if board[i] == row || board[i] == row-d || board[i] == row+d {
			return false
		}
	}
	return true
}

func main() {
	solveNQueens(4)
}
