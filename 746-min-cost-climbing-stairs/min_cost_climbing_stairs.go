package solution

//
// solution 1
//

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i < len(cost)+1; i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[len(dp)-1]
}

//
// solution 2
//

func minCostClimbingStairs2(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[len(dp)-1], dp[len(dp)-2])
}

//
// solution 2
//

func minCostClimbingStairs3(cost []int) int {
	var a, b int
	for _, c := range cost {
		a, b = b, c+min(a, b)
	}
	return min(a, b)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
