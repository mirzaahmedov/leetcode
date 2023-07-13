// https://leetcode.com/problems/same-tree/
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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	nodes := []*TreeNode{p, q}

	for len(nodes) > 0 {
		p = nodes[0]
		q = nodes[1]

		if p == nil && q == nil {
			nodes = nodes[2:]
			continue
		}
		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		nodes = append(nodes[2:], p.Left, q.Left, p.Right, q.Right)
	}

	return true
}
func main() {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	fmt.Println(isSameTree(p, q), "should be true")

	p = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q = &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	fmt.Println(isSameTree(p, q), "should be false")

	p = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	q = &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}

	fmt.Println(isSameTree(p, q), "should be false")
}
