// https://leetcode.com/problems/ugly-number/

/**
 * @param {number} num
 * @return {boolean}
 * 
 * Given num = 14, return false
 */
var isUgly = function(num) {
  if (num < 1) {
    return false
  }
  while (num > 5) {
    if (num % 2 == 0 || num % 3 == 0 || num % 5 == 0) {
      while (num % 5 == 0) {
        num /= 5
      }
      while (num % 3 == 0) {
        num /= 3
      }
      while (num % 2 == 0) {
        num /= 2
      }
      continue
    }
    // if (num % 5 == 0) {
    //   num /= 5
    //   continue
    // }
    // if (num % 3 == 0) {
    //   num /= 3
    //   continue
    // }
    // if (num % 2 == 0) {
    //   num /= 2
    //   continue
    // }
    return false
  } 
  return true
};

console.log(isUgly(15))