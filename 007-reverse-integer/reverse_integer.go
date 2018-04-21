package solution

import "math"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		x = x / 10
	}
	if ret < math.MinInt32 || ret > math.MaxInt32 {
		return 0
	}
	return ret
}
