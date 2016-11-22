// https://leetcode.com/problems/intersection-of-two-arrays-ii/

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 *
 * Given nums1 = [1, 2, 2, 1] nums2 = [2, 2], return [2, 2]
 */
var intersect = function(nums1, nums2) {
  if (nums1.length == 0 || nums2.length == 0) {
    return []
  }
  if (nums1.length > nums2.length) {
    [nums1, nums2] = [nums2, nums1]
  }
  var hash = {}, result = [], i = 0
  while (i < nums1.length) {
    if (hash[nums1[i]] == undefined) {
      hash[nums1[i]] = 1
    } else {
      hash[nums1[i]]++
    }
    i++
  }
  i = 0
  while (i < nums2.length) {
    if (hash[nums2[i]]) {
      result.push(nums2[i])
      hash[nums2[i]]--
    }
    i++
  }
  return result
};

// console.log('', intersect([1, 2, 2, 1], [2, 2]))
