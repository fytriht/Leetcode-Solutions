let containsDuplicate = nums => {
  nums.sort()
  return nums.some((n, i) => n == nums[i + 1])
}
