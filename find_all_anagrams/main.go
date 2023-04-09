// https://leetcode.com/problems/find-all-anagrams-in-a-string/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(findAnagrams("cbaebabacd", "abc"))
  fmt.Println(findAnagrams("abab", "ab"))
}
// func isAnagram(s string, t string, l int, r int) bool {
//   count := [256]int{}
//
//   for i := 0; i < len(t); i++ {
//     count[t[i] - 'a']++
//   }
//   for j := l; j < r; j++ {
//     count[s[j] - 'a']--
//     if count[s[j] - 'a'] < 0 {
//       return false
//     }
//   }
//
//   return true
// }
// func findAnagrams(s string, p string) []int {
//   results := []int{}
//
//   for i := 0; i <= len(s) - len(p); i++ {
//     if isAnagram(s, p, i, i + len(p)) {
//       results = append(results, i)
//     }
//   }
//
//   return results
// }
func findAnagrams(s string, p string) []int {
  results := []int{}
  count := [26]int{}
  l, r := 0, 0
  
  for i := 0; i < len(p); i++ {
    count[p[i] - 'a']++
  }

  for ;r < len(s); r++ {
    ch := s[r] - 'a'

    count[ch]--

    for count[ch] < 0 {
      count[s[l] - 'a']++
      l++
    }

    if r - l == len(p) - 1 {
      results = append(results, l)
    }
  }

  return results
}
