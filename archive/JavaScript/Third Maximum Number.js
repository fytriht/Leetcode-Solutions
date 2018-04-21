// https://leetcode.com/problems/third-maximum-number/

/**
 * @param {number[]} nums
 * @return {number}
 */
var thirdMax = function(nums) {
  var 
    max = Number.NEGATIVE_INFINITY, 
    secondMax = Number.NEGATIVE_INFINITY, 
    thirdMax = Number.NEGATIVE_INFINITY,
    i

  for (i = 0; i < nums.length; i++) {
    if (nums[i] >= max) {
      if (nums[i] != max) {
        thirdMax = secondMax
        secondMax = max
      }
      max = nums[i]
      continue
    } 
    if (nums[i] >= secondMax) {
      if (nums[i] != secondMax) {
        thirdMax = secondMax        
      }
      secondMax = nums[i]
      continue
    } 
    if (nums[i] >= thirdMax) {
      thirdMax = nums[i]
    }
  }
  return thirdMax == Number.NEGATIVE_INFINITY ? max : thirdMax
};

// console.log(thirdMax([1,1,2]))
