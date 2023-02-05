// https://leetcode.com/problems/merge-k-sorted-lists/

package main

import (
	"fmt"
	"math"
)

// [[1,4,5],[1,3,4],[2,6]]

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

	lists := []*ListNode{list1, list2, list3}

	fmt.Println(mergeKLists(lists))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	mergedList := &ListNode{}
	mergedListHead := mergedList

	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	for {
		idx := -1
		min := math.Inf(1)

		for i := 0; i < len(lists); i++ {
			if lists[i] != nil && float64(lists[i].Val) < min {
				idx = i
				min = float64(lists[i].Val)
			}
		}

		if idx != -1 {
			mergedList.Next = lists[idx]
			mergedList = mergedList.Next
			lists[idx] = lists[idx].Next
		}

		allDone := true
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				allDone = false
				break
			}
		}

		if allDone {
			break
		}
	}

	return mergedListHead.Next
}
