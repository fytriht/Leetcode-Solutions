/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
  var i = 0, j = 0
// [3,5,6,7], 4, [2,5,8], 3
// [2,3,5,6,7] [2,5,8]
// [2,3,5,5,6,7] [2,5,8]
  while (i <= m && j <= n) {
    if (nums1[i] > nums2[j]) {
      move[nums1, i]
      nums1[i] = nums2[j]
      if (i <= m && j <= n) {
        i++
        j++
      }
      while (nums1[i] <= nums1[j] && i <= m && j <= n) {
        i++
      }
    }
    move(nums1, i)
    nums1[i] = nums2[j]
    if (i <= m && j <= n) {
      i++
      j++
    }
  } 

  function move(arr, i) {
    var j = arr.length -1
    while (j >= i) {
      arr[j+1] = arr[j]
      arr[j] = undefined 
      j--
    }
    return arr
  }
  return nums1
  // return
};
// [1,5,6,7], [2,5,8]

// [1,2,5,6,7] [5,8]
// [1,2,5,5,6,7] [8]
// [1,2,5,5,6,7,8] 

console.log(merge([3,5,6,7], 4, [2,5,8], 3))