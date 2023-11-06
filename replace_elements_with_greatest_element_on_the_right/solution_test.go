package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestReplaceElements(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{17, 18, 5, 4, 6, 1},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		{
			nums: []int{400},
			want: []int{-1},
		},
	}

	for i, tc := range testCases {
		got := replaceElements(tc.nums)
		if !utils.EqualSlices(tc.want, got) {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
