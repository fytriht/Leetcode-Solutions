// https://leetcode.com/problems/length-of-last-word/

/**
 * @param {string} s
 * @return {number}
 * 
 * Given s = 'hello world', return 5
 */
var lengthOfLastWord = function(s) {
  if (!s) return 0
  var i = s.length, counts = 0
  // find the last index of not blank(' ') value
  while (s[--i] === ' ');
  // count the length of last word
  while (i >= 0 && s[i--] !== ' ') counts++;
  return counts
};
