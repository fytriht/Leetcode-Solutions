/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
 
  let result = '', y = x
  
  while (y > 0) {
    result += (y % 10)
    y = Math.floor(y / 10)
  }

  return  y > 0 || result == x
};

/**
 *  if x = 0
 *  then y = 0, result = ''
 *  y > 0 ==> false, result = x ==> true
 *  return false
 *  
 *  if x < 0
 *  then y < 0
 *  y > 0 ==> true
 *  return true
 *
 *  I'm not sure if this is a good practice
 */