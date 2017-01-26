// https://leetcode.com/problems/happy-number/

/**
 * @param {number} n
 * @return {boolean}
 *
 *  Given n = 19, return true
 */
var isHappy = function(n) {
  var hash = {}, i, sum
  function getDigitsSquaresSum(num) { 
    // Given num = 18, reutrn sum = 65 (1^2 + 8^2). 
    // Given num = 62, return sum = 40 (6^2 + 2^2).
    sum = 0
    i = 10
    while (num) {
      sum += Math.pow(num % i / i * 10, 2)
      num -= num % i
      i *= 10
    } 
    return sum
  }
  while (hash[n] == undefined) {
    hash[n] = n
    n = getDigitsSquaresSum(n)
    if (n === 1) {
      return true
    }
  }
  return false
};

// console.log(isHappy(19))

// 18
// 1 + 64 = 65
// 36 + 25 = 61
// 36 + 1 = 37
// 9 + 48 = 57
// 25 + 49 = 74
// 49 + 16 = 65
// 36 + 25 = 61
