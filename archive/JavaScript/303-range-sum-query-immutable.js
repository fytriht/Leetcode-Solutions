let NumArray = function(nums) {
  this.nums = [0, ...nums].reduce((acc, n) => {
    acc.push(n + (acc[acc.length - 1] || 0))
    return acc
  }, [])
}

NumArray.prototype.sumRange = function(i, j) {
  return this.nums[j + 1] - this.nums[i]
}
