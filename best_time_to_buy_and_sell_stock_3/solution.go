package solution

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	result := 0
	indexes := [2]int{}
	maxProfit := 0
	minPrice := prices[0]
	minPriceIndex := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
			indexes[0], indexes[1] = minPriceIndex, i
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
			minPriceIndex = i
		}
	}

	result += maxProfit

	maxProfit = 0
	if indexes[0] != 0 {
		minPrice = prices[0]
	} else {
		minPrice = indexes[1] + 1
	}

	if result == 0 {
		return result
	}
	fmt.Println(indexes, result)

	for i := 1; i < indexes[0]; i++ {
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = i
		}
	}
	for i := indexes[1] + 1; i < len(prices); i++ {
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return result + maxProfit
}
