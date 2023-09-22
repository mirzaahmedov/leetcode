package solution

import (
	"testing"
)

func TestNumberOfEmployeesWhoMetTheTarget(t *testing.T) {
	testCases := []struct {
		hours  []int
		target int
		want   int
	}{
		{
			hours:  []int{0, 1, 2, 3, 4},
			target: 2,
			want:   3,
		},
		{
			hours:  []int{5, 1, 4, 2, 2},
			target: 6,
			want:   0,
		},
	}

	for i, tc := range testCases {
		got := numberOfEmployeesWhoMetTarget(tc.hours, tc.target)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
