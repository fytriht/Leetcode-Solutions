let rotate = (nums, k) => {
  k %= nums.length
  reverse(nums, 0, nums.length - 1)
  reverse(nums, 0, k - 1)
  reverse(nums, k, nums.length - 1)
}

let reverse = (nums, start, end) => {
  while (start < end) [nums[start++], nums[end--]] = [nums[end], nums[start]]
}
