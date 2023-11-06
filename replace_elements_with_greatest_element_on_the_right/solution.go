package solution

func replaceElements(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	results := make([]int, len(nums))
	maxValue := -1
	maxIndex := 0

	for i := 0; i < len(nums); i++ {
		if maxIndex <= i {
			max := -1
			index := 0
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > max {
					max = nums[j]
					index = j
				}
			}
			maxValue = max
			maxIndex = index
		}
		results[i] = maxValue
	}

	return results
}
