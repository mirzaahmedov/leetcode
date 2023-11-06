package solution

import (
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 2, 1, 1, 3},
			want: true,
		},
		{
			nums: []int{1, 2},
			want: false,
		},
		{
			nums: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			want: true,
		},
	}

	for i, tc := range testCases {
		got := uniqueOccurrences(tc.nums)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
