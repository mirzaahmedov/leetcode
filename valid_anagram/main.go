// https://leetcode.com/problems/valid-anagram/
package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

const NUR_OF_CHARS = 256

func isAnagram(a string, b string) bool {
  if len(a) != len(b) {
    return false
  }

  count := [NUR_OF_CHARS]int{}

  for i := 0; i < len(a); i++ {
    count[int(a[i])]++
  }
  for i := 0; i < len(b); i++ {
    count[int(b[i])]--
    if count[int(b[i])] < 0 {
      return false
    }
  }

  return true
}
