package solution

var m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	ret := []string{}
	if len(digits) == 0 {
		return ret
	}

	ret = append(ret, "")
	for _, d := range digits {
		next := []string{}
		for _, b := range m[d-'2'] {
			for _, s := range ret {
				next = append(next, s+string(b))
			}
		}
		ret = next
	}
	return ret
}
