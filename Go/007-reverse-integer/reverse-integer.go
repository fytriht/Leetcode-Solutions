package twoSum

import "math"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		t := x % 10
		ret = ret*10 + t
		x = x / 10
	}
	if ret < math.MinInt32 || ret > math.MaxInt32 {
		return 0
	}
	return ret
}
