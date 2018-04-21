package solution

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	fst := strs[0]
	for i := 1; i <= len(fst); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i || strs[j][:i] != fst[:i] {
				return fst[:i-1]
			}
		}
	}
	return fst
}
