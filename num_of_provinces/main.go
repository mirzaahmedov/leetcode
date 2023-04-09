// https://leetcode.com/problems/number-of-provinces/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(findCircleNum([][]int{{1,1,0},{1,1,0},{0,0,1}}))
}
func dfs(isConnected [][]int, i int) {
  for j := 0; j < len(isConnected); j++ {
    if isConnected[i][j] == 1 {
      isConnected[i][j] = 0
      isConnected[j][i] = 0
      dfs(isConnected, j)
    }
  }
}
func findCircleNum(isConnected [][]int) int {
  count := 0

  for i := 0; i < len(isConnected); i++ {
    if isConnected[i][i] == 1 {
      count++
      dfs(isConnected, i)
    }
  }

  return count
}
