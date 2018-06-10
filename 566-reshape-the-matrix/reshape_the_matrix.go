package solution

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	ret := [][]int{}
	row := []int{}
	for _, ns := range nums {
		for _, n := range ns {
			row = append(row, n)
			if len(row) == c {
				ret = append(ret, row)
				row = []int{}
			}
		}
	}
	return ret
}
