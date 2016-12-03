// https://leetcode.com/submissions/detail/84572166/

/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
 var merge = function(nums1, m, nums2, n) {
  var i = 0, j = 0
  nums1.length = m
  nums2.length = n

  // ([1,4,5,8],1) => [ 1, undefined, 4, 5, 8 ]
  function move(arr, i) {
    var j = arr.length - 1
    while (j >= i) {
      arr[j+1] = arr[j]
      j--
    }
  }
  if (m == 0) {
    while (n > 0) {
      nums1[n-1] = nums2[n-1]
      n--
    }
  } else {
    if (n != 0) {
      while (nums1.length != (m + n)) {
        while (nums1[i] < nums2[j] && nums1[i+1] != undefined) {
          i++
        } 
        if (nums1[i+1] == undefined) {
          if (nums1[i] < nums2[j]) {
            nums1[i+1] = nums2[j]
          } else {
            nums1[i+1] =nums1[i]
            nums1[i] = nums2[j]
          }
        } else {
          move(nums1, i)
          nums1[i] = nums2[j]
        }
        j++
      }
    } 
  }
};

