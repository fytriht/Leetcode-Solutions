// https://leetcode.com/problems/number-of-1-bits/

/**
 * @param {number} n - a positive integer
 * @return {number}
 * 
 * Given n = 11, return 3
 */
var hammingWeight = function(n) {
  var counts = 0
  while (n) {
    if (n % 2 == 1) {
      counts++
    }
    n = (n - n % 2) / 2
  }    
  return counts
};

// console.log(hammingWeight(11))

// 11 % 2 = 5 R 1
// 5 % 2 = 2 R 1
// 2 % 2 = 1 R 0
// 1 % 2 = 0 R 1