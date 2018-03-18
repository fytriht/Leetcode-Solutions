package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	var ret string
	if len(strs) == 0 {
		return ret
	}
	for i, fst := 1, strs[0]; i <= len(fst); i++ {
		tmp := fst[:i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i || strs[j][:i] != tmp {
				return ret
			}
		}
		ret = tmp
	}
	return ret
}
