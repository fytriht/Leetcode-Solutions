// https://leetcode.com/problems/plus-one/

/**
 * @param {number[]} digits
 * @return {number[]}
 *
 * Given nums = [9,9], return [1,0,0]
 */
var plusOne = function(digits) {
  var i = digits.length - 1
  digits[i]++
  while (i > 0) {
    if (digits[i] != 10) {
      return digits
    }
    digits[i] = 0
    digits[i-1]++      
    i--
  }
  if (digits[0] == 10) {
    digits[0] = 1
    digits[digits.length] = 0
  }
  return digits
};
