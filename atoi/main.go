// https://leetcode.com/problems/string-to-integer-atoi/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("2003"))
}
func myAtoi(s string) int {
	count := 0
	result := 0
	isNegative := false
	isPositive := true

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 48 && s[i] <= 57 {
			result += int(s[i]-48) * int(math.Pow10(count))
			count++
		} else if s[i] == 45 {
			if isPositive {
				return 0
			}
			isNegative = true
		} else if s[i] == 32 {
			continue
		} else if s[i] == 43 {
			if isNegative {
				return 0
			}
			isPositive = true
		} else if s[i] == 46 {
			result /= int(math.Pow10(count))
			count = 0
		} else if count > 0 {
			return 0
		}

	}

	if isNegative {
		result = -result
	}
	if result < -2147483648 {
		result = -2147483648
	}
	if result > 2147483647 {
		result = 2147483647
	}

	return result
}
