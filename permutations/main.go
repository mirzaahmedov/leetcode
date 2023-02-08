// https://leetcode.com/problems/permutations/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums[:1], nums[1+1:])

	fmt.Println(permute([]int{1, 2, 3}))

	fmt.Println("Hello World")
}
func permute(nums []int) [][]int {
	results := [][]int{}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	for i := 0; i < len(nums); i++ {
		rest := append([]int{}, nums[:i]...)
		rest = append(rest, nums[i+1:]...)

		result := permute(rest)

		for j := 0; j < len(result); j++ {
			result[j] = append(result[j], nums[i])
		}

		results = append(results, result...)
	}

	return results
}
