// https://leetcode.com/problems/backspace-string-compare/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  // fmt.Println(normalize("123#####4"))
  // fmt.Println(normalize("123####4"))
  // fmt.Println(normalize("123###4"))
  // fmt.Println(normalize("123##4"))
  // fmt.Println(normalize("c#d#"))
  // fmt.Println(normalize("ab##"))
  fmt.Println(backspaceCompare("a##c","#a#c"))
  fmt.Println(backspaceCompare("y#fo##f","y#f#o##f"))
  fmt.Println(normalize("y#fo##f"), normalize("y#f#o##f"))
}
func normalize(s string) string {
  count := 0

  for i := len(s) - 1; i >= 0;i-- {
    if s[i] == '#' {
      count++
      s = s[:i] + s[i + 1:]
    } else {
      if count > 0 {
        count--
        s = s[:i] + s[i + 1:]
      }
    }
  }

  return s
}
func backspaceCompare(s string, t string) bool {
  return normalize(s) == normalize(t)
}
