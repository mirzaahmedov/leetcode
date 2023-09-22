package solution

import (
	"testing"
)

// ))
// 	"/../"))
// 	"/home/foo"))
// 	"/..."))
// 	"/.../"))
// 	"/..hidden"))
// 	"/.....hidden"))
// 	"/hello../world"))

func TestSimplifyPath(t *testing.T) {
	testCases := []struct {
		path string
		want string
	}{
		{
			path: "/home/",
			want: "/home",
		},
	}

	for i, tc := range testCases {
		got := simplifyPath(tc.path)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
