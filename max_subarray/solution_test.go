package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want:  6,
		},
		{
			input: []int{1},
			want:  1,
		},
		{
			input: []int{5, 4, -1, 7, 8},
			want:  23,
		},
	}

	t.Run("Should return correct max sub array sum", func(t *testing.T) {
		for _, tc := range tests {
			got := maxSubArray(tc.input)

			if got != tc.want {
				t.Fatalf("expected %v but got %v", tc.want, got)
			}
		}
	})
}
