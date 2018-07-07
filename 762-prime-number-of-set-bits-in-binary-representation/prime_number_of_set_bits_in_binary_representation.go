package solution

//
// solution 1
//

var primes = map[int]bool{
	2:  true,
	3:  true,
	5:  true,
	7:  true,
	11: true,
	13: true,
	17: true,
	19: true,
}

func countPrimeSetBits(L int, R int) int {
	var cnt int
	for ; L <= R; L++ {
		if primes[countSetBits(L)] {
			cnt++
		}
	}
	return cnt
}

//
// solution 2
//

func countPrimeSetBits2(L int, R int) int {
	var cnt int
	for ; L <= R; L++ {
		if isPrime(countSetBits(L)) {
			cnt++
		}
	}
	return cnt
}

func countSetBits(n int) int {
	var cnt int
	for n != 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}

// https://en.wikipedia.org/wiki/AKS_primality_test
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
