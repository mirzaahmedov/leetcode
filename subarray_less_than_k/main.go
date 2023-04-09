// https://leetcode.com/problems/subarray-product-less-than-k/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(numSubarrayProductLessThanK([]int{ 10, 5, 2, 6 }, 100))
  fmt.Println(numSubarrayProductLessThanK([]int{ 1, 2, 3 }, 0))
}
func numSubarrayProductLessThanK(nums []int, k int) int {
  result := 0
  product := 1

  for l, r := 0, 0; r < len(nums); r++ {
    product *= nums[r]

    for l <= r && product >= k {
      product /= nums[l]
      l++ 
    }

    result += r - l + 1
  }

  return result
}
