// https://leetcode.com/problems/number-of-islands/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(numIslands([][]byte{
    []byte("11110"),
    []byte("11010"),
    []byte("11000"),
    []byte("00000"),
  }))
  fmt.Println(numIslands([][]byte{
    []byte("11000"),
    []byte("11000"),
    []byte("00100"),
    []byte("00011"),
  }))
}
func bfs(x int, y int, grid [][]byte){
    queue := make([][]int, 0)
    directions := [][]int{ { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } }

    grid[x][y] = '0'
    queue = append(queue, []int{ x, y })

    for len(queue) > 0 {
      cell := queue[0]
      queue = queue[1:]

      for _, d := range directions {
        x := cell[0] + d[0]
        y := cell[1] + d[1]

        if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
          grid[x][y] = '0'
          queue = append(queue, []int{ x, y })
        }
      }
    }

}
func numIslands(grid [][]byte) int {
    count := 0

    for i := 0; i < len(grid); i++ {
      for j := 0; j < len(grid[i]); j++ {
        if grid[i][j] == '1' {
          bfs(i, j, grid)
          count++
        }
      }
    }

    return count
}
