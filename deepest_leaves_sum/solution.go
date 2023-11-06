package solution

import (
	"github.com/mirzaahmedov/leetcode/utils"
)

// func deepestLeavesSum(root *utils.TreeNode) int {
// 	result := 0
// 	leaves := [][]int{}
//
// 	findLeaves(&leaves, root, []int{})
//
// 	for i := 0; i < len(leaves); i++ {
// 		result += leaves[i][len(leaves[i])-1]
// 	}
//
// 	return result
// }
// func findLeaves(leaves *[][]int, root *utils.TreeNode, result []int) {
// 	result = appendCopy(result, root.Val)
//
// 	if root.Left == nil && root.Right == nil {
// 		if len(*leaves) > 0 && len((*leaves)[0]) == len(result) {
// 			*leaves = append(*leaves, result)
// 		} else if len(*leaves) == 0 || (len(*leaves) > 0 && len(result) > len((*leaves)[0])) {
// 			*leaves = [][]int{result}
// 		}
// 		return
// 	}
//
// 	if root.Left != nil {
// 		findLeaves(leaves, root.Left, result)
// 	}
// 	if root.Right != nil {
// 		findLeaves(leaves, root.Right, result)
// 	}
// }
// func appendCopy(target []int, value int) []int {
// 	dest := make([]int, len(target)+1)
// 	copy(dest, target)
// 	dest[len(dest)-1] = value
// 	return dest
// }

func deepestLeavesSum(root *utils.TreeNode) int {
	level := deepestLevel(root)
	return calcDeepestLevelSum(level, root)
}
func calcDeepestLevelSum(level int, root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	if level == 1 {
		return root.Val
	}

	return calcDeepestLevelSum(
		level-1,
		root.Left,
	) + calcDeepestLevelSum(
		level-1,
		root.Right,
	)
}
func deepestLevel(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(
		deepestLevel(root.Left),
		deepestLevel(root.Right),
	) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
