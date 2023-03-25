// https://leetcode.com/problems/climbing-stairs/?envType=study-plan&id=algorithm-i
package main

import (
  "fmt"
)

func main(){
  fmt.Println(climbStairs(2))
  fmt.Println(climbStairs(3))
}
func climbStairs(n int) int {
  table := make([]int, n + 1)

  if n < 3 {
    return n
  }

  table[1] = 1
  table[2] = 2

  for i := 3; i < n + 1; i++ {
    table[i] = table[i - 1] + table[i - 2]
  }

  return table[n]
}
