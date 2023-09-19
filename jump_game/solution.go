package solution

import (
	"fmt"
)

func canJump(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if i == len(nums)-1 && nums[0] == i {
			return true
		}

		if nums[len(nums)-i] == i {
			nums = nums[:len(nums)-i]
			i = 0
		}

		fmt.Println(nums)
	}

	return false
}
