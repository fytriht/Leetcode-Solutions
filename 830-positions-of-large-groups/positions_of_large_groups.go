package solution

func largeGroupPositions(S string) [][]int {
	ret := [][]int{}
	for i := 0; i < len(S); {
		j := i + 1
		for j < len(S) && S[i] == S[j] {
			j++
		}
		if j-i >= 3 {
			ret = append(ret, []int{i, j - 1})
		}
		i = j
	}
	return ret
}
