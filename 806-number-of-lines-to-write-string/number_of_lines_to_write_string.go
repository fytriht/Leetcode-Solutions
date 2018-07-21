package solution

func numberOfLines(widths []int, S string) []int {
	ret := []int{1, 0}
	for _, r := range S {
		w := widths[r-'a']
		ret[1] += w
		if ret[1] > 100 {
			ret[0]++
			ret[1] = w
		}
	}
	return ret
}
