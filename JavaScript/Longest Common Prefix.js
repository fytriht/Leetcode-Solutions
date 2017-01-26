// https://leetcode.com/problems/longest-common-prefix/

/**
 * @param {string[]} strs
 * @return {string}
 *
 * Given strs=['11','1111','122'], return '1'
 */
var longestCommonPrefix = function(strs) {
  if (!strs.length) {
    return ''
  }
  var prefix = strs[0], len
  // find the string with shortest length.
  for (var str of strs) {
    if (str.length < prefix.lenght) {
      prefix = str
    }
  }
  for (var str of strs) {
    prefix = findShortestPrefix(str, prefix)
    if (!prefix.length) {
      return ''
    }
  }
  function findShortestPrefix(str, prefix) {
    len = prefix.length  
    prefix = prefix.split('')
    while (len) {
      if (prefix[len-1] != str[len-1]) {
        prefix[len-1] = undefined // remove the last charater of `prefix`
        prefix.length--
      }
      len--
    }
    return prefix.join('')
  }
  return prefix
};
