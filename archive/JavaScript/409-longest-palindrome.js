let longestPalindrome = s => {
  let m = {}
  for (let b of s) {
    if (!m.hasOwnProperty(b)) m[b] = 0
    m[b]++
  }
  let odd = 0
  for (let k in m) odd += m[k] & 1
  if (odd > 0) odd -= 1
  return s.length - odd
}
