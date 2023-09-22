package solution

import (
	"testing"
)

func TestFinalString(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "string",
			want: "rtsng",
		},
		{
			s:    "poiinter",
			want: "ponter",
		},
	}

	for i, tc := range testCases {
		got := finalString(tc.s)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
