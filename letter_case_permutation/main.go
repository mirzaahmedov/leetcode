// https://leetcode.com/problems/letter-case-permutation/?envType=study-plan&id=algorithm-i
package main

import "fmt"

func main(){
  fmt.Println(letterCasePermutation("a1b2"))
}
func copyAndAppend(str string, c byte) string {
  s := []byte(str)
  result := make([]byte, len(s) + 1)
  for i, letter := range s {
    result[i] = letter
  }
  result[len(s)] = c
  
  return string(result)
}
func helper(results *[]string, prefix string,s string){
  if len(s) == 0 {
    *results = append(*results, prefix)
    return
  }

  letter := s[0]

  if letter >= 97 && letter <= 122 {
    helper(results, copyAndAppend(prefix, letter - 32), s[1:])
  }
  if letter >= 65 && letter <= 90 {
    helper(results, copyAndAppend(prefix, letter + 32), s[1:])
  }

  helper(results, copyAndAppend(prefix, letter), s[1:])
}
func letterCasePermutation(s string) []string {
  results := []string{}

  helper(&results, "", s)
  
  return results
}
