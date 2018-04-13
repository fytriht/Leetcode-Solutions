let hammingWeight = n => {
  let cnt = 0
  while (n) {
    n &= n - 1
    cnt++
  }
  return cnt
}
