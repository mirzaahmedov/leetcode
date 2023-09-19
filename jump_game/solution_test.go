package solution

import (
	"testing"
)

func TestJumpGame(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{2, 3, 1, 1, 4},
			want:  2,
		},
		{
			input: []int{2, 3, 0, 1, 4},
			want:  2,
		},
		{
			input: []int{1, 2},
			want:  1,
		},
	}

	for i, tc := range tests {
		got := jump(tc.input)
		if got != tc.want {
			t.Fatalf("#%d test - expected %d but got %d", i+1, tc.want, got)
		}
	}
}
