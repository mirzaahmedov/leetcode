// https://leetcode.com/problems/interval-list-intersections/?envType=study-plan&id=algorithm-ii
package main

import "fmt"

func main(){
  fmt.Println(intervalIntersection([][]int{ { 1, 3 } }, [][]int{ { 1, 2 } }))
  fmt.Println(intervalIntersection([][]int{{0,2},{5,10},{13,23},{24,25}},[][]int{{1,5},{8,12},{15,24},{25,26}}))
  fmt.Println(intervalIntersection([][]int{{ 1,3 }, { 5,9 }},[][]int{ }))
}
func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
  result := [][]int{}

  i, j := 0, 0   

  for i < len(firstList) && j < len(secondList) {
    a1, b1 := firstList[i][0], firstList[i][1]
    a2, b2 := secondList[j][0], secondList[j][1]

    if a2 >= a1 && a2 <= b1 {
      result = append(result, []int{ a2, min(b1, b2) })
    } else if a1 >= a2 && a1 <= b2 {
      result = append(result, []int{ a1, min(b1, b2) })
    }

    if b1 > b2 {
      j++
    } else {
      i++
    }
    
  }

  return result
}
