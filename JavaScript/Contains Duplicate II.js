// https://leetcode.com/submissions/detail/84636616/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 *
 * Given nums = [1,0,1,1]  k = 1, return true 
 */
var containsNearbyDuplicate = function(nums, k) {
  var hash = {}, i

  for (i = 0; i < nums.length; i++) {
    if (i - hash[nums[i]] <= k) { // i - undefined <= k  -->  false
      return true
    }
    hash[nums[i]] = i
  }
  return false
};

console.log(containsNearbyDuplicate([1,0,1,1], 1))