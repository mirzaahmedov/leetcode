// https://leetcode.com/problems/multiply-strings/

package main

import "fmt"

func main() {
		fmt.Println(multiply("11", "11"))
}

func multiplyDigits(a int, b int) (int, int) {
	product := a * b
	reminder := product / 10 
	carry := product % 10

	return carry, reminder
}
func multiply(a string, b string) string {
	maxLength := 0
	res := []byte{}
	
	if (len(a) > len(b) ) {
		maxLength = len(a)
	} else {
		maxLength = len(b)
	}

	for i := 0; i < maxLength; i++ {
		p, _ := multiplyDigits(int(a[i] - 48), int(b[i] - 48))

		res = append(res, byte(p) + 48)
	}

	return string(res)
}