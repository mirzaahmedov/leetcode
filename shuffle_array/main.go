// https://leetcode.com/problems/shuffle-the-array/

package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{1, 2, 3, 4}, 2))
}

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)

	for i := 0; i < n; i += 1 {
		res[2*i], res[2*i+1] = nums[i], nums[n+i]
	}

	return res
}
