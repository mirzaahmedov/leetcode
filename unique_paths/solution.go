package solution

func uniquePaths(m int, n int) int {
	if m < n {
		return calcPaths(m+n-2, m-1)
	}
	return calcPaths(m+n-2, n-1)
}
func calcPaths(n int, k int) int {
	r := 1
	for i := 1; i <= k; i++ {
		r *= (n - i + 1)
		r /= i
	}
	return r
}

// func uniquePaths(m int, n int) int {
// 	if m == 0 || n == 0 {
// 		return 0
// 	}
//
// 	grid := make([][]int, m)
//
// 	for i := 0; i < m; i++ {
// 		grid[i] = make([]int, n)
// 	}
//
// 	grid[0][0] = 1
//
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if i < m-1 {
// 				grid[i+1][j] += grid[i][j]
// 			}
// 			if j < n-1 {
// 				grid[i][j+1] += grid[i][j]
// 			}
// 		}
// 	}
//
// 	return grid[m-1][n-1]
// }
