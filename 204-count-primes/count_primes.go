package solution

// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

func countPrimes(n int) int {
	m := make([]bool, n)
	for i := range m {
		m[i] = true
	}

	for i := 2; i*i < n; i++ {
		if m[i] {
			for j := i * i; j < n; j += i {
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
