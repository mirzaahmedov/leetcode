package solution

import "math"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			min := math.MaxInt

			if i > 0 {
				min = grid[i-1][j]
			}

			if j > 0 && grid[i][j-1] < min {
				min = grid[i][j-1]
			}

			if min != math.MaxInt {
				grid[i][j] += min
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
