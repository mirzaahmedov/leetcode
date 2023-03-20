// https://leetcode.com/problems/01-matrix/description/
package main

import "fmt"

func main(){
  data := [][]int{ 
    { 0,0,0 },
    { 0,1,0 },
    { 1,1,1 },
  }

  fmt.Println(updateMatrix(data))
}

func updateMatrix(mat [][]int) [][]int {
  queue := [][]int{}

  for i := 0; i < len(mat); i++ {
    for j := 0; j < len(mat[i]); j++ {
      if mat[i][j] == 0 {
        queue = append(queue, []int{ i, j })
      } else {
        mat[i][j] = -1
      }
    }
  }

  for i := 0; i < len(queue); i++ {
    row, column := queue[i][0], queue[i][1]

    for _, v := range [][]int{ { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } } { 
      dx, dy := v[0], v[1]
      nr := row + dx
      nc := column + dy

      if 0 <= nr && nr < len(mat) && 0 <= nc && nc < len(mat[0]) && mat[nr][nc] == -1 {
        mat[nr][nc] = mat[row][column] + 1
        queue = append(queue, []int{ nr, nc })
      }
    }
  }

  return mat
}
