// https://leetcode.com/problems/power-of-two/?envType=study-plan&id=algorithm-i
package main

import (
  "fmt"
)

func main(){
  fmt.Println(isPowerOfTwo(8))
  fmt.Println(isPowerOfTwo(5))
  fmt.Println(isPowerOfTwo(2))
  fmt.Println(isPowerOfTwo(10))
  fmt.Println(isPowerOfTwo(80))
}
func isPowerOfTwo(n int) bool {
  num := 1
  for num <= n {
    if num == n {
      return true
    }
    num = num << 1
  }

  return false
}
