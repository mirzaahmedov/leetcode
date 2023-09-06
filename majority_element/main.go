package main

import "fmt"

func main() {
	fmt.Println("3 == ", majorityElement([]int{3, 2, 3}))
	fmt.Println("2 == ", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
func majorityElement(nums []int) int {
	counter := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		counter[nums[i]] += 1
		if counter[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}

	return -1
}
