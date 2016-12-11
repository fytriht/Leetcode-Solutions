// https://leetcode.com/problems/valid-palindrome/

// /**
//  * @param {string} s
//  * @return {boolean}
//  */ 
var isPalindrome = function(s) {
  var 
    i = 0, 
    j = s.length -1,
    r = /[A-Za-z0-9]/

  function isEqual(m, n) {
    // console.log(Math.abs(m.charCodeAt(0) - n.charCodeAt(0)))
    return Math.abs(m.charCodeAt(0) - n.charCodeAt(0)) == 32 
      && /[A-Za-z]/.test(m)
      && /[A-Za-z]/.test(n)
  }

  while (i < j) {
    while (!r.test(s[i])) {
      i++
    }
    while (!r.test(s[j])) {
      j--
    }
    if (i === j) {
      return true
    }
    if (s[i] != s[j] && !isEqual(s[i], s[j])) {
      return false
    }
    i++
    j--
  }
  return true
};

console.log(isPalindrome("aA"))
