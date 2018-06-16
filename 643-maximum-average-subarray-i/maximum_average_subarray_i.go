package solution

func findMaxAverage(nums []int, k int) float64 {
	var ret, sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
		ret += nums[i]
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
