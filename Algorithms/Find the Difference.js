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
    j = s.length -1,
    i

  for (i = 0; i < j; i++, j--) {
    diff += t[i].charCodeAt(0)
    diff += t[j].charCodeAt(0)
    diff -= s[i].charCodeAt(0)
    diff -= s[j].charCodeAt(0)
  }
  if (i == j) {
    diff += t[i].charCodeAt(0)
    diff -= s[i].charCodeAt(0)
  }

  return String.fromCharCode(diff + t[len].charCodeAt(0))
};
