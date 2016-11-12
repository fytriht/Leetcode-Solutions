// https://leetcode.com/problems/remove-element/

/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 *
 * Given nums=[3,2,2,1] val=3, return 2
 */
var removeElement = function(nums, val) {
  var i = nums.length
  while (i--) {
    if (nums[i] == val) {
      nums[i] = nums[nums.length-1]
      nums.length--
    }
  }
  return nums.length
};

// [3,2,2,1]
