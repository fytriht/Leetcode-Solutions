let thirdMax = nums => {
  let i = (j = k = Number.NEGATIVE_INFINITY)
  nums.forEach(n => {
    if (n > i) {
      ;[k, j, i] = [j, i, n]
    } else if (n == i) {
      return
    } else if (n > j) {
      ;[k, j] = [j, n]
    } else if (n == j) {
      return
    } else if (n > k) {
      k = n
    }
  })
  return k > Number.NEGATIVE_INFINITY ? k : i
}
