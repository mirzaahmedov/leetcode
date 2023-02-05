// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1, 1, 1, 2, 4, 5, 5, 7}, 5))
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

func searchRange(nums []int, target int) []int {
	i := search(nums, target)
	j := i

	if i < 0 {
		return []int{-1, -1}
	}

	for i > 0 && nums[i-1] == target || j < len(nums)-1 && nums[j+1] == target {
		if i > 0 && nums[i-1] == target {
			i -= 1
		}
		if j < len(nums)-1 && nums[j+1] == target {
			j += 1
		}
	}

	return []int{i, j}
}
