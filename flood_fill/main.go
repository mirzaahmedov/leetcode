// https://leetcode.com/problems/flood-fill/?envType=study-plan&id=algorithm-i
package main

import "fmt"

func main(){
  image := [][]int{ { 1,1,1 },{ 1,1,0 },{ 1,0,1 } }
  sr := 1 
  sc := 1 
  color := 2

  floodFill(image, sr, sc, color)

  fmt.Println(image)
}

func recurse(image [][]int, sr int, sc int, color int, currentColor int){
  if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != currentColor || image[sr][sc] == color {
    return
  }

  image[sr][sc] = color

  recurse(image, sr - 1, sc, color, currentColor)
  recurse(image, sr + 1, sc, color, currentColor)
  recurse(image, sr, sc - 1, color, currentColor)
  recurse(image, sr, sc + 1, color, currentColor)
}
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
  recurse(image, sr, sc, color, image[sr][sc])
  return image
}
