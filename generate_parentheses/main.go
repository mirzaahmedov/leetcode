// https://leetcode.com/problems/generate-parentheses/

package main

import "fmt"

func main() {
  results := generateParenthesis(3)
  fmt.Println(results)
}

var randoms []string

func recursive(value string, n int) {
  if n == 0 {
    randoms = append(randoms, value)
    return
  }

  recursive(value + "(", n - 1)
  recursive(value + ")", n - 1)
}
func isValid(value string) bool {
  count := 0
  for _, v := range value {
    if v == '(' {
      count++
    } else {
      count--
    }
    if count < 0 {
      return false
    }
  }
  return count == 0
}
func generateParenthesis(n int) []string {
  results := []string{}

  recursive("", n * 2)

  for i := 0; i < len(randoms); i++ {
    if isValid(randoms[i]) {
      results = append(results, randoms[i])
    }
  }

  return results
}
