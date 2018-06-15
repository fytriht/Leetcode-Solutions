package solution

import (
	"math"
)

func judgeSquareSum(c int) bool {
	m := int(math.Sqrt(float64(c / 2)))
	left, right := m, m
	for left >= 0 && right*right <= c {
		sum := left*left + right*right
		if sum > c {
			left--
		} else if sum < c {
			right++
		} else {
			return true
		}
	}
	return false
}
