// https://leetcode.com/problems/symmetric-tree/
package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodes := []*TreeNode{root.Left, root.Right}

	for len(nodes) > 0 {
		p, q := nodes[0], nodes[len(nodes)-1]

		if p == nil && q == nil {
			nodes = nodes[1 : len(nodes)-1]
			continue
		}
		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}
		left := nodes[1 : len(nodes)/2]
		right := nodes[len(nodes)/2 : len(nodes)-1]
		nodes = append(left, p.Left, p.Right, q.Right, q.Left)
		nodes = append(nodes, right...)
	}

	return true
}
func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}

	root.Left.Left = &TreeNode{Val: 3}

	root.Right.Right = &TreeNode{Val: 3}

	fmt.Println(isSymmetric(root))
}
