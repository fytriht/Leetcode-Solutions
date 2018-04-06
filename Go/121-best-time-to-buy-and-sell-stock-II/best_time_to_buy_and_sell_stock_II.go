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

func maxProfit2(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		profit := prices[i+1] - prices[i]
		if profit > 0 {
			max += profit
		}
	}
	return max
}
