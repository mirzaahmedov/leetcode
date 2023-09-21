package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		list *utils.ListNode
		k    int
		want *utils.ListNode
	}{
		{
			list: utils.GenerateList([]int{1, 2, 3, 4, 5}),
			k:    2,
			want: utils.GenerateList([]int{4, 5, 1, 2, 3}),
		},
		{
			list: utils.GenerateList([]int{0, 1, 2}),
			k:    4,
			want: utils.GenerateList([]int{2, 0, 1}),
		},
	}

	for i, tc := range tests {
		got := rotateRight(tc.list, tc.k)
		if !utils.EqualList(got, tc.want) {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
