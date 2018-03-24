let removeDuplicates = nums => {
  let i = -1
  for (let n of nums) if (i == -1 || nums[i] != n) nums[++i] = n
  return i + 1
}
