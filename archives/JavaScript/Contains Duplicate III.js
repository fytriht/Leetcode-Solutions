// https://leetcode.com/submissions/detail/84647013/

/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number} t
 * @return {boolean}
 *
 * Given nums=[2,1,3,4,5]  k=1 t=1, return true 
 */
var containsNearbyAlmostDuplicate = function(nums, k, t) {
  var i, j
 
  for (i = 0; i < nums.length; i++) {
    for (j = 0; j < nums.length; j++) {
      if (i != j && Math.abs(nums[i] - nums[j]) <= t && Math.abs(i -j) <= k) {
        return true
      }
    }
  }
  return false
};

// console.log(containsNearbyAlmostDuplicate([2,1,3,4,5], 1,1))