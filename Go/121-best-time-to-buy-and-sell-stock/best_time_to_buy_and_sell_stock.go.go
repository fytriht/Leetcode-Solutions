package maxProfit

func maxProfit(prices []int) int {
	var begin, end, max int
	for end < len(prices) {
		profit := prices[end] - prices[begin]
		if profit <= 0 {
			begin = end
		} else if profit > max {
			max = profit
		}
		end++
	}
	return max
}
