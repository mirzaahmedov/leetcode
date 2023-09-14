package solution

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
	}{
		{
			input: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			want:  7,
		},
		{
			input: [][]int{{1, 2, 3}, {4, 5, 6}},
			want:  12,
		},
	}
	t.Run("Should return correct min path sum", func(t *testing.T) {
		for _, tc := range tests {
			got := minPathSum(tc.input)
			if got != tc.want {
				t.Fatalf("got %v but expected %v", got, tc.want)
			}
		}
	})
}
