// https://leetcode.com/problems/ransom-note/

/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 *
 * canConstruct("aa", "aab") -> true
 */
var canConstruct = function(ransomNote, magazine) {
  var 
    i = 0, 
    len = magazine.length, 
    cache = {}
  
  while (i < len) {
    if (!cache[magazine[i]]) {
      cache[magazine[i]] = 0
    }
    cache[magazine[i]]++
    i++
  }

  i = 0
  len = ransomNote.length
  
  while (i < len) {
    if (!cache[ransomNote[i]]) {
      return false
    }
    cache[ransomNote[i]]--
    if (cache[ransomNote[i]] < 0) {
      return false
    }
    i++
  }
 
  return true
};

// console.log(canConstruct('', 'a'))