package solution

//
// solution 1
//

func findDisappearedNumbers(nums []int) []int {
	for _, n := range nums {
		if n = abs(n); nums[n-1] > 0 {
			nums[n-1] *= -1
		}
	}
	var ret []int
	for i, n := range nums {
		if n > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//
// solution 2
//

func findDisappearedNumbers2(nums []int) []int {
	m := map[int]int{}
	for i := 1; i <= len(nums); i++ {
		m[i]++
	}
	for _, n := range nums {
		m[n]--
	}

	ret := []int{}
	for k, v := range m {
		if v > 0 {
			ret = append(ret, k)
		}
	}
	return ret
}
