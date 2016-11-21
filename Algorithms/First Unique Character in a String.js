// https://leetcode.com/problems/first-unique-character-in-a-string/

/**
 * @param {string} s
 * @return {number}
 *
 * Given s = 'leetcodle', return 0
 */
var firstUniqChar = function(s) {
  var cache = {}, i, len = s.length
  for (i = 0; i < len; i++) {
    if (cache[s[i]] == undefined) {
      cache[s[i]] = 0
    } else {
      cache[s[i]]++
    }
  }
  for (i = 0; i < len; i++) {
    if (cache[s[i]] === 0) {
      return i
    }
  }  
  return -1
};

// console.log(firstUniqChar('leetcolde'))