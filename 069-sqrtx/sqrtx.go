package solution

func mySqrt(x int) int {
	ret := x
	for ret*ret > x {
		ret = (ret + x/ret) / 2
	}
	return ret
}
