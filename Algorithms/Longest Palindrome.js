// https://leetcode.com/problems/longest-palindrome/

/**
 * @param {string} s
 * @return {number}
 *
 * Given s='abccccdd', return 7
 */
var longestPalindrome = function(s) {
  if (!s) {
    return 0
  }
  var i, hash = {}, len = s.length, sum = 0, addOne = false
  for (i = 0; i < len; i++) {
    if (hash[s[i]] == undefined) {
      hash[s[i]] = 1
    } else {
      hash[s[i]]++
    }
  }
  for (var key in hash) {
    if (hash[key] % 2 == 0) {
      sum += hash[key]
    } else {
      sum += hash[key] - 1
      addOne = true
    }
  }
  return addOne ? sum + 1 : sum
};

// console.log(longestPalindrome('aabbbbccd'))