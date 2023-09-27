package solution

func intersect(nums1 []int, nums2 []int) []int {
	results := []int{}
	counts := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		counts[nums1[i]] += 1
	}

	for i := 0; i < len(nums2); i++ {
		if counts[nums2[i]] > 0 {
			results = append(results, nums2[i])
			counts[nums2[i]] -= 1
		}
	}

	return results
}
