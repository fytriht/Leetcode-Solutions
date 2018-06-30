package solution

func findShortestSubArray(nums []int) int {
	m, minM, maxM := map[int]int{}, map[int]int{}, map[int]int{}
	for i, n := range nums {
		m[n]++
		if _, ok := minM[n]; !ok {
			minM[n] = i
		}
		maxM[n] = i
	}

	degree, elms := 0, []int{}
	for k, v := range m {
		if v > degree {
			degree = v
			elms = []int{k}
		} else if v == degree {
			elms = append(elms, k)
		}
	}

	ret := len(nums)
	for _, el := range elms {
		ret = min(ret, maxM[el]-minM[el]+1)
	}
	return ret
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
