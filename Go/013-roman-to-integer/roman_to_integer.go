package romanToInt

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ret, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		tmp := m[s[i]]
		if tmp < prev {
			ret = ret - tmp
		} else {
			ret = ret + tmp
		}
		prev = tmp
	}
	return ret
}
