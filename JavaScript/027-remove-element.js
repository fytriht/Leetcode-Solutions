let removeElement = (nums, val) => {
  let i = 0
  let j = nums.length
  while (i < j) nums[i] == val ? (nums[i] = nums[--j]) : i++
  return j
}
