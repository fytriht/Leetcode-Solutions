package solution

func letterCasePermutation(S string) []string {
	ret := []string{""}
	for _, b := range S {
		if b > '9' {
			next := []string{}
			for _, s := range ret {
				next = append(next, s+string(b), s+string(b^32))
			}
			ret = next
		} else {
			for i := range ret {
				ret[i] += string(b)
			}
		}
	}
	return ret
}
