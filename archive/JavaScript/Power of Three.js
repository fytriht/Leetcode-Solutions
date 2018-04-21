// https://leetcode.com/problems/power-of-three/

/**
 * @param {number} n
 * @return {boolean}
 */
var isPowerOfThree = function(n) {
  while (n % 3 == 0 && n > 1) { 
    n /= 3
  }
  return n == 1
};

// 1
// 3
// 9
// 27
// 81
// 243
// 729
// 2187
// 6561
// 19683
// 59049