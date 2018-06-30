package solution

func countBinarySubstrings(s string) int {
	var ret int
	for i := 0; i+1 < len(s); i++ {
		if s[i] != s[i+1] {
			for j, k := i, i+1; j >= 0 && k < len(s) && s[j] == s[i] && s[k] == s[i+1]; {
				ret++
				k++
				j--
			}

		}
	}
	return ret
}
