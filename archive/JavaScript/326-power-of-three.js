let isPowerOfFour = num => {
  if (num == 1) return true
  if (num == 0 || num % 3 != 0) return false
  return isPowerOfFour(num / 3)
}
