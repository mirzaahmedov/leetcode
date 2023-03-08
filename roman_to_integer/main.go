// https://leetcode.com/problems/roman-to-integer/
package main

import (
  "fmt"
)


func main(){
  fmt.Println(romanToInt("III"), 3)
  fmt.Println(romanToInt("LVIII"), 58)
  fmt.Println(romanToInt("MCMXCIV"), 1994)
}
func romanToInt(s string) int {
  values := map[byte]int {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
  }

  value := 0

  for i := len(s) - 1; i >= 0; i-- {
    if i < len(s) - 1 && values[s[i]] < values[s[i + 1]] {
      value -= values[s[i]]
    } else {
      value += values[s[i]]
    }
  }

  return value
}
