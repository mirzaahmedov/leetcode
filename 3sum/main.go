// https://leetcode.com/problems/3sum/?envType=study-plan&id=algorithm-ii
package main

import "sort"

func main(){
}
func threeSum(nums []int) [][]int {
  res := [][]int{}
  prev := -1

  l, r := 0, 0
  
  sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
  })

  for i := 0; i < len(nums); i++ {
    if prev == -1 && nums[prev] != nums[i] {
      l, r = i + 1, len(nums) - 1

      for r > l {

      }
    }

  }

  return res
}

