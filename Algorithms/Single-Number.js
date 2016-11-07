// https://leetcode.com/problems/single-number/


/**
 * @param {number[]} nums
 * @return {number}
 *
 * Given nums = [1,3,4,2,2,3,1], return 4
 */
var singleNumber = function(nums) {
  
  var theNum
  for (var num of nums) {
    theNum ^= num
  }
  return theNum
}
