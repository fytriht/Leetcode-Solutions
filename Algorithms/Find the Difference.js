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
    diff = 0,
    i

  for (i = 0; i <= len / 2 - 1; i++) {
    diff += t[i].charCodeAt(0) + t[len-i-1].charCodeAt(0)
    diff -= s[i].charCodeAt(0) + s[len-i-1].charCodeAt(0)
  }

  if (i != len / 2) {
    diff += t[i].charCodeAt(0)
    diff -= s[i].charCodeAt(0)
  }

  return String.fromCharCode(diff + t[len].charCodeAt(0))
};
