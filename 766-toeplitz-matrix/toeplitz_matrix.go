package solution

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		currR, nextR := matrix[i], matrix[i+1]
		if !isShallowEq(currR[:len(currR)-1], nextR[1:]) {
			return false
		}
	}
	return true
}

func isShallowEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}
