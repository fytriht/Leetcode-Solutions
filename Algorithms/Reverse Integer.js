// https://leetcode.com/problems/reverse-integer

/**
 * @param {number} x
 * @return {number}
 *
 * Given x = -1234, return -4321
 */
var reverse = function(x) {
  if (x == 0)
    return 0

  var isNeg = x < 0, result = ''
  
  x = Math.abs(x)
  while (x) {
    result += (x % 10)
    x = Math.floor(x / 10)
  }
  result = parseInt(result)
  
  if (result > Math.pow(2, 31) - 1) 
    return 0

  return  isNeg ? -(result) : result
};