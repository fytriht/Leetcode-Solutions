// https://leetcode.com/problems/contains-duplicate/ 

/**
 * @param {number[]} nums
 * @return {boolean}
 * 
 * Givne nums = [0,4,5,0,3,6], return true
 */
var containsDuplicate = function(nums) {
  var hash = {}, i = 0, j = nums.length - 1
  while (i < j) {
    if (hash[nums[i]] != undefined) {
      return true
    }
    hash[nums[i]] = nums[i]
    if (hash[nums[j]] != undefined) {
      return true
    }
    hash[nums[j]] = nums[i]
    i++
    j--
  }
  if (i == j && hash[nums[i]] != undefined) {
    return true
  }
  return false
};

// console.log(containsDuplicate([0,4,5,0,3,6]))