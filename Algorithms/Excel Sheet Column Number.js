// https://leetcode.com/problems/excel-sheet-column-number/

/**
 * @param {string} s
 * @return {number}
 *
 * Given s='AB', return 28
 */
var titleToNumber = function(s) {
  var i = s.length,
      result = 0,
      base = 'A'.charCodeAt(0) - 1,
      base1 = s.length - 1
  while (--i >= 0) {
    result += (s.charCodeAt(i) - base) * Math.pow(26, (base1-i))
  }
  return result
};

console.log(titleToNumber('ABCD'))