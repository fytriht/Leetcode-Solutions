let isPalindrome = x => {
  if (x < 0) return false
  let bit = 1
  while (x / bit > 10) {
    bit *= 10
  }
  while (x != 0) {
    if (Math.floor(x / bit) !== x % 10) {
      return false
    }
    x = Math.floor((x % bit) / 10)
    bit /= 100
  }
  return true
}
