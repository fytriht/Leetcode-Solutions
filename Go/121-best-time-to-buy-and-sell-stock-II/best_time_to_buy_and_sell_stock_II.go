package maxProfit

func maxProfit(prices []int) int {
	max, tmpMax := 0, 0
	for begin, end := 0, 0; end < len(prices); end++ {
		profit := prices[end] - prices[begin]
		if profit > tmpMax {
			tmpMax = profit
		}
		if end+1 == len(prices) || prices[end+1] < prices[end] {
			max += tmpMax
			tmpMax = 0
			begin = end + 1
		}
	}
	return max
}
