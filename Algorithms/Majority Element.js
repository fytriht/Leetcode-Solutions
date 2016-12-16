// https://leetcode.com/submissions/detail/85799911/

/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function(nums) {
  var hash = {} ,i = 0, len = nums.length

  while (i < len) {
    if (hash[nums[i]] == undefined) {
      hash[nums[i]] = 1
    } else {
      hash[nums[i]]++
    }
    if (hash[nums[i]] > len / 2) {
      return nums[i]
    }
    i++  
  }
};

// console.log(majorityElement([1,2,2,2]))