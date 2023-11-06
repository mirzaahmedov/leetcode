package solution

func uniqueOccurrences(nums []int) bool {
	counts := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		counts[nums[i]] += 1
	}

	results := make(map[int]int)

	for num, count := range counts {
		if _, ok := results[count]; ok {
			return false
		}
		results[count] = num
	}

	return true
}
