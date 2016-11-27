// https://leetcode.com/problems/reverse-vowels-of-a-string/

/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function(s) {
  if (!s) {
    return s
  }
  var i = 0, 
      j = s.length - 1, 
      hash = { a:'',e:'',i:'',o:'',u:'',A:'',E:'',I:'',O:'',U:'' }
  
  s = s.split('')
  
  while (i < j) {
    while (!hash.hasOwnProperty(s[i]) && i < j) {
      i++
    }
    while (!hash.hasOwnProperty(s[j]) && i < j) {
      j--
    }
    [s[i], s[j]] = [s[j], s[i]]
    i++
    j--
  }
  return s.join('')
};
