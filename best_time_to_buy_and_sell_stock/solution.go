package solution

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
}
