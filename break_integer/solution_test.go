package solution

import (
	"testing"
)

func TestIntegerBreak(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    2,
			want: 1,
		},
		{
			n:    10,
			want: 36,
		},
		{
			n:    5,
			want: 6,
		},
		{
			n:    6,
			want: 9,
		},
		{
			n:    3,
			want: 2,
		},
	}

	for i, tc := range testCases {
		got := integerBreak(tc.n)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
