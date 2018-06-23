package solution

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tmpLen, ret := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tmpLen++
			ret = max(ret, tmpLen)
		} else {
			tmpLen = 1
		}
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
