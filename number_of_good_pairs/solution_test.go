package solution

import (
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 1, 1, 3},
			want: 4,
		},
		{
			nums: []int{1, 1, 1, 1},
			want: 6,
		},
		{
			nums: []int{1, 2, 3},
			want: 0,
		},
	}

	for i, tc := range testCases {
		got := numIdenticalPairs(tc.nums)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
