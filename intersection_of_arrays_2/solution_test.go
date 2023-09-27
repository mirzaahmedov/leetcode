package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{9, 4},
		},
	}

	for i, tc := range testCases {
		got := intersect(tc.nums1, tc.nums2)
		if !utils.EqualSlices(got, tc.want) {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
