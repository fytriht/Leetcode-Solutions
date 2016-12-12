// https://leetcode.com/submissions/detail/85415898/

var rotate = function(nums, k) {
  if (k != 0) {
    var arr = nums
    arr = arr.slice(nums.length - k).concat(arr.slice(0, nums.length - k))
    arr.forEach((num, i) => nums[i] = num)
  }
};

// [1,2,3,4,5,6,7]
// [5,6,7,1,2,3,4]
console.log(rotate([1,2], 1))