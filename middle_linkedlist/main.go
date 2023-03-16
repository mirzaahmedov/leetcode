// https://leetcode.com/problems/middle-of-the-linked-list/?envType=study-plan&id=algorithm-i

package main

// also fast and slow can be used as well
type ListNode struct {
    Val int
    Next *ListNode
}
func middleNode(head *ListNode) *ListNode {
  count := 0
  current := head
  for current.Next != nil {
    current = current.Next
    count++
  }

  count = count / 2

  for count > 0 {
    head = head.Next
    count--
  }

  return head
}
