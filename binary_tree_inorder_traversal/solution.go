package solution

import (
	"github.com/mirzaahmedov/leetcode/utils"
)

func inorderTraversal(root *utils.TreeNode) []int {
	values := []int{}
	if root == nil {
		return values
	}

	traverse(&values, root)

	return values
}
func traverse(results *[]int, root *utils.TreeNode) {
	if root == nil {
		return
	}

	traverse(results, root.Left)
	*results = append(*results, root.Val)
	traverse(results, root.Right)
}
