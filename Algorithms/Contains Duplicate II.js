// https://leetcode.com/submissions/detail/84636616/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 *
 * Given nums = [1,0,1,1]  k = 1, return true 
 */
var containsNearbyDuplicate = function(nums, k) {
  var hash = {}, minDiff = Number.MAX_VALUE , i

  for (i = 0; i < nums.length; i++) {
    if (hash[nums[i]] == undefined) {
      hash[nums[i]] = i
    } else {
      if (i - hash[nums[i]] < minDiff) {
        minDiff = i - hash[nums[i]]
        hash[nums[i]] = i
      }
    }
  }
  return minDiff <= k
};

// console.log(containsNearbyDuplicate([1,0,1,1], 1))