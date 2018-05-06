package solution

const (
	maxHexLen = 8
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var ret string
	for num != 0 && len(ret) < maxHexLen {
		h := num & 15
		if h >= 10 {
			ret = string('a'+h-10) + ret
		} else {
			ret = string('0'+h) + ret
		}
		num >>= 4
	}
	return ret
}
