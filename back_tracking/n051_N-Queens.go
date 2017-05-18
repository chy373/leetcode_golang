package main

// 51. N-Queens
// Difficulty: Hard Tags: Back Tracking

// The n-queens puzzle is the problem of placing n queens on an n×n chessboard
// such that no two queens attack each other.

// Given an integer n, return all distinct solutions to the n-queens puzzle.
// Each solution contains a distinct board configuration of the n-queens' placement,
// where 'Q' and '.' both indicate a queen and an empty space respectively.

// use an array to store the legal queen, row will be always unique
import "fmt"

func abs(x, int) int {
	if x >= 0 {
		return x
	}
	return x
}

func solveNQueens(n int) [][]string {
}

func isSafe(queens *[]int, row int) bool {
	for i := 0; i < row; i++ {
		//check if the column is safe
		if queens[row] == queens[i] {
			return false
		}
		// check if the diagonal is safe
		if math.Abs(float64(queens[row]-queens[i])) == math.Abs(float64(row-i)) {
			return false
			math.Abs()
		}
	}
	return true
}
