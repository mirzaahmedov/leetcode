// https://leetcode.com/problems/find-peak-element/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(findPeakElement([]int{ 1,2,3,1 }))
  fmt.Println(findPeakElement([]int{ 1,2,3,4,5 }))
  fmt.Println(findPeakElement([]int{ 6,5,4,3,2,1 }))
  fmt.Println(findPeakElement([]int{ 1,10,3,2,1,6,1 }))
  fmt.Println(findPeakElement([]int{ 1,2,1,3,5,6,4 }))
  fmt.Println(findPeakElement([]int{ 1,2,3,4,5,6,10,3,4 }))
  fmt.Println(findPeakElement([]int{ 1,2,2,1 }))
}
// func findPeakElement(nums []int) int {
//   left := 0
//   right := len(nums) - 1
//
//   for right - left > 2 {
//     mid := (left + right) / 2
//
//     if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
//       return mid
//     }
//
//     if nums[mid - 1] > nums[mid] {
//       right = mid
//     } else {
//       left = mid
//     }
//   }
//
//   if nums[left] > nums[right] {
//     return left
//   }
//
//   return right
// }
func findPeakElement(nums []int) int {
  left := 0
  right := len(nums) - 1

  for right > left {
    mid := (left + right) / 2

    if mid > 0 && mid < len(nums) - 1 && nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
      return mid
    }

    if nums[mid] < nums[mid + 1] {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }

  return left
}
