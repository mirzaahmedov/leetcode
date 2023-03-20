// https://leetcode.com/problems/merge-two-binary-trees/?envType=study-plan&id=algorithm-i
package main

func main() {
}

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func merge(root1 *TreeNode, root2 *TreeNode) *TreeNode {
  if root1 == nil && root2 == nil {
    return nil
  }
  if root1 == nil {
    root1 = &TreeNode{
      Val: root2.Val,
      Left: root2.Left,
      Right: root2.Right,
    }
    return root1
  }
  if root2 == nil {
    return root1
  }
  root1.Val += root2.Val

  root1.Left = merge(root1.Left, root2.Left)
  root1.Right = merge(root1.Right, root2.Right)
  
  return root1
}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
  root1 = merge(root1, root2)

  return root1
}
