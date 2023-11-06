package solution

import (
	"testing"
)

func TestDecodeMessage(t *testing.T) {
	testCases := []struct {
		key     string
		message string
		want    string
	}{
		{
			key:     "the quick brown fox jumps over the lazy dog",
			message: "vkbs bs t suepuv",
			want:    "this is a secret",
		},
		{
			key:     "eljuxhpwnyrdgtqkviszcfmabo",
			message: "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
			want:    "the five boxing wizards jump quickly",
		},
	}

	for i, tc := range testCases {
		got := decodeMessage(tc.key, tc.message)
		if got != tc.want {
			t.Fatalf("#%v expected %v; got %v", i, tc.want, got)
		}
	}
}
