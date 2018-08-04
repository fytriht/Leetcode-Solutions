package solution

func maxArea(height []int) int {
	ret := 0
	for i, j := 0, len(height)-1; i < j; {
		h := min(height[i], height[j])
		ret = max(ret, h*(j-i))
		for i < j && height[i] <= h {
			i++
		}
		for i < j && height[j] <= h {
			j--
		}
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
