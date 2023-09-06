package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(5))
}
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	rows := make([][]int, rowIndex+1)

	rows[0] = []int{1}

	for i := 1; i < len(rows); i++ {
		rows[i] = make([]int, i+1)

		for j := 0; j < len(rows[i]); j++ {
			if j-1 >= 0 {
				rows[i][j] += rows[i-1][j-1]
			}
			if j < len(rows[i-1]) {
				rows[i][j] += rows[i-1][j]
			}
		}
	}

	return rows[len(rows)-1]
}
