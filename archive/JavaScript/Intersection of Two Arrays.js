// https://leetcode.com/problems/intersection-of-two-arrays/

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 *
 * Given nums1 = [1, 2, 2, 1] nums2 = [2, 2], return [2]
 */
var intersection = function(nums1, nums2) {
  var i = 0,
      result = [],
      cache = {}
  if (nums1.length < nums2.length) {
    [nums1, nums2] = [nums2, nums1]
  }
  while (i < nums1.length) {
    cache[nums1[i]] = nums1[i]
    i++
  }
  i = 0
  while (i < nums2.length) {
    if (cache[nums2[i]] != undefined) {
      result.push(nums2[i])
      cache[nums2[i]] = undefined
    }
    i++
  }
  return result
};

console.log(intersection([1, 2, 2, 1], [2, 2]))

