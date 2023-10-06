package solution

func integerBreak(n int) int {
	if n == 3 {
		return 2
	}
	if n == 2 {
		return 1
	}

	prod := 1

	for n-3 >= 0 {
		if n == 4 {
			n -= 2
			prod *= 2
			break
		}
		n -= 3
		prod *= 3
	}
	for n-2 >= 0 {
		n -= 2
		prod *= 2
	}

	return prod
}
