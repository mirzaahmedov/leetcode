// https://leetcode.com/problems/rotting-oranges/?envType=study-plan&id=algorithm-i
package main

import "fmt"

func main(){
  oranges := [][]int{ {2,1,1},{1,1,0},{0,1,1} }
  // oranges := [][]int{ {2,1,1},{0,1,1},{1,0,1} }

  fmt.Println(orangesRotting(oranges))
}

func orangesRotting(grid [][]int) int {
  queue := [][]int{}
  time := 0
  fresh := 0

  for i := range grid {
    for j := range grid[i] {
      if grid[i][j] == 1 {
        fresh++
      }
      if grid[i][j] == 2 {
        queue = append(queue, []int{i, j})
      }
    }
  }

  directions := [][]int{ {0, 1}, {0, -1}, {1, 0}, {-1, 0} }

  for len(queue) > 0 && fresh != 0 {
    length := len(queue)
    newQueue := [][]int{}

    for i := 0; i < length; i++ {

      for _, direction := range directions {
        x := queue[i][0] + direction[0]
        y := queue[i][1] + direction[1]

        if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
          grid[x][y] = 2
          fresh--
          newQueue = append(newQueue, []int{x, y})
        }
      }
    }
    
    queue = newQueue
    time++
  }

  if fresh == 0 {
    return time
  } else {
    return -1
  }
}
