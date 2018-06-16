package solution

func findMaxAverage(nums []int, k int) float64 {
	var ret, sum int
	for _, n := range nums[:k] {
		sum += n
		ret += n
	}
	for i := 1; i < len(nums)-k+1; i++ {
		sum -= nums[i-1] - nums[i+k-1]
		ret = max(ret, sum)
	}
	return float64(ret) / float64(k)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
