// https://leetcode.com/problems/word-pattern/

/**
 * @param {string} pattern
 * @param {string} str
 * @return {boolean}
 *
 * Given pattern = 'abba', str = 'dog cat cat dog'
 */
var wordPattern = function(pattern, str) {
  var hash = {}, 
      hash1 = {}, 
      len = pattern.length, 
      arr = str.split(' '), 
      i
  if (len != arr.length) {
    return false
  }
  for (i = 0; i < len; i++) {
    if (hash[pattern[i]] == undefined) {
      if (hash1[arr[i]] in hash1) {
        return false
      }
      hash[pattern[i]] = arr[i]
      hash1[arr[i]] = arr[i]
    } else {
      if (hash[pattern[i]] != arr[i]) {
        return false
      } 
    }
  }
  return true
};

console.log(wordPattern('aaa', 'aa aa aa aa'))