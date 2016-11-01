// https://leetcode.com/problems/single-number/


/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
  let theNum

  for (let num of nums) {
    theNum ^= num
  }

  return theNum
}