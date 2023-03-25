// https://leetcode.com/problems/triangle/?envType=study-plan&id=algorithm-i
package main

import "fmt"

func main(){
  triangle := [][]int{
    {2},
    {3,4},
    {6,5,7},
    {4,1,8,3},
  }
  fmt.Println(minimumTotal(triangle))
  fmt.Println(triangle)
}
func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}
func minimumTotal(triangle [][]int) int {
  for i := 1; i < len(triangle); i++ {
    for j := 0; j < len(triangle[i]); j++ {
      if j > 0 && j < len(triangle[i]) - 1 {
        triangle[i][j] += min(triangle[i - 1][j - 1], triangle[i - 1][j])
      } else if j == 0 {
        triangle[i][j] += triangle[i - 1][j]
      } else {
        triangle[i][j] += triangle[i - 1][j - 1]
      }
    }
  }

  results := triangle[len(triangle) - 1]
  min := results[0]

  for i := 0; i < len(results); i++ {
    if results[i] < min {
      min = results[i]
    }
  }

  return min
}
