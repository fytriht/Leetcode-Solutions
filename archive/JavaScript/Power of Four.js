// https://leetcode.com/problems/power-of-four/

/**
 * @param {number} num
 * @return {boolean}
 */
var isPowerOfFour = function(num) {
  while (num % 4 == 0 && num > 1) { 
    num /= 4
  }
  return num == 1
};