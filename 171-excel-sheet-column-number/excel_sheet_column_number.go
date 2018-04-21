package solution

func titleToNumber(s string) int {
	ret := 0
	for _, b := range s {
		ret = ret*26 + int(b-'A'+1)
	}
	return ret
}
