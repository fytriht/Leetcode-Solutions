// https://leetcode.com/problems/reverse-string/

/**
 * @param {string} s
 * @return {string}
 */
var reverseString = function(s) {
  var 
    result = [], 
    i, 
    j = s.length -1  

  for (i = 0; i <= j; i++, j--) {
    result[i] = s[j]
    result[j] = s[i]
  }
  
  return result.join('')
};

// console.log(reverseString(''))