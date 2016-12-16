// https://leetcode.com/problems/majority-element-ii/

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var majorityElement = function(nums) {
  var 
    result = [], 
    hash = {}, 
    len = nums.length,
    cond = len % 3 ? Math.ceil(len / 3) : len / 3 + 1
    i = 0

  while (i < len) {
    if (hash[nums[i]] == undefined) {
      hash[nums[i]] = 1
    } else {
      hash[nums[i]]++
    }
    if (hash[nums[i]] == cond) {
      result.push(nums[i])
      if (result.length == 2) break
    }
    i++
  }
  return result
};

console.log(majorityElement([1]))