package maxProfit

func maxProfit(prices []int) int {
	max, tmpMax, begin, end := 0, 0, 0, 0
	for end < len(prices) {
		profit := prices[end] - prices[begin]
		if profit > tmpMax {
			tmpMax = profit
		}
		if end+1 < len(prices) && prices[end+1] < prices[end] {
			max += tmpMax
			tmpMax = 0
			begin = end + 1
		}
		end++
	}
	max += tmpMax
	return max
}
