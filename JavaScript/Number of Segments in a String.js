// https://leetcode.com/problems/number-of-segments-in-a-string/

/**
 * @param {string} s
 * @return {number}
 */
var countSegments = function(s) {
  var i, counts = 0

  for (i = 0; i < s.length; i++) {
    if (s[i] !== ' ' && (s[i+1] === ' ' || s[i+1] === undefined)) {
      counts++
    }
  }
  return counts
};

// console.log(countSegments("    1     11   1111   1"))
// " 1   1"
