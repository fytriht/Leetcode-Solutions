// https://leetcode.com/problems/find-the-difference/

/**
 * @param {string} s
 * @param {string} t
 * @return {character}
 *
 * Given s='abcd' t='bacde', return 'e'
 */
var findTheDifference = function(s, t) {
  var 
    len = s.length, 
    diff = t.charCodeAt(len),
    i

  for (i = 0; i <= len / 2 - 1; i++) {
    diff += t.charCodeAt(i) + t.charCodeAt(len-i-1) 
      - s.charCodeAt(i) - s.charCodeAt(len-i-1)
  }

  if (i != len / 2) {
    diff += t.charCodeAt(i) - s.charCodeAt(i)
  }

  return String.fromCharCode(diff)
};