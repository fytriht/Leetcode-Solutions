// https://leetcode.com/problems/find-the-difference/

/**
 * @param {string} s
 * @param {string} t
 * @return {character}
 *
 * Given s='abcd' t='bacde', return 'e'
 */
var findTheDifference = function(s, t) {
  if (!s) return t
  var i, len = s.length, cache = {}
  for (i = 0; i < len; i++) {
    if (cache[s[i]] == undefined) 
      cache[s[i]] = 0
    if (cache[t[i]] == undefined)
      cache[t[i]] = 0
    cache[s[i]]++
    cache[t[i]]--
  }
  if (!cache[t[i]])
    return t[i]
  else
    cache[t[i]]--
  while (i--) {
    if (cache[t[i]]) 
      return t[i]
  }
};
