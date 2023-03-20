// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/?envType=study-plan&id=algorithm-i
package main

import "fmt"

func main() {
  nums := []int{ 2,7,11,15 }

  fmt.Println(twoSum(nums, 9))
}

func twoSum(numbers []int, target int) []int {
  for i := 0; i < len(numbers); i++ {
    for j := i + 1; j < len(numbers); j++ {
      if numbers[i] + numbers[j] == target {
        return []int{ i + 1, j + 1 }
      }
    }
  }

  return []int{ -1, -1 }
}
