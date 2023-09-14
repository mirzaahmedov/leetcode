package solution

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		want int
	}{
		{
			m:    3,
			n:    7,
			want: 28,
		},
		{
			m:    3,
			n:    2,
			want: 3,
		},
	}
	t.Run("Should return number of unique paths", func(t *testing.T) {
		for _, tc := range tests {
			got := uniquePaths(tc.m, tc.n)

			if got != tc.want {
				t.Fatalf("got %v but expected %v", got, tc.want)
			}
		}
	})
}
