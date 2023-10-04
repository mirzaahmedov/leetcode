package solution

import (
	"testing"

	"github.com/mirzaahmedov/leetcode/utils"
)

func TestMyHashMap(t *testing.T) {
	testCases := []struct {
		commands []string
		args     [][]int
		want     []*int
	}{
		{
			commands: []string{"MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"},
			args:     [][]int{{}, {1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2}},
			want:     []*int{nil, nil, nil, utils.Pointer(1), utils.Pointer(-1), nil, utils.Pointer(1), nil, utils.Pointer(-1)},
		},
		{
			commands: []string{"MyHashMap", "put", "put", "put", "get", "get", "remove", "get"},
			args:     [][]int{{}, {1, 1}, {10001, 2}, {10001, 3}, {1}, {10001}, {10001}, {10001}},
			want:     []*int{nil, nil, nil, nil, utils.Pointer(1), utils.Pointer(3), nil, utils.Pointer(-1)},
		},
	}

	for j, tc := range testCases {
		var (
			hashMap MyHashMap
			results []*int
		)

		for i, cmd := range tc.commands {
			switch cmd {
			case "MyHashMap":
				hashMap = Constructor()
				results = []*int{nil}
			case "put":
				hashMap.Put(tc.args[i][0], tc.args[i][1])
				results = append(results, nil)
			case "get":
				results = append(results, utils.Pointer(hashMap.Get(tc.args[i][0])))
			case "remove":
				hashMap.Remove(tc.args[i][0])
				results = append(results, nil)
			}
		}

		if len(results) != len(tc.want) {
			t.Log("length of the arrays does not match")
			t.Fatalf("#%d test - expected but %v got %v", j, tc.want, results)
		}

		for i := 0; i < len(tc.want); i++ {
			if results[i] == nil && tc.want[i] == nil {
				continue
			}
			if *results[i] != *tc.want[i] {
				t.Log(i, *results[i], *tc.want[i])
				t.Fatalf("#%d test - i = %d expected %v but got %v", j, i, *tc.want[i], *results[i])
			}
		}
	}
}
