package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		encoded []int
		first   int
		want    []int
	}{
		{
			encoded: []int{1, 2, 3},
			first:   1,
		},
	}
	for i, tc := range testCases {
		got := decode(tc.encoded, tc.first)
		if utils.EqualSlices(got, tc.want) {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
