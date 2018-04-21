// https://leetcode.com/problems/add-strings/

/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 *
 * Given num1='584' num2='18', return '602'
 */
var addStrings = function(num1, num2) {
  if (num1.length < num2.length) {
    [num1, num2] = [num2, num1];
  }
  var diff = num1.length - num2.length, 
      sum = [], 
      i = num1.length,      // cureent index of sum[]
      j,                    // current index of num1[]
      k                     // current index of num2[]
  while (i > 0) {
    j = i - 1
    k = i - diff - 1
    if (!sum[i]) {
      sum[i] = 0
    }
    if (k >= 0) {
      sum[i] += +num1[j] + +num2[k]
    } else {
      sum[i] += +num1[j]
    }
    if (sum[i] >= 10) {
      sum[i - 1] = 1
      sum[i] -= 10
    }
    i--
  }
  return sum.join('')
};

// [ , 4, 7, 8, 4]
// console.log(addStrings('584', '18'))
// "9333852702227987"
//   "85731737104263"

