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
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for i, tc := range testCases {
		got := maxProfit(tc.prices)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
