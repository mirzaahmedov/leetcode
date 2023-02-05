// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extra := 3

	fmt.Println(kidsWithCandies(candies, extra))

}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0

	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	res := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		res[i] = candies[i]+extraCandies >= max
	}

	return res
}
