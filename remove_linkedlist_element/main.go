// https://leetcode.com/problems/remove-linked-list-elements/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(removeElements(&ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}, 6))
	fmt.Println(removeElements(&ListNode{1, &ListNode{1, nil}}, 1))
	fmt.Println(removeElements(&ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}}, 7))
	fmt.Print(removeElements(&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}, 2))
}
func removeElements(head *ListNode, val int) *ListNode {
	for head == nil {
		return head
	}
	for head.Val == val {
		if head.Next != nil {
			head = head.Next
		} else {
			return head.Next
		}
	}

	prev := head
	curr := head.Next

	if curr == nil {
		return head
	}

	for curr.Next != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
	}

	if curr.Val == val {
		prev.Next = curr.Next
	}

	return head
}
