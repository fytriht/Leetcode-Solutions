package solution

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	ret[0] = 1
	for i := 0; i < rowIndex; i++ {
		for j := i + 1; j > 0; j-- {
			ret[j] += ret[j-1]
		}
	}
	return ret
}
