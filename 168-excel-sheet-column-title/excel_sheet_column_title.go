package convert

func convertToTitle(n int) string {
	ret := ""
	for n != 0 {
		ret = string((n-1)%26+'A') + ret
		n = (n - (n-1)%26) / 26
	}
	return ret
}
