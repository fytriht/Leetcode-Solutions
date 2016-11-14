// https://leetcode.com/problems/climbing-stairs/

/**
 * @param {number} n
 * @return {number}
 *
 * Given n=5, return 8
 */
var climbStairs = function(n) {
  if (!n) return 0
  var a = 1, i = 0, counts = 1, temp

  while (++i < n) {
    temp = counts
    counts += a
    a = temp
  }
  return counts
};
// 1 1 =>1
// 2 11 2 => 2
// 3 111 12 21 =>3
// 4 1111 112 121 211 22 =>5
// 5 11111 1112 1121 1211 2111 122 212 221 =>8
// 6 111111 11112 11121 11211 12111 21111 1122 1212 2112 1221 2121 2211 222 =>13

//// With tail call optimization, the recursive version is still slower
// 'use strict'
// var climbStairs = function(n) {
//   function recur(n, a, b) {
//     if (n > 0) 
//       return recur(n - 1, b, a + b)
//     else 
//       return a
//   }
//   return recur(n, 1, 1)
// }