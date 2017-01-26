// https://leetcode.com/problems/add-binary/

/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 *
 * Given a='1111' b='1111', return 11110
 */
var addBinary = function(a, b) {
  if (a.length < b.length)
    [a, b] = [b, a]
  var arr = [''], 
      i = a.length,
      diff = a.length - b.length
  while (--i >= 0) {
    if (!arr[i+1]) {
      arr[i+1] = +(a[i])
    }
    else {
      arr[i+1] += +(a[i])
    }
    if (i >= diff) {
      arr[i+1] += +(b[i-diff])
    }
    if (arr[i+1] >= 2) {
      arr[i+1] -= 2
      arr[i] = 1
    }
  }
  return arr.join('')
};

// '1111' '1111'
// [1,1,1,1,0]
// 