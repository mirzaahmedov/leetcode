package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestRotateRight(t *testing.T) {
	testCases := []struct {
		list *utils.ListNode
		k    int
		want *utils.ListNode
	}{
		{
			list: utils.GenerateLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
			k:    2,
			want: utils.GenerateLinkedListFromSlice([]int{4, 5, 1, 2, 3}),
		},
		{
			list: utils.GenerateLinkedListFromSlice([]int{0, 1, 2}),
			k:    4,
			want: utils.GenerateLinkedListFromSlice([]int{2, 0, 1}),
		},
	}

	for i, tc := range testCases {
		got := rotateRight(tc.list, tc.k)
		if !utils.EqualLinkedLists(got, tc.want) {
			t.Fatalf("#%d test - expected %v but got %v", i, utils.LinkedListToString(tc.want), utils.LinkedListToString(got))
		}
	}
}
