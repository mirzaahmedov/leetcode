package solution

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if len(s) == 1 {
		return s == goal
	}

	count := len(goal)

	for i := 0; i < len(s); i++ {
		if s[i] == goal[0] {
			j, k := 1, i+1
			count -= 1

			if k == len(s) {
				k = 0
			}

			for j < len(goal) {
				if goal[j] == s[k] {
					j, k = j+1, k+1
					count -= 1
				} else {
					count = len(goal)
					break
				}

				if k >= len(s) {
					k = 0
				}
				if count == 0 {
					return true
				}

			}
		}
	}

	return false
}
