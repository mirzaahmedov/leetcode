// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/

package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n > 9 {
		product *= n % 10
		sum += n % 10
		n = n / 10
	}

	product *= n
	sum += n

	return product - sum
}
