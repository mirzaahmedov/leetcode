// https://leetcode.com/problems/reverse-words-in-a-string-iii/?envType=study-plan&id=algorithm-i

package main

import "fmt"

func main() {
    fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
  r := []byte(s)
  count := 0

  for i := 0; i < len(r); i++ {
    if r[i] == ' ' {
      for j, k := i - count, i - 1; k > j; j, k = j + 1, k - 1 {
        r[j], r[k] = r[k], r[j]
      }
      count = 0
    } else {
      count++
    }
  }
  for j, k := len(r) - count, len(r) - 1; k > j; j, k = j + 1, k - 1 {
    r[j], r[k] = r[k], r[j]
  }

  return string(r)
} 
