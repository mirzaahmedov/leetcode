package solution

var (
	memo = make(map[int]int)
)

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	return helper(nums, 0)
}
func helper(nums []int, j int) int {
	min := 100000000000000

	if len(nums) < 2 {
		return 1
	}
	if len(nums) == 2 {
		if nums[0] != 0 {
			return 1
		}
		return 0
	}

	for i := 0; i <= nums[0]; i++ {
		if val, ok := memo[j+i]; ok {
			if val < min {
				min = val
			}
			continue
		}

		res := helper(nums[i+1:], j+i)
		memo[j+i] = res
		if res < min {
			min = res
		}
	}

	return min + 1
}
