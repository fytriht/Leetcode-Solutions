package solution

import (
	"math"
)

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	minA, minB := math.MaxInt32, math.MaxInt32
	for _, op := range ops {
		minA, minB = min(minA, op[0]), min(minB, op[1])
	}
	return minA * minB
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
