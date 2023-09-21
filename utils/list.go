package utils

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val: nums[0],
	}

	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{
			Val: nums[i],
		}
		curr = curr.Next
	}

	return head
}
func EqualList(head *ListNode, target *ListNode) bool {
	for head != nil && target != nil {
		if head.Val != target.Val {
			return false
		}

		head = head.Next
		target = target.Next
	}

	return head == nil && target == nil
}
func PrintList(head *ListNode) string {
	result := "["

	for head != nil {
		result += " " + strconv.Itoa(head.Val)
		head = head.Next
	}

	return result + " ]"
}
