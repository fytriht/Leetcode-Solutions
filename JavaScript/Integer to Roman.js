// https://leetcode.com/problems/integer-to-roman/

/**
 * @param {number} num
 * @return {string}
 *
 * intToRoamn(2274) ==> MMCCLXXIV
 */

var intToRoman = function(num) {
  var hash = [1,5,10,50,100,500,1000]
  var vals = ['I','V','X','L','C','D','M']
  var arr = [], result = '', r
  // if num=2273, result in arr=[3,7,2,2]
  while (num) {
    r = num % 10
    arr.push(r)
    num -= r
    num /= 10
  }
  for (i = arr.length - 1; i >= 0; i--) {
    result += digitToRoman(i, arr[i]) 
  }
  function digitToRoman(i, num) {
    var result = '', d
    if (num == 4) {
      return digitToRoman(i, 1) + digitToRoman(i, 5)
    }
    if (num == 9) {
      return digitToRoman(i, 1) + digitToRoman(i+1, 1)
    }
    num *= Math.pow(10, i)
    d = 2 * (i + 1) - 1  
    if (hash[d] && num >= hash[d]) {
      result += vals[d]
      num -= hash[d]
    }
    d--
    while (num) {
      result += vals[d]
      num -= hash[d]
    }
    return result
  }
  return result
};

// 3256 => MMMCCLVI 3000 + 200 + 50 + 6
// 2234 => MMCCXXXIV 2000 + 200 + 30 + 4
// 2274 => MMCCLXXIV
// 2273 => MMCCLXXIII
// 1996 => M CM XC VI 1000 + (1000 - 100) + (100 - 10) + 5 + 1

// I V X   L   C   D   M
// 1 5 10  50  100 500 1000