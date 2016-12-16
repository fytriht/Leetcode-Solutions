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
    cond = len % 3 ? Math.ceil(len / 3) : len / 3 + 1,
    i

  for (i = 0; i < len; i++) {
    if (hash[nums[i]] == undefined) {
      hash[nums[i]] = 1
    } else {
      if (hash[nums[i]] == cond) {
        continue
      }
      hash[nums[i]]++
    }
    if (hash[nums[i]] == cond) {
      result.push(nums[i])
      if (result.length == 2) break
    }
  }
  return result
}

  // console.log(majorityElement([1]))