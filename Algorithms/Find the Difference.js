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
    hash = {}, 
    result,
    i

  for (i = 0; i < len; i++) {
    if (hash[t[i]] == undefined) {
      hash[t[i]] = 1
    } else {
      hash[t[i]]++
    }
   
    if (hash[s[i]] == undefined) {
      hash[s[i]] = -1
    } else {
      hash[s[i]]--
    }
  
  }
  
  if (!hash[t[i]]) {
    return t[i]
  } else {
    hash[t[i]]++
  }
  
  for (var item in hash) {
    if (hash[item] == 1) {
      return item
    }
  }
};
