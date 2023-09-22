package solution

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			nums: []int{1, 1, 1, 1},
			want: true,
		},
		{
			nums: []int{1, 1, 0, 1},
			want: false,
		},
		{
			nums: []int{0, 1},
			want: false,
		},
		{
			nums: []int{1, 2},
			want: true,
		},
		{
			nums: []int{1},
			want: true,
		},
		{
			nums: []int{0},
			want: true,
		},
	}

	for i, tc := range testCases {
		got := canJump(tc.nums)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
