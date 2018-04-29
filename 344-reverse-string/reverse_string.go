package solution

func reverseString(s string) string {
	var ret []byte
	for i := len(s) - 1; i >= 0; i-- {
		ret = append(ret, s[i])
	}
	return string(ret)
}
