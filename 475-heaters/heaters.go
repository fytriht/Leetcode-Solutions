package solution

import (
	"sort"
)

// TODO: gofmt
// TODO: test

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var ret, i int
	for _, h := range houses {
		for i < len(heaters) - 1 && abs(heaters[i +1] - h) <= abs(heaters[i] - h) {
			i++
		}
		ret = max(ret, abs(heaters[i] - h)) 
	}
	return ret
}

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