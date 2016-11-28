// https://leetcode.com/problems/arranging-coins/

/**
 * @param {number} n
 * @return {number}
 *
 * n = 8, retrurn 3
 */
var arrangeCoins = function(n) {
  return Math.floor(-1/2 + Math.pow(1/4 + 2 * n, 1/2)) 
};

// console.log(arrangeCoins(8))

// 1/2 * i^2 + 1/2 * i - n

// (-1/2 +_ Math.pow((1/4 + 2n), 1/2))

// 1: -1/2 + 1/4 + 2n => 2n - 1/4
// 2: -1/2 - 1/4 - 2n => -2n - 3/4
