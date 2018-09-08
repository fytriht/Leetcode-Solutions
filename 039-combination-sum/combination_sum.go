package solution

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(candidates)
	dfs(candidates, &ret, []int{}, target, 0)
	return ret
}

func dfs(candidates []int, ret *[][]int, sol []int, target int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*ret = append(*ret, sol)
	} else {
		for i := start; i < len(candidates); i++ {
			sol = sol[:len(sol):len(sol)] // XXX
			dfs(candidates, ret, append(sol, candidates[i]), target-candidates[i], i)
		}
	}
}
