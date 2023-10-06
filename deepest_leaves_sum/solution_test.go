package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestDeepestLeavesSum(t *testing.T) {
	testCases := []struct {
		root *utils.TreeNode
		want int
	}{
		{
			want: 15,
			root: &utils.TreeNode{
				Val: 1,
				Left: &utils.TreeNode{
					Val: 2,
					Left: &utils.TreeNode{
						Val: 4,
						Left: &utils.TreeNode{
							Val: 7,
						},
					},
					Right: &utils.TreeNode{
						Val: 5,
					},
				},
				Right: &utils.TreeNode{
					Val: 3,
					Right: &utils.TreeNode{
						Val: 6,
						Right: &utils.TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
		{
			want: 19,
			root: &utils.TreeNode{
				Val: 6,
				Left: &utils.TreeNode{
					Val: 7,
					Left: &utils.TreeNode{
						Val: 2,
						Left: &utils.TreeNode{
							Val: 9,
						},
					},
					Right: &utils.TreeNode{
						Val: 7,
						Left: &utils.TreeNode{
							Val: 1,
						},
						Right: &utils.TreeNode{
							Val: 4,
						},
					},
				},
				Right: &utils.TreeNode{
					Val: 8,
					Left: &utils.TreeNode{
						Val: 1,
					},
					Right: &utils.TreeNode{
						Val: 3,
						Right: &utils.TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		got := deepestLeavesSum(tc.root)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
