// https://leetcode.com/problems/shortest-path-in-binary-matrix/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(shortestPathBinaryMatrix([][]int{{0,0,0},{1,1,0},{1,1,1}}))
}
func shortestPathBinaryMatrix(grid [][]int) int {
  if len(grid) == 0 || grid[0][0] == 1 {
    return -1
  }
  if len(grid) == 1 {
    return 1
  }

  var (
    queue = [][]int{{0, 0}}
  )

  grid[0][0] = 1

  directions := [][]int{ { -1, -1 }, { -1, 0 }, { -1, 1 }, { 0, -1 }, { 0, 1 }, { 1, -1 }, { 1, 0 }, { 1, 1 } }
  for len(queue) > 0 {
    for d := range directions {
      x := queue[0][0] + directions[d][0]
      y := queue[0][1] + directions[d][1]      

      if x == len(grid)-1 && y == len(grid[0])-1 && grid[x][y] == 0 {
        fmt.Println(grid , x, y, queue)
        return grid[queue[0][0]][queue[0][1]] + 1
      }
      if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 0 {
        queue = append(queue, []int{x, y})
        grid[x][y] = grid[queue[0][0]][queue[0][1]] + 1
      }
    }
    queue = queue[1:]
  }

  return -1
}
