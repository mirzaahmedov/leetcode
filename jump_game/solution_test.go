package solution

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{
			input: []int{2, 3, 1, 1, 4},
			want:  true,
		},
		{
			input: []int{3, 2, 1, 0, 4},
			want:  false,
		},
	}

	for i, tc := range tests {
		got := canJump(tc.input)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
