// https://leetcode.com/problems/number-of-1-bits/?envType=study-plan&id=algorithm-i
package main

import (
  "fmt"
)

func main() {
  fmt.Println(hammingWeight(11))
}
func hammingWeight(num uint32) int {
  var (
    count int
    i uint32
  )
  
  for i = 0; i < 32; i++ {
    if (num & (1 << i)) == (1 << i) {
      count++
    }
  }

  return count
}
