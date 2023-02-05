// https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/

package main

import "fmt"

func main() {
	values := []int{1, 15, 6, 3}

	fmt.Println(differenceOfSum(values))
}

func differenceOfSum(nums []int) int {
	elemSum := 0
	digitSum := 0

	for i := 0; i < len(nums); i++ {
		elemSum += nums[i]

		rem := nums[i]

		for rem > 0 {
			digitSum += rem % 10
			rem = rem / 10
		}

		digitSum += rem
	}

	diff := elemSum - digitSum

	if diff < 0 {
		return -diff
	}

	return diff
}
