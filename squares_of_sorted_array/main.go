// https://leetcode.com/problems/squares-of-a-sorted-array/?envType=study-plan&id=algorithm-i
package main

import (
  "fmt"
)

func main(){
  fmt.Println(sortedSquares([]int{ -4,-1,0,3,10 }))
  fmt.Println(sortedSquares([]int{ -7,-3,2,3,11 }))
}
func sortedSquares(nums []int) []int {
  var (
    i int = 0
    j, k int = len(nums) - 1, len(nums) - 1
  )

  var results = make([]int, len(nums))

  for i <= j {
    if nums[i]*nums[i] < nums[j]*nums[j] {
      results[k] = nums[j] * nums[j]
      j -= 1
    } else {
      results[k] = nums[i] * nums[i]
      i += 1
    }
    k -= 1
  }

  return results
}
