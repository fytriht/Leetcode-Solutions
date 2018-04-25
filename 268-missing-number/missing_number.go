package solution

func missingNumber(nums []int) int {
	ret := (len(nums) + len(nums)*len(nums)) / 2
	for _, n := range nums {
		ret -= n
	}
	return ret
}
