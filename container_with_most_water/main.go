// https://leetcode.com/problems/container-with-most-water/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(maxArea([]int{ 1,8,6,2,5,4,8,3,7 }))
  fmt.Println(maxArea([]int{ 1,1 }))
}
// brute force solution
func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}
func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}
// func maxArea(height []int) int {
//   maxHeight := 0
//   res := 0
//   i, j := 0, len(height) - 1
//
//   for i := range height {
//     if height[i] > maxHeight {
//       maxHeight = height[i]
//     }
//   }
//
//   for i < len(height) {
//     j = len(height) - 1
//     for j > (res / maxHeight) {
//       res = max(res, (j - i) * min(height[i], height[j]))
//       j--
//     }
//     i++
//   }
//
//   return res
// }

func maxArea(height []int) int {
  res := 0
  l, r := 0, len(height) - 1

  for l < r {
    res = max(res, (r - l) * min(height[l], height[r]))

    if height[l] < height[r] {
      l++
    } else {
      r--
    }

  }

  return res
}
