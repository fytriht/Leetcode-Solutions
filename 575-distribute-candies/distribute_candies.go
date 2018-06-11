package solution

func distributeCandies(candies []int) int {
	m := map[int]bool{}
	for _, c := range candies {
		m[c] = true
	}
	return min(len(m), len(candies)/2)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
