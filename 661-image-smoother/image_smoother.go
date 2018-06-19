package solution

func imageSmoother(M [][]int) [][]int {
	ret := [][]int{}
	for i, m := range M {
		row := []int{}
		for j := range m {
			row = append(row, calcGray(M, i, j))
		}
		ret = append(ret, row)
	}
	return ret
}

func calcGray(M [][]int, x int, y int) int {
	var sum, cnt int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			x, y := x+i, y+j
			if 0 <= x && x < len(M) && 0 <= y && y < len(M[0]) {
				sum += M[x][y]
				cnt++
			}
		}
	}
	return sum / cnt
}
