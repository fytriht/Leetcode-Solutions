// https://leetcode.com/problems/arranging-coins/

/**
 * @param {number} n
 * @return {number}
 *
 * n = 8, retrurn 3
 */
var arrangeCoins = function(n) {
  var i = 1
  while ((1 + i) * i / 2 <= n) {
    i++
  }
  return i - 1
};

console.log(arrangeCoins(30))