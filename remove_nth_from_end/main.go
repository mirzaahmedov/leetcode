// https://leetcode.com/problems/remove-nth-node-from-end-of-list/?envType=study-plan&id=algorithm-i

package main

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  nodes := []*ListNode{}

  var current *ListNode
  for current != nil {
    nodes = append(nodes, current)
    current = current.Next
  }

  if len(nodes) == n {
    return head.Next
  }
  if len(nodes) > n {
    nodes[len(nodes) - n - 1].Next = nodes[len(nodes) - n].Next
  }
  
  return head
}
