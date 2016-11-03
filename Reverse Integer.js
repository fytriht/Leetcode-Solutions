// https://leetcode.com/problems/reverse-integer

/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
  if (x == 0)
    return 0

  const isNeg = x < 0
  let result = ''
  
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