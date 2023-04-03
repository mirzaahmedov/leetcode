// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/?envType=study-plan&id=algorithm-ii
package main

type ListNode struct {
    Val int
    Next *ListNode
}

func main(){
}
func deleteDuplicates(head *ListNode) *ListNode {
  var (
    current *ListNode = head
    prev *ListNode = nil
  )

  for current != nil {
    if current.Next != nil && current.Val == current.Next.Val {
      for current.Next != nil && current.Val == current.Next.Val {
        current.Next = current.Next.Next
      } 

      current = current.Next
      if prev != nil {
        prev.Next = current
      } else {
        head = current
      }
    } else {
      prev = current
      current = current.Next
    }
  }

  return head
}
