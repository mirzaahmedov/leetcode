package solution

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	count := 0

	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			count += 1
		}
	}

	return count
}
