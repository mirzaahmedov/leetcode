// https://leetcode.com/problems/search-a-2d-matrix/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  nums := [][]int{
    { 1, 3, 5, 7 },
    { 10, 11, 16, 20 },
    { 23, 30, 34, 50 },
  }

  fmt.Println(searchMatrix(nums, 3))
  fmt.Println(searchMatrix([][]int{ { 1 , 1 } }, 2))
}
func searchMatrix(matrix [][]int, target int) bool {
  row := len(matrix)
  col := len(matrix[0])
  
  left := 0
  right := row * col - 1
  mid := 0

  for left <= right {
    fmt.Println(left, right, mid)
    i := mid / col
    j := mid % col
    fmt.Println(i, j)

    if matrix[i][j] == target {
      return true
    }

    if target > matrix[i][j] {
      left = mid + 1
    } else {
      right = mid - 1
    }

    mid = (left + right) / 2
  }

  return false
}
