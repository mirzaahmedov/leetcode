package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
	fmt.Println(addDigits(0))
	fmt.Println(addDigits(8))
	fmt.Println(addDigits(200))
}
func addDigits(num int) int {
	if num < 10 {
		return num
	}

	sum := 0

	for num >= 1 {
		sum += num % 10
		num = num / 10
	}

	return addDigits(sum)
}
