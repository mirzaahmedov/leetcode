// https://leetcode.com/problems/binary-search/

package main

import "fmt"

func main() {
	list := []int{2, 5}
	fmt.Println(search(list, 0))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
