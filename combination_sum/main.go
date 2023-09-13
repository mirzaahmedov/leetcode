package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("[2,3,6,7] => 7", combinationSum([]int{2, 3, 6, 7}, 7))
	// fmt.Println("[2,3,5] => 8", combinationSum([]int{2, 3, 5}, 8))

	// combinationSum([]int{2, 3, 6, 7}, 7)
	// fmt.Println(combinationSum([]int{2, 3}, 6))
	for _, v := range combinationSum([]int{2, 3, 4}, 9) {
		fmt.Println(v)
	}
}
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return recurse(candidates, target, 0)
}
func recurse(candidates []int, target int, start int) [][]int {
	var r [][]int

	for i := start; i < len(candidates); i++ {
		d := target - candidates[i]

		if d < 0 {
			continue
		}
		if d == 0 {
			r = append(r, []int{candidates[i]})
			continue
		}

		for _, v := range recurse(candidates, d, i) {
			v = append(v, candidates[i])
			r = append(r, v)
		}
	}

	return r
}
