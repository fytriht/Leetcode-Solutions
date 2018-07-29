package solution

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += 2*numRows - 2 {
			ret += string(s[j])
			if k := 2*numRows - 2 + j - 2*i; k < len(s) && i != 0 && i != numRows-1 {
				ret += string(s[k])
			}
		}
	}
	return ret
}
