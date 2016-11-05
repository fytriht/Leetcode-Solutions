// https://leetcode.com/problems/roman-to-integer/


/**
 * @param {string} s
 * @return {number}
 * 
 * romanToInt('DCXXI') ==> 621
 */
var romanToInt = function(s) {

  // var hash = [['I', 1], ['V', 5], ['X', 10], ['L', 50], ['C', 100], ['D', 500], ['M', 1000]]
  // var map = new Map(hash)
  // 
  // Though intuitive, it seems use the above approach to defined a map is much slower.
  // 
 
  var map = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000
  }
  var len = s.length, result = 0, cur, next

  for (var i = 0; i < len; i++) {
    cur = map[s[i]]
    next = map[s[i + 1]]
    result += (cur < next) ?　-(cur) : cur 
  }

  return result
};

// 3256 => MMMCCLVI 3000 + 200 + 50 + 6
// 2234 => MMCCXXXIV 2000 + 200 + 30 + 4
// 2274 => MMCCLXXIV
// 2273 => MMCCLXXIII
// 1996 => MCMXCVI

// I V X   L   C   D   M
// 1 5 10  50  100 500 1000
