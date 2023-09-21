package solution

import (
	"github.com/mirzaahmedov/leetcode/utils"
)

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil {
		return nil
	}

	length := 0
	curr := head

	for curr != nil {
		length += 1
		curr = curr.Next
	}

	k %= length

	count := length - k
	curr = head
	for curr != nil {
		count -= 1

		if count == 0 {
			newHead := curr.Next
			curr.Next = nil

			if newHead == nil {
				return head
			}

			node := newHead
			for node.Next != nil {
				node = node.Next
			}
			node.Next = head
			return newHead
		}

		curr = curr.Next
	}

	return head
}
