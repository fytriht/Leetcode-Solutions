let isHappy = n => {
  let m = {}
  while (true) {
    n = getSumOfDigitsSuqare(n)
    if (m[n]) return false
    if (n == 1) return true
    m[n] = true
  }
}

let getSumOfDigitsSuqare = n => {
  let s = 0
  for (let d of String(n)) s += Math.pow(d, 2)
  return s
}
