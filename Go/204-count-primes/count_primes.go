package count

func countPrimes(n int) int {
	m := make([]bool, n)
	for i := range m {
		m[i] = true
	}

	for i := 2; i < n; i++ {
		if m[i] {
			for j := 2 * i; j < n; j += i {
				m[j] = false
			}
		}
	}

	cnt := 0
	for i := 2; i < n; i++ {
		if m[i] {
			cnt++
		}
	}
	return cnt
}
