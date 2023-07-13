// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

outer:
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if i+j >= len(haystack) || haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}

	return -1
}
func main() {
	fmt.Println(strStr("leetcode", "leeto"))
}
