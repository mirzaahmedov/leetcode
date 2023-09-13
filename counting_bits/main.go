package main

import (
	"fmt"
)

func main() {
	// fmt.Println("count:", 2, count(2))
	// fmt.Println("count:", 3, count(3))
	// fmt.Println("count:", 4, count(4))
	// fmt.Println("count:", 5, count(5))

	fmt.Println("countBits:", 2, countBits(2))
	fmt.Println("countBits:", 3, countBits(3))
	fmt.Println("countBits:", 4, countBits(4))
	fmt.Println("countBits:", 5, countBits(5))

	// fmt.Println("count:", 2, toBinary(2))
	// fmt.Println("count:", 3, toBinary(3))
	// fmt.Println("count:", 4, toBinary(4))
	// fmt.Println("count:", 5, toBinary(5))
}

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	ans := make([]int, n+1)

	ans[0] = 0
	ans[1] = 1

	count := func(n int) int {
		count := 0

		for n > 0 {
			if ans[n] != 0 {
				return count + ans[n]
			}
			if n%2 != 0 {
				count += 1
			}
			n = n / 2
		}

		return count
	}

	for i := 2; i <= n; i++ {
		ans[i] = count(i)
	}

	return ans
}

func toBinary(n int) string {
	res := ""

	for n > 0 {
		if n%2 != 0 {
			res = "1" + res
		} else {
			res = "0" + res
		}

		n = n / 2
	}

	return res
}
