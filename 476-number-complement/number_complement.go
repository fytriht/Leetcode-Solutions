package solution

// TODO: gofmt
// TODO: test

func findComplement(num int) int {
	var match bool
	for i := 31; i >= 0; i-- {
		if (num & (1 << uint(i)) != 0) {
			match = true
		}
		if match {
            num ^= (1 << uint(i))
		}
	}
	return num
}