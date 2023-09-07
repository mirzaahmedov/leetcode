package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("some", "yaya"))
	fmt.Println(isIsomorphic("paper", "title"))
}
func isIsomorphic(s string, t string) bool {
	mappings := make(map[byte]byte)
	reverse := make(map[byte]byte)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if mappings[s[i]] == 0 {
			mappings[s[i]] = t[i]
			if reverse[t[i]] != 0 {
				return false
			}
			reverse[t[i]] = s[i]
			continue
		}
		if mappings[s[i]] != t[i] {
			return false
		}
	}

	return true
}
