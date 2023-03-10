// https://leetcode.com/problems/rotate-array/?envType=study-plan&id=algorithm-i
package main

import "fmt"

func main(){
  nums := []int{ 1, 2, 3, 4, 5, 6, 7 }

  rotate(nums, 2)

  fmt.Println(nums)
}
func rotate(nums []int, k int) {
  temp := make([]int, k)

  k %= len(nums)

  for i := len(nums) - 1; i >= 0; i-- {
      if len(nums) - i <= k {
        temp[k - (len(nums) - i)] = nums[i]
      } 
      
      if i < k {
        nums[i] = temp[i]
      } else {
        nums[i] = nums[i - k]
      }
  }
}
