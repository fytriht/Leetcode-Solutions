// https://leetcode.com/problems/valid-anagram/

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 *
 * Given s = "anagram" t = "nagaram", return true
 */
var isAnagram = function(s, t) {
  if (s.length != t.length) {
    return false
  }
  var hash = {}, i, len = s.length
  for (i = 0; i < len; i++) {
    if (hash[s[i]] == undefined) {
      hash[s[i]] = 1
    } else {
      hash[s[i]]++
    }
  }
  for (i = 0; i < len; i++) {
    if (hash[t[i]] == undefined) {
      return false
    }
    hash[t[i]]--
    if (hash[t[i]] == -1) {
      return false
    }
  }
  return true
};

// console.log(isAnagram('anagram', 'nagaram'))