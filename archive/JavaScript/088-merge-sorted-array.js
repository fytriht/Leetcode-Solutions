let merge = (nums1, m, nums2, n) => {
  let i = m + n
  for (
    m--, n--, i--;
    m >= 0 && n >= 0;
    nums1[i--] = nums1[m] > nums2[n] ? nums1[m--] : nums2[n--]
  ) {}
  for (; n >= 0; nums1[i--] = nums2[n--]) {}
}
