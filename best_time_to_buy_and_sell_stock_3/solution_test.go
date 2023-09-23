package solution

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   6,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{2, 1, 2, 0, 1},
			want:   2,
		},
		{
			prices: []int{6, 1, 3, 2, 4, 7},
			want:   7,
		},
	}

	for i, tc := range testCases {
		got := maxProfit(tc.prices)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
