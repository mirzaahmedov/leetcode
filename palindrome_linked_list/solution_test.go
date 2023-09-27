package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		list *utils.ListNode
		want bool
	}{
		{
			list: utils.GenerateLinkedListFromSlice([]int{1, 2, 2, 1}),
			want: true,
		},
		{
			list: utils.GenerateLinkedListFromSlice([]int{1, 2}),
			want: false,
		},
	}

	for i, tc := range testCases {
		got := isPalindrome(tc.list)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
