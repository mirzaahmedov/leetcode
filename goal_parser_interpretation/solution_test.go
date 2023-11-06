package solution

import (
	"testing"
)

func TestInterpret(t *testing.T) {
	testCases := []struct {
		command string
		want    string
	}{
		{
			command: "G()(al)",
			want:    "Goal",
		},
		{
			command: "G()()()()(al)",
			want:    "Gooooal",
		},
		{
			command: "(al)G(al)()()G",
			want:    "alGalooG",
		},
	}

	for i, tc := range testCases {
		got := interpret(tc.command)
		if got != tc.want {
			t.Fatalf("#%d test - expected %v but got %v", i, tc.want, got)
		}
	}
}
