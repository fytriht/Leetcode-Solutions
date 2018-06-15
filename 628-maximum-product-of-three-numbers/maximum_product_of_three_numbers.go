package solution

import (
	"math"
	"sort"
)

//
// solution 1
//

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	ret := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]
	if neg := nums[0] * nums[1] * nums[len(nums)-1]; neg > ret {
		return neg
	}
	return ret
}

//
// solution 2
//

func maximumProduct2(nums []int) int {
	min1, min2 := math.MaxInt32, math.MaxInt32
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	for _, n := range nums {
		if n < min1 {
			min1, min2 = n, min1
		} else if n < min2 {
			min2 = n
		}
		if n > max1 {
			max1, max2, max3 = n, max1, max2
		} else if n > max2 {
			max2, max3 = n, max2
		} else if n > max3 {
			max3 = n
		}
	}
	return max(min1*min2*max1, max1*max2*max3)
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
