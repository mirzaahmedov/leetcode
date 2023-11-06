package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		list *utils.ListNode
		val  int
		want *utils.ListNode
	}{
		{
			list: utils.GenerateLinkedListFromSlice([]int{1, 2, 6, 3, 4, 5, 6}),
			val:  6,
			want: utils.GenerateLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
		},
		{
			list: utils.GenerateLinkedListFromSlice([]int{}),
			val:  1,
			want: utils.GenerateLinkedListFromSlice([]int{}),
		},
		{
			list: utils.GenerateLinkedListFromSlice([]int{7, 7, 7, 7}),
			val:  7,
			want: utils.GenerateLinkedListFromSlice([]int{}),
		},
	}

	for i, tc := range testCases {
		got := removeElements(tc.list, tc.val)
		if !utils.EqualLinkedLists(got, tc.want) {
			t.Fatalf("#%d test - expected %v but got %v", i, utils.LinkedListToString(tc.want), utils.LinkedListToString(got))
		}
	}
}
