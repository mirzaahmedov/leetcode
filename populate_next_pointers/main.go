// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/?envType=study-plan&id=algorithm-ii
package main

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}
func main(){
}
func connect(root *Node) *Node {
  if root == nil {
    return root
  }

  var (
    node *Node
    queue = []*Node{root}
    temp []*Node
  )

  for len(queue) > 0 {
    temp = []*Node{}
    for i := 0; i < len(queue); i++ {
      node = queue[i]
      if i < len(queue) - 1 {
        node.Next = queue[i + 1]
      }
 
      if node.Left != nil {
        temp = append(temp, node.Left)
      }
      if node.Right != nil {
        temp = append(temp, node.Right)
      }
    }
    queue = temp
  }

  return root
}
