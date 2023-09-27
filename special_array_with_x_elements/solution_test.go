package solution

import (
	"testing"
)

func TestSpecialArray(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 5},
			want: 2,
		},
		{
			nums: []int{0, 0},
			want: -1,
		},
		{
			nums: []int{0, 4, 3, 0, 4},
			want: 3,
		},
	}

	for i, tc := range testCases {
		got := specialArray(tc.nums)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
