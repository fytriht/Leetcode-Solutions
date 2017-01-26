// https://leetcode.com/problems/move-zeroes/

/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 *
 * Given nums = [0, 1, 0, 3, 12], nums should be [1, 3, 12, 0, 0].
 */
var moveZeroes = function(nums) {
  var 
    i = 0, 
    j = 1, 
    len = nums.length
  while (j < len) {
    while (nums[i] != 0 && i < len - 1) {
      i++
    }
    while ((j <= i || nums[j] == 0) && j < len - 1) {
      j++
    }
    [nums[j], nums[i]] = [nums[i], nums[j]]
    j++
  }
  return nums
};
// [1,3,12,0,0]
// [1, 3, 12, 0, 0, 0] => [1, 3, 12, 0, 0]
// [0, 0]