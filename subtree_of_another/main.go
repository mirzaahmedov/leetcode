// https://leetcode.com/problems/subtree-of-another-tree/?envType=study-plan&id=algorithm-ii
package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func main(){

}
func compare(root *TreeNode, subRoot *TreeNode) bool {
  if root == nil && subRoot == nil {
    return true
  }
  if root == nil || subRoot == nil {
    return false
  }
  if root.Val != subRoot.Val {
    return false
  }
  return compare(root.Left, subRoot.Left) && compare(root.Right, subRoot.Right)
}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
  if root == nil {
    return root == subRoot
  }

  var (
    node *TreeNode
    queue []*TreeNode = []*TreeNode{root}
  )

  for len(queue) > 0 {
    node, queue = queue[0], queue[1:]

    if compare(node, subRoot) {
      return true
    }

    if node.Left != nil {
      queue = append(queue, node.Left)
    }
    if node.Right != nil {
      queue = append(queue, node.Right)
    }
  }

  return false
}
