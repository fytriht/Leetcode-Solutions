/*
 * @constructor
 * @param {number[]} nums
 */
var NumArray = function(nums) {
  this.nums = nums
};

/**
 * @param {number} i
 * @param {number} j
 * @return {number}
 */
NumArray.prototype.sumRange = function(i, j) {
  var sum = 0
  while (i <= j) {
    sum += this.nums[i]
    i++
  }
  return sum
};

// var nums = [1,2,3,4,5]

// var numArray = new NumArray(nums);
// console.log(numArray.sumRange(0, 2))
// numArray.sumRange(0, 2);
