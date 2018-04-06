package maxProfit

// TODO: more efficient

func maxProfit(prices []int) int {
	max := 0
	i, j := 0, len(prices)-1
	for i < j {
		for i < j {
			if d := prices[j] - prices[i]; d > max {
				max = d
			}
			j--
		}
		i++
		j = len(prices) - 1
	}
	return max
}
