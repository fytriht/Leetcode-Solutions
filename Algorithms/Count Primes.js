// https://leetcode.com/problems/count-primes/

/**
 * @param {number} n
 * @return {number}
 */
var countPrimes = function(n) {
  var i = 2, counts = 0
  
  while (i < n) {
    isPrime(i) && counts++
    i++
  }
  return counts
};

var isPrime = function(n) {
  if (n == 2) 
    return true
  
  var m = Math.ceil(Math.sqrt(n)), 
      i = 2

  while (i <= m) {
    if (n % i == 0) return false
    i++
  }
  return true
}

// console.log(countPrimes(999983))


// Suppose n is a whole number, and we want to test it to see if it is prime.   First, we take the square root (or the 1/2 power) of n; then we round this number up to the next highest whole number.  Call the result m.  We must find all of the following quotients:

// qm = n / m
// q(m-1) = n / (m-1)
// q(m-2) = n / (m-2)
// q(m-3) = n / (m-3)
// . . .
// q3 = n / 3
// q2 = n / 2

// The number n is prime if and only if none of the q's, as derived above, are whole numbers.