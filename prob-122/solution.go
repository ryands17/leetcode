package prob122

func MaxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	var profit int
	currentPrice := prices[0]
	for i := 1; i < l; i++ {
		if prices[i] > currentPrice {
			profit += (prices[i] - currentPrice)
		}
		currentPrice = prices[i]
	}
	return profit
}
