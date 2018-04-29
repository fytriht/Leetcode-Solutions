let isPowerOfFour = num => {
  for (let c = 1; c <= num; c *= 4) {
    if (c == num) {
      return true
    }
  }
  return false
}
