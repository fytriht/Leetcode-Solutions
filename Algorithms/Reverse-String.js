// https://leetcode.com/problems/reverse-string/

 /**
  * @param {string} s
  * @return {string}
  *
  * Given s = 'abcd', return 'dcba'
  */
var reverseString = function(s) {
    
  var newArr = [], i = 0, j = s.length - 1
  
  while (i <= j) {
    newArr[i] = s[j]
    newArr[j--] = s[i++]
  }
  return newArr.join('')
}