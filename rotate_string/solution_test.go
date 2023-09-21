package solution

import (
	"testing"
)

func TestRotateString(t *testing.T) {
	tests := []struct {
		input string
		goal  string
		want  bool
	}{
		{
			input: "abcde",
			goal:  "cdeab",
			want:  true,
		},
		{
			input: "abcde",
			goal:  "abced",
			want:  false,
		},
		{
			input: "gcmbf",
			goal:  "fgcmb",
			want:  true,
		},
	}

	for i, tc := range tests {
		got := rotateString(tc.input, tc.goal)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
