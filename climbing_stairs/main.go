// https://leetcode.com/problems/climbing-stairs/?envType=study-plan&id=algorithm-i
package main

import "fmt"

func main(){
  fmt.Println(climbStairs(2))
  fmt.Println(climbStairs(3))
  fmt.Println(climbStairs(44))
}
func helper(count *int, n int) {
  if n == 0 {
    *count = *count + 1
    return
  }
  if n < 0 {
    return
  }

  helper(count, n - 1)
  helper(count, n - 2)
}
func climbStairs(n int) int {
  count := 0

  helper(&count, n)

  return count
}
