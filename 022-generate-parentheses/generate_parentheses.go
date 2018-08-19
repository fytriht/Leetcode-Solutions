package solution

func generateParenthesis(n int) []string {
	ret := []string{}

	var dfs func(int, int, string)
	dfs = func(left, right int, s string) {
		if left > right {
			return
		}
		if left == 0 && right == 0 {
			ret = append(ret, s)
		} else {
			if left > 0 {
				dfs(left-1, right, s+"(")
			}
			if right > 0 {
				dfs(left, right-1, s+")")
			}
		}
	}

	dfs(n, n, "")
	return ret
}
