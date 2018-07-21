package solution

func shortestToChar(S string, C byte) []int {
	ret := make([]int, len(S))
	for i := range S {
		ret[i] = len(S)
	}

	prev := -len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			prev = i
		}
		ret[i] = min(ret[i], abs(prev-i))

	}
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == C {
			prev = i
		}
		ret[i] = min(ret[i], abs(prev-i))
	}

	return ret
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
