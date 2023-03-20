// https://leetcode.com/problems/max-area-of-island/?envType=study-plan&id=algorithm-i

package main

func main(){

}

func traverse(grid [][]int, i int, j int, area int) int {
  if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
    return 0
  }

  grid[i][j] = 0
  area++

  area += traverse(grid, i + 1, j, area)
  area += traverse(grid, i - 1, j, area)
  area += traverse(grid, i, j + 1, area)
  area += traverse(grid, i, j - 1, area)

  return area
}
func maxAreaOfIsland(grid [][]int) int {
  max := 0

  for i := 0; i < len(grid); i ++ {
    for j := 0; j < len(grid[0]); j++ {
      if grid[i][j] == 1 {
        area := traverse(grid, i, j, 0)

        if area > max {
          max = area
        }
      }
    }
  }

  return max
}
