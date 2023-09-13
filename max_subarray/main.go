package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
	fmt.Println(maxSubArray([]int{1}), 1)
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}), 23)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var (
		sum = 0
		min = 0
		max = -99999
	)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum-min > max {
			max = sum - min
		}

		if sum < min {
			min = sum
		}
	}

	return max
}
