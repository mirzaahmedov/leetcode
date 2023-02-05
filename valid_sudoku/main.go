package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := [9]int{}
	cols := [9]int{}
	sqrs := [3][3]int{}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - 48

			if readBit(&rows[i], num) == 1 || readBit(&cols[j], num) == 1 || readBit(&sqrs[i/3][j/3], num) == 1 {
				return false
			} else {
				writeBit(&rows[i], num)
				writeBit(&cols[j], num)
				writeBit(&sqrs[i/3][j/3], num)
			}
		}
	}

	return true
}

func readBit(val *int, n byte) int {
	if (*val & (1 << n)) == (1 << n) {
		return 1
	} else {
		return 0
	}
}
func writeBit(val *int, n byte) {
	*val = *val | (1 << n)
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board))
	board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board))
}
