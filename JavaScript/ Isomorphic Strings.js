// https://leetcode.com/problems/isomorphic-strings/

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isIsomorphic = function(s, t) {
  var hash = {}, hash1 = {}, len = s.length, result = [], i
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
  return result.join('') == t
};

console.log(isIsomorphic('', ''))  // false
