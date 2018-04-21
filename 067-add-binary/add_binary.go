package solution

import "strconv"

func addBinary(a string, b string) string {
	ret := ""
	carry := 0
	la, lb := len(a), len(b)
	for i := 0; i < max(la, lb); i++ {
		va, vb := 0, 0
		if i < la {
			va, _ = strconv.Atoi(a[la-i-1 : la-i])
		}
		if i < lb {
			vb, _ = strconv.Atoi(b[lb-i-1 : lb-i])
		}
		tmp := carry + va + vb
		carry = tmp / 2
		ret = strconv.Itoa(tmp%2) + ret
	}
	if carry != 0 {
		ret = strconv.Itoa(carry) + ret
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
