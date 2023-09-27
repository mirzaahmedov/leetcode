// https://leetcode.com/problems/remove-linked-list-elements/
package solution

import (
	"github.com/mirzaahmedov/leetcode/utils"
)

func removeElements(head *utils.ListNode, val int) *utils.ListNode {
	var (
		prev *utils.ListNode
		curr *utils.ListNode
	)

	for head != nil && head.Val == val {
		head = head.Next
	}

	curr = head

	for curr != nil {
		if curr.Val == val {
			if prev != nil {
				prev.Next = curr.Next
			}
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	return head
}
