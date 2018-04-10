package majority

// http://www.cs.utexas.edu/~moore/best-ideas/mjrty/example.html

func majorityElement(nums []int) int {
	cdd, cnt := 0, 0
	for _, n := range nums {
		if cnt == 0 {
			cdd = n
		}
		if cdd == n {
			cnt++
		} else {
			cnt--
		}
	}
	return cdd
}
