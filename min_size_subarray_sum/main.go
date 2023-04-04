// https://leetcode.com/problems/minimum-size-subarray-sum/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(minSubArrayLen(7,[]int{ 2,3,1,2,4,3 }))
  fmt.Println(minSubArrayLen(11,[]int{ 1,1,1,1,1,1,1,1 }))
}
func minSubArrayLen(target int, nums []int) int {
  res := len(nums)
  sum := 0

  for l, r := 0, 0; r < len(nums); r++ {
    sum += nums[r]

    if sum < target {
      continue
    }

    for sum >= target {
      sum -= nums[l]
      l++
    } 
    l--
    sum += nums[l]

    if r - l + 1 < res {
      res = r - l + 1
    }
  }

  if res == len(nums) && sum < target {
    return 0
  }

  return res
}
