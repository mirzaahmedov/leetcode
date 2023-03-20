// https://leetcode.com/problems/reverse-linked-list/?envType=study-plan&id=algorithm-i
package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func main(){
  head := &ListNode{Val: 1, Next: &ListNode{
    Val: 2, Next: &ListNode{
      Val: 3, Next: &ListNode{
        Val: 4, Next: nil,
      },
    },
  },
}

  newHead := reverseList(head)

  fmt.Println(newHead)
}
func reverseList(head *ListNode) *ListNode {
  var (
    prev *ListNode
    current *ListNode
    next *ListNode
  )

  current = head

  for current != nil {
    next = current.Next
    current.Next = prev
    prev = current

    current = next
  }
  
  return prev
}
