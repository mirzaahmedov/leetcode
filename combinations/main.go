// https://leetcode.com/problems/combinations/?envType=study-plan&id=algorithm-i
package main

import "fmt"

func main(){
  fmt.Println(combine(4, 2))
}

func helper(results *[][]int, n int, k int, prefix []int){
    if k <= 0 {
      temp := make([]int, len(prefix))
      copy(temp, prefix)
      *results = append(*results, temp)
      return
    }

    for i := prefix[len(prefix) - k] + 1; i <= n; i++ {
      fmt.Println(k, i)
      prefix[len(prefix) - k] = i
      helper(results, n, k - 1, prefix)
    }
}
func combine(n int, k int) [][]int {
  results := [][]int{}

  helper(&results, n, k, make([]int, k))
    
  return results
}
