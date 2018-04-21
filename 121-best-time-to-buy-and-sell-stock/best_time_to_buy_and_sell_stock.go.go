package solution

import (
	"math"
)

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

func maxProfit2(prices []int) int {
	maxProfit, minPrice := 0, math.MaxInt32
	for _, p := range prices {
		minPrice = min(minPrice, p)
		maxProfit = max(maxProfit, p-minPrice)
	}
	return maxProfit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
