package solution

import (
	"math"
)

func minMoves(nums []int) int {
	sum, min := 0, math.MaxInt32
	for _, n := range nums {
		if n < min {
			min = n
		}
		sum += n
	}
	return sum - len(nums)*min
}
