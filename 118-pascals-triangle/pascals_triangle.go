package solution

func generate(numRows int) [][]int {
	var ret [][]int
	if numRows == 0 {
		return ret
	}
	ret = append(ret, []int{1})
	for i := 1; i < numRows; i++ {
		curr := []int{1}
		prev := ret[i-1]
		for i := 0; i < len(prev)-1; i++ {
			curr = append(curr, prev[i]+prev[i+1])
		}
		ret = append(ret, append(curr, 1))
	}
	return ret
}
