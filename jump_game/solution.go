package solution

// simple solution but slow
//
//	func canJump(nums []int) bool {
//		if len(nums) == 1 {
//			return true
//		}
//		for i := 1; i < len(nums) && i <= nums[0]; i++ {
//			res := canJump(nums[i:])
//			if res {
//				return true
//			}
//		}
//
//		return false
//	}
// func canJump(nums []int) bool {
// 	if len(nums) == 0 {
// 		return false
// 	}
//
// 	max := 0
// 	results := make([]bool, len(nums))
// 	results[0] = true
//
// 	for i := 0; i < len(nums)-1; i++ {
// 		if !results[i] {
// 			continue
// 		}
// 		for j := max - i + 1; i+j < len(nums) && j <= nums[i]; j++ {
// 			if i+j == len(nums)-1 {
// 				return true
// 			}
// 			if i+j > max {
// 				max = i + j
// 			}
// 			if !results[i+j] {
// 				results[i+j] = true
// 			}
// 		}
// 	}
//
// 	return results[len(results)-1]
// }

// solution copied from submissions! very smart!
func canJump(nums []int) bool {
	max := 0

	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}

		if i+nums[i] > max {
			max = i + nums[i]

			if max > len(nums)-1 {
				return true
			}
		}
	}

	return true
}
