package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(maxDepth(root))
}
func maxDepth(root *TreeNode) int {
	return depth(root)
}
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	a := depth(root.Left)
	b := depth(root.Right)

	if a > b {
		return a + 1
	}

	return b + 1
}
