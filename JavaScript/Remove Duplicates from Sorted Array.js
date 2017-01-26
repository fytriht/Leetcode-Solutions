// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

/**
 * @param {number[]} nums
 * @return {number}
 */
 // [1,2,3,u,u,4,4,5]
var removeDuplicates = function(nums) {
  var i, j, len = nums.length
  for(i = 0; i < len; i++) {
    j = i + 1
    if (nums[i] == undefined) {
      continue
    }
    while (nums[i] == nums[j]) {
      nums[j] = undefined
      j++
    }
  }
  for(i = 0; i < len - 1; i++) {
    j = i
    if (nums[i] == undefined) {
      while (nums[++j] == undefined && j < len - 1);
      nums[i] = nums[j]
      nums[j] = undefined
    }
  }
  for(i = 0, j = 0; i < nums.length; i++) {
    if (nums[i] == undefined)
      j++
  }

  return nums.length - j
};

console.log(removeDuplicates([1,2,3,3,4,4,4,5]))
                          // [1,2,3,u,4,u,u,5]
