package solution

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	pre, tmpLen, ret := nums[0], 1, 1
	for _, n := range nums[1:] {
		if n > pre {
			tmpLen++
			ret = max(ret, tmpLen)
		} else {
			tmpLen = 1
		}
		pre = n
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
