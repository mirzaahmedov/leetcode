package main

import (
	"fmt"
)

func main() {
	fmt.Println("1,2,3,1", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println("1,2,3,4", containsDuplicate([]int{1, 2, 3, 4}))
}
func containsDuplicate(nums []int) bool {
	memo := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if memo[nums[i]] {
			return true
		}
		memo[nums[i]] = true
	}

	return false
}
