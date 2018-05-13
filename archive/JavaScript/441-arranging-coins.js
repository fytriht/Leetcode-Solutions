let arrangeCoins = n => {
  let left = 1,
    right = n
  while (left <= right) {
    let mid = left + Math.floor((right - left) / 2)
    if (2 * n < mid * (mid + 1)) {
      right = mid - 1
    } else {
      left = mid + 1
    }
  }
  return left - 1
}
