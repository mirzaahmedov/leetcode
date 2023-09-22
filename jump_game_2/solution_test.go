package solution

import (
	"testing"
)

func TestJump(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			nums: []int{1, 2},
			want: 1,
		},
		{
			nums: []int{1, 2, 3},
			want: 2,
		},
		{
			nums: []int{1, 1, 1, 1},
			want: 3,
		},
	}

	for i, tc := range testCases {
		got := jump(tc.nums)
		if got != tc.want {
			t.Fatalf("#%d test - expected %d but got %d", i+1, tc.want, got)
		}
	}
}
