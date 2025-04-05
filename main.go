package main

import "fmt"

func sudoku(board [][]int) {
	dfs(board, 0, 0)
}

func dfs(board [][]int, x, y int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	j := y
	for i := x; i < 9; i++ {
		for ; j < 9; j++ {
			if board[i][j] == 0 {
				for k := 1; k < 10; k++ {
					if isValidNumber(board, i, j, k) {
						board[i][j] = k
						if dfs(board, i, j+1) {
							return true
						}
						board[i][j] = 0
					}
				}

				return false
			}
		}
		j %= 9
	}
	return false
}

func isValidNumber(board [][]int, i, j, k int) bool {
	//if board[i][j] != 0 {
	//	return false
	//}

	if !isRowValid(board, i, k) {
		return false
	}

	if !isColValid(board, j, k) {
		return false
	}

	if !isGridValid(board, i, j, k) {
		return false
	}

	return true
}

func isRowValid(board [][]int, row int, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	return true
}

func isColValid(board [][]int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	return true
}

func isGridValid(board [][]int, x, y int, num int) bool {
	for i := x - x%3; i <= 3*(x/3)+2; i++ {
		for j := y - y%3; j <= 3*(y/3)+2; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}

func hasEmptyCell(board [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}

	return false
}

func main() {
	board := [][]int{
		{0, 0, 0, 5, 0, 0, 4, 2, 0},
		{0, 5, 0, 0, 0, 9, 6, 0, 0},
		{6, 8, 7, 0, 0, 0, 0, 1, 5},
		{0, 0, 9, 6, 5, 8, 1, 3, 2},
		{0, 0, 2, 0, 4, 0, 0, 0, 8},
		{0, 0, 0, 0, 9, 1, 0, 6, 4},
		{3, 0, 0, 0, 0, 2, 0, 0, 0},
		{7, 2, 0, 0, 1, 0, 3, 4, 9},
		{8, 9, 1, 0, 0, 7, 0, 5, 0},
	}

	sudoku(board)

	for _, row := range board {
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println()
	}
}
