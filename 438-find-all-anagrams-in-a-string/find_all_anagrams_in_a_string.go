package solution

func findAnagrams(s, p string) []int {
	var ret []int
	if len(s) == 0 {
		return ret
	}

	m := make(map[byte]int)
	for _, r := range p {
		m[byte(r)]++
	}

	left, right, cnt := 0, 0, len(p)
	for right < len(s) {
		if m[s[right]] > 0 {
			cnt--
		}
		m[s[right]]--
		right++
		if cnt == 0 {
			ret = append(ret, left)
		}
		if right-left == len(p) {
			if m[s[left]] >= 0 {
				cnt++
			}
			m[s[left]]++
			left++
		}
	}

	return ret
}
