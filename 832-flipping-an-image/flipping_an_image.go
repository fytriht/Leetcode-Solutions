package solution

func flipAndInvertImage(A [][]int) [][]int {
	for i, sl := range A {
		A[i] = flip(reverse(sl))
	}
	return A
}

func reverse(sl []int) []int {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return sl
}

func flip(sl []int) []int {
	for i, n := range sl {
		if n == 0 {
			sl[i] = 1
		} else {
			sl[i] = 0
		}
	}
	return sl
}
