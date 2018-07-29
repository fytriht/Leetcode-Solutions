package solution

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	ret := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			dp[j][i] = s[j] == s[i] && (i-j <= 1 || dp[j+1][i-1])
			if dp[j][i] && i-j+1 > len(ret) {
				ret = s[j : i+1]
			}
		}
	}
	return ret
}
