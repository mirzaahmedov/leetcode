// https://leetcode.com/problems/symmetric-tree/
package main

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
	nodes := []*TreeNode{root, root}

	for len(nodes) > 0 {
		left, right := nodes[0], nodes[len(nodes)/2]

		if left == nil && right == nil {
			nodes = append(nodes[1:len(nodes)/2], nodes[len(nodes)/2+1:]...)
			continue
		}
		if left == nil || right == nil {
			return false
		}
	}

	return true
}
func main() {

}
