// https://leetcode.com/problems/permutation-in-string/?envType=study-plan&id=algorithm-i

package main

import "fmt"

func main(){
  cases := [][]string{
    {"adc", "dcda"},
    {"ab", "eidbaooo"},
    {"ab", "eidboaoo"},
    {"hello", "ooolleoooleh"},
  }

  for _, c := range cases {
    fmt.Println(checkInclusion(c[0], c[1]))
  }
}

func checkInclusion(s1 string, s2 string) bool {
  memo := [129]int{} 
  l := 0
  r := 0
  
  for i := 0; i < len(s1); i++ {
    if memo[s1[i]] == 0 {
      memo[s1[i]] = 1
    }
    memo[s1[i]]++
  }

  for ; r < len(s2); r++ {
    if r - l == len(s1) {
      return true
    }

    if memo[s2[r]] == 0 {
      for ; l < r; l++ {
        memo[s2[l]]++
      }
      l++
    } else if memo[s2[r]] == 1 {
      for ; memo[s2[r]] == 1; l++ {
        memo[s2[l]]++
      }
      memo[s2[r]]--
    } else {
      memo[s2[r]]--
    }
  }

  return r - l == len(s1)
}
