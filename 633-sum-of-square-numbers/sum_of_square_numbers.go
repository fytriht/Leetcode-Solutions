package solution

import (
	"math"
)

func judgeSquareSum(c int) bool {
	for i := 0; i*i <= c; i++ {
		j := int(math.Sqrt(float64(c - i*i)))
		if i*i+j*j == c {
			return true
		}
	}
	return false
}
