// https://leetcode.com/submissions/detail/85515961/

/**
 * @param {string} version1
 * @param {string} version2
 * @return {number}
 */
var compareVersion = function(version1, version2) {
  version2 = version2.split('.')
  version1 = version1.split('.')

  var 
    i = 0, 
    len = Math.max(version1.length, version2.length), 
    m, 
    n

  while (i < len) {
    if (version1[i] == undefined) {
      while (version2[i] == 0) i++
      return version2[i] == undefined ? 0 : -1
    }
    if (version2[i] == undefined) {
      while (version1[i] == 0) i++ 
      return version1[i] == undefined ? 0 : 1
    }
    m = Number(version1[i])
    n = Number(version2[i])
    if (m > n) {
      return 1
    }
    if (m < n) {
      return -1
    }
    i++
  }
  return 0
};

console.log(compareVersion("01", "1" ))