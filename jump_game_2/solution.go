package solution

// my solution not so fast
// func jump(nums []int) int {
// 	max := 0
// 	results := make([]int, len(nums))
//
// 	for i := 0; i < len(nums); i++ {
// 		if i > max {
// 			return 0
// 		}
// 		for j := 1; i+j < len(nums) && j <= nums[i]; j++ {
// 			if i+j > max {
// 				max = i + j
// 			}
//
// 			if results[i+j] == 0 || results[i+j] > results[i]+1 {
// 				results[i+j] = results[i] + 1
// 			}
// 		}
// 	}
//
// 	return results[len(results)-1]
// }

func jump(nums []int) int {
	result := 0
	totalMax := 0
	currentMax := 0

	for i := 0; i < len(nums) && totalMax < len(nums)-1; {
		currentMax = 0

		for ; i <= totalMax; i++ {
			currentMax = max(currentMax, i+nums[i])
		}

		totalMax = currentMax
		result++
	}

	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
