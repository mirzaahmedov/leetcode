// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/?envType=study-plan&id=algorithm-i
package main

func main() {
}
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

func connect(root *Node) *Node {
  nodes := []*Node{root}
	
  for len(nodes) != 0 && nodes[0] != nil {
    children := make([]*Node, len(nodes) * 2)

    for i := 0; i < len(nodes); i++ {
      if i < len(nodes) - 1 {
        nodes[i].Next = nodes[i + 1]
      }

      if nodes[i].Left != nil && nodes[i].Right != nil {
        children[2 * i] = nodes[i].Left
        children[2 * i + 1] = nodes[i].Right
      }
    }

    nodes = children
  }
  return root
}
