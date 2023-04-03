// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(findMin([]int{ 3, 4, 5, 1, 2 }))
  fmt.Println(findMin([]int{ 4, 5, 6, 7, 0, 1, 2 }))
  fmt.Println(findMin([]int{ 11, 13, 15, 17 }))
  fmt.Println(findMin([]int{ 1 }))
}

func findMin(nums []int) int {
  left := 0
  right := len(nums) - 1
  
  for right >= left {
      mid := (left + right) / 2

      if mid != 0 && nums[mid - 1] > nums[mid] {
        return nums[mid]
      }
      if mid == 0 && nums[right] > nums[mid] {
        return nums[0]
      }

      if nums[mid] >= nums[0] {
        if nums[right] > nums[mid] {
          right = mid - 1
        } else {
          left = mid + 1
        }
      } else {
        right = mid - 1
      }

  }
  
  return nums[0]
}
