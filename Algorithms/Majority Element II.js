// https://leetcode.com/problems/majority-element-ii/

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var majorityElement = function(nums) {
  var result = [], hash = {}, i = 0, len = nums.length

  while (i < len) {
    if (hash[nums[i]] == undefined) {
      hash[nums[i]] = 1
    } else {
      hash[nums[i]]++
    }
    if (hash[nums[i]] > len / 3 && !result.includes(nums[i])) {
      result.push(nums[i])
    }
    if (result.length == 2) {
      break
    }
    i++
  }
  return result
};
