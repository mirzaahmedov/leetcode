package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		intersection([]int{1, 2, 2, 1}, []int{2, 2}),
	)
	fmt.Println(
		intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}),
	)
}

// func intersection(nums1 []int, nums2 []int) []int {
// 	memo := make(map[int]int)
// 	res := make([]int, 0)
//
// outer:
// 	for i := 0; i < len(nums1); i++ {
// 		if memo[nums1[i]] == 1 {
// 			continue
// 		}
// 		for j := 0; j < len(nums2); j++ {
// 			if nums1[i] == nums2[j] {
// 				res = append(res, nums1[i])
// 				memo[nums1[i]] = 1
// 				continue outer
// 			}
// 		}
// 	}
//
// 	return res
// }

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	r := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = struct{}{}
	}

	for j := 0; j < len(nums2); j++ {
		if _, ok := m[nums2[j]]; ok {
			r = append(r, nums2[j])
			delete(m, nums2[j])
		}
	}

	return r
}
