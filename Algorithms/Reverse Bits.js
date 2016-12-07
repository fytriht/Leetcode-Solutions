// https://leetcode.com/problems/reverse-bits/

/**
 * @param {number} n - a positive integer
 * @return {number} - a positive integer
 *
 * Given n = 43261596, return 964176192
 */
var reverseBits = function(n) {
  var i = 0, s = 0
  while (n) {
    if (n % 2 == 1) {
      s += Math.pow(2, 31 - i)
    }
    n = (n - n % 2) / 2
    i++
  }
  return s
};


// reverseBits(0)

// 11
// 00111001011110000010100101000000
// 5 R 1 1
// 2 R 1 2
// 1 R 0
// 0 R 1 8

// 1011