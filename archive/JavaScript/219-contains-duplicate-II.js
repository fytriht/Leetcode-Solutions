let containsNearbyDuplicate = (nums, k) => {
  let m = {}
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] in m && i - m[nums[i]] <= k) {
      return true
    }
    m[nums[i]] = i
  }
  return false
}
