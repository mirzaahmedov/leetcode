package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	common := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 40,
			Next: &ListNode{
				Val: 100,
			},
		},
	}

	headA := &ListNode{Val: 1, Next: &ListNode{Val: 25, Next: common}}
	headB := &ListNode{Val: 2, Next: common}

	fmt.Println(getIntersectionNode(headA, headB))

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	countA, countB := 0, 0

	for curr := headA; curr != nil; {
		countA += 1
		curr = curr.Next
	}
	for curr := headB; curr != nil; {
		countB += 1
		curr = curr.Next
	}

	if countA > countB {
		for i := 0; i < countA-countB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < countB-countA; i++ {
			headB = headB.Next
		}
	}

	currA, currB := headA, headB
	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA = currA.Next
		currB = currB.Next
	}

	return nil
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	memo := make(map[*ListNode]bool)
//
// 	for headA != nil {
// 		memo[headA] = true
// 		headA = headA.Next
// 	}
// 	for headB != nil {
// 		if memo[headB] {
// 			return headB
// 		}
//
// 		headB = headB.Next
// 	}
//
// 	return nil
// }

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	memo := make(map[*ListNode]bool)
//
// 	for headA != nil || headB != nil {
// 		if headA == headB {
// 			return headA
// 		}
// 		if memo[headA] {
// 			return headA
// 		}
// 		if memo[headB] {
// 			return headB
// 		}
//
// 		if headA != nil {
// 			memo[headA] = true
// 			headA = headA.Next
// 		}
// 		if headB != nil {
// 			memo[headB] = true
// 			headB = headB.Next
// 		}
// 	}
//
// 	return nil
// }
