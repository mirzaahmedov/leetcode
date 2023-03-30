// https://leetcode.com/problems/search-in-rotated-sorted-array/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  var nums = []int{4,5,6,7,0,1,2}
  var target = 0
  fmt.Println(search(nums, target))
}
func search(nums []int, target int) int {
  var (
    left = 0
    right = len(nums) - 1
    mid = 0
  )

  for left <= right {
    mid = (left + right) / 2

    if nums[mid] == target {
      return mid
    }

    if nums[mid] >= nums[left] {
      if nums[mid] > target && nums[left] <= target {
        right = mid - 1
      } else {
        left = mid + 1
      }
    } else {
      if nums[mid] < target && nums[right] >= target {
        left = mid + 1
      } else {
        right = mid - 1
      }
    }
  }

  return -1
}
