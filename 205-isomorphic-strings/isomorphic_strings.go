package solution

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s2t, t2s := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if r, ok := s2t[s[i]]; ok {
			if r != t[i] {
				return false
			}
		} else {
			if _, ok := t2s[t[i]]; ok {
				return false
			}
			s2t[s[i]] = t[i]
			t2s[t[i]] = s[i]
		}
	}
	return true
}

func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return isHalfIsmph(s, t) && isHalfIsmph(t, s)
}

func isHalfIsmph(s, t string) bool {
	m := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if r, ok := m[s[i]]; !ok {
			m[s[i]] = t[i]
		} else if r != t[i] {
			return false
		}
	}
	return true
}
