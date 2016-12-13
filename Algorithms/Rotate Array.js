// https://leetcode.com/submissions/detail/85415898/

var rotate = function(nums, k) {
  if (k != 0) {
    var arr = nums, i = nums.length - k
    arr = arr.slice(i).concat(arr.slice(0, i))
    arr.forEach((num, i) => nums[i] = num)
  }
};

// [1,2,3,4,5,6,7]
// [u,6,7,1,2,3,4]
// [5,6,7,1,2,3,4]
console.log(rotate([1,2], 1)) 