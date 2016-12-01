// https://leetcode.com/problems/isomorphic-strings/

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isIsomorphic = function(s, t) {
  if (s.length == 0) {
    return true
  }
  var hash = {}, hash1 = {}, i, len = s.length, result = []
  for (i = 0; i < len; i++) {
    if (hash[s[i]] == undefined) {
      if (t[i] in hash1) {
        return false
      }
      hash[s[i]] = t[i]
      hash1[t[i]] = t[i]
    } else {
      if (hash[s[i]] != t[i]) {
        return false
      }
    }
    result[i] = t[i] 
  }
  // return hash1
  return result.join('') == t
};

// console.log(isIsomorphic('title', 'paper')) // true
console.log(isIsomorphic('ab', 'aa'))  // false
