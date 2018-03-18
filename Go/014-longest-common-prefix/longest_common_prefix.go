package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	var ret string
	if len(strs) == 0 {
		return ret
	}
	for i := 1; i <= len(strs[0]); i++ {
		tmp := strs[0][:i]
		matched := true
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i || strs[j][:i] != tmp {
				matched = false
				break
			}
		}
		if !matched {
			return ret
		}
		ret = tmp
	}
	return ret
}
