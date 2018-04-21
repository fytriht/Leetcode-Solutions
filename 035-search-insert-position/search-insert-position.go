package searchInsert

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if target <= nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
