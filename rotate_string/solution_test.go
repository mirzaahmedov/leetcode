package solution

import (
	"testing"
)

func TestRotateString(t *testing.T) {
	testCases := []struct {
		s    string
		goal string
		want bool
	}{
		{
			s:    "abcde",
			goal: "cdeab",
			want: true,
		},
		{
			s:    "abcde",
			goal: "abced",
			want: false,
		},
		{
			s:    "gcmbf",
			goal: "fgcmb",
			want: true,
		},
	}

	for i, tc := range testCases {
		got := rotateString(tc.s, tc.goal)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
