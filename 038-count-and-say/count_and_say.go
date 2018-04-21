package solution

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return getNext(countAndSay(n - 1))
}

func getNext(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		count := 1
		for i < len(s)-1 && s[i] == s[i+1] {
			i++
			count++
		}
		ret += strconv.Itoa(count) + s[i:i+1]
	}
	return ret
}
