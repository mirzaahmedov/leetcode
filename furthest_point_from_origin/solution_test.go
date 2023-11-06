package solution

import (
	"testing"
)

func TestFurthestDistanceFromOrigin(t *testing.T) {
	testCases := []struct {
		moves string
		want  int
	}{
		{
			moves: "L_RL__R",
			want:  3,
		},
		{
			moves: "_R__LL_",
			want:  5,
		},
		{
			moves: "_______",
			want:  7,
		},
	}

	for i, tc := range testCases {
		got := furthestDistanceFromOrigin(tc.moves)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
