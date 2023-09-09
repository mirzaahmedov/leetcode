package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(nums)

	sortColors(nums)

	fmt.Println(nums)

	nums2 := []int{2, 0, 1}
	fmt.Println(nums2)

	sortColors(nums2)

	fmt.Println(nums2)
}

//	func sortColors(nums []int) {
//		for i := 0; i < len(nums); i++ {
//			for j := 1; j < len(nums)-i; j++ {
//				if nums[j-1] > nums[j] {
//					nums[j-1], nums[j] = nums[j], nums[j-1]
//				}
//			}
//		}
//	}
// func sortColors(nums []int) {
// 	l, r := 0, len(nums)-1
//
// 	for i := 0; i <= r; i++ {
// 		switch true {
// 		case nums[i] == 0 && i != l:
// 			nums[i], nums[l] = nums[l], nums[i]
// 			l++
// 		case nums[i] == 2 && i != r:
// 			nums[i], nums[r] = nums[r], nums[i]
// 			r--
// 		}
// 	}
// }

// func sortColors(nums []int) {
// 	i, l, r := 0, 0, len(nums)-1
//
// 	for i <= r {
// 		if nums[i] == 2 && i != r {
// 			nums[i], nums[r] = nums[r], nums[i]
// 			r -= 1
// 		} else if nums[i] == 0 && i != l {
// 			nums[i], nums[l] = nums[l], nums[i]
// 			l += 1
// 		} else {
// 			i += 1
// 		}
// 	}
// }

func sortColors(nums []int) {
	var (
		i int = 0
		l int = 0
		r int = len(nums) - 1
		t int
	)

	for i <= r {
		switch true {
		case nums[i] == 2 && i != r:
			t = nums[i]
			nums[i] = nums[r]
			nums[r] = t
			r -= 1
		case nums[i] == 0 && i != l:
			t = nums[i]
			nums[i] = nums[l]
			nums[l] = t
			l += 1
		default:
			i += 1
		}
	}
}
