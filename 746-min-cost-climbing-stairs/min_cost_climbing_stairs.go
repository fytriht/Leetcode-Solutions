package solution

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i < len(cost)+1; i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[len(dp)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
