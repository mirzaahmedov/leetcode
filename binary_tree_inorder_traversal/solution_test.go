package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		tree *utils.TreeNode
		want []int
	}{
		{
			tree: &utils.TreeNode{
				Val: 1,
				Right: &utils.TreeNode{
					Val: 2,
					Left: &utils.TreeNode{
						Val: 3,
					}}},
			want: []int{1, 3, 2},
		},
		{
			tree: nil,
			want: []int{},
		},
		{
			tree: &utils.TreeNode{Val: 1},
			want: []int{1},
		},
	}

	for i, tc := range testCases {
		got := inorderTraversal(tc.tree)
		if !utils.EqualSlices(got, tc.want) {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
