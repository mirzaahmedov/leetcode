package main

import "fmt"

func getDigits(n int) []int {
	digits := []int{}

	for n >= 10 {
		digits = append(digits, n%10)
		n = n / 10
	}

	digits = append(digits, n)

	return digits
}
func isHappy(n int) bool {
	memo := map[int]bool{}

	for true {
		result := 0
		digits := getDigits(n)

		for _, digit := range digits {
			result += digit * digit
		}

		if result == 1 {
			return true
		}
		if memo[result] {
			return false
		}

		memo[result] = true
		n = result
	}

	return false
}
func main() {
	fmt.Println(isHappy(2), "== false")
	fmt.Println(isHappy(19), "== true")
	fmt.Println(isHappy(1), "== true")
}
