package solution

func findShortestSubArray(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
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
		i, j := 0, len(nums)-1
		for ; nums[i] != el; i++ {
		}
		for ; nums[j] != el; j-- {
		}
		ret = min(ret, j-i+1)
	}
	return ret
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
