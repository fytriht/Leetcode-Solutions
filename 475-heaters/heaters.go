package solution

import (
	"math"
	"sort"
)

// TODO: test

//
// solution 1
//

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var ret, i int
	for _, h := range houses {
		for i < len(heaters)-1 && abs(heaters[i+1]-h) <= abs(heaters[i]-h) {
			i++
		}
		ret = max(ret, abs(heaters[i]-h))
	}
	return ret
}

//
// solution 2
//

func findRadius2(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var ret int
	for _, h := range houses {
		i := sort.Search(len(heaters), func(i int) bool {
			return heaters[i] >= h
		})
		r := math.MaxInt32
		if i == len(heaters) {
			r = h - heaters[i-1]
		} else if i == 0 {
			r = heaters[i] - h
		} else {
			r = min(heaters[i]-h, h-heaters[i-1])
		}
		ret = max(ret, r)
	}
	return ret
}

//
// util
//

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
