// https://leetcode.com/problems/subsets/?envType=study-plan&id=algorithm-ii
package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
	// fmt.Println(subsets([]int{1, 2, 3, 4}))
}
func subsets(nums []int) [][]int {
	return findSubsets([]int{}, nums)
}
func findSubsets(curr, nums []int) [][]int {
	results := [][]int{curr}

	for i := range nums {
		s := append(curr, nums[i])

		for _, r := range findSubsets(s, nums[i+1:]) {
			fmt.Println(nums[i], r, results)
			results = append(results, r)
		}
	}

	return results
}
