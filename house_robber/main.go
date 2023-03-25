// https://leetcode.com/problems/house-robber/?envType=study-plan&id=algorithm-i
package main

import (
  "fmt"
)

func main(){
  fmt.Println(rob([]int{ 1, 2, 3, 1 }))
}
func max(a, b int) int {
  if a > b {
    return a
  }

  return b
}
func rob(nums []int) int {
  var (
    curr_included = 0
    max_curr_excluded = 0
  )

  for i := 0; i < len(nums); i++ {
    curr_included, max_curr_excluded = max_curr_excluded + nums[i], max(curr_included, max_curr_excluded)
  }

  return max(curr_included, max_curr_excluded)
}

