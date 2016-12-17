// https://leetcode.com/problems/number-of-segments-in-a-string/

/**
 * @param {string} s
 * @return {number}
 */
var countSegments = function(s) {
  var i = 0, counts = 0

  while (i < s.length) {
    if (s[i] !== ' ' && (s[i+1] === ' ' || s[i+1] == undefined)) {
      counts++
    }
    i++
  }
  return counts
};

// console.log(countSegments("    1     11   1111   1"))
// " 1   1"