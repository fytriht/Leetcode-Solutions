package solution

func letterCasePermutation(S string) []string {
	ret := []string{""}
	for _, b := range S {
		if b >= '0' && b <= '9' {
			for i := range ret {
				ret[i] += string(b)
			}
		} else {
			next := []string{}
			for _, s := range ret {
				next = append(next, s+string(b), s+string(b^32))
			}
			ret = next
		}
	}
	return ret
}
