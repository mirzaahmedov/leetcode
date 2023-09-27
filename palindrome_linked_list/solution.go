package solution

import (
	"github.com/mirzaahmedov/leetcode/utils"
)

func isPalindrome(head *utils.ListNode) bool {
	length := 0
	curr := head
	values := []int{}

	for curr != nil {
		length += 1
		values = append(values, curr.Val)
		curr = curr.Next
	}

	i := length - 1
	for i >= 0 && head != nil {
		if head.Val != values[i] {
			return false
		}

		i -= 1
		head = head.Next
	}

	return true
}
