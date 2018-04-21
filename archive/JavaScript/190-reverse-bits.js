let reverseBits = n => {
  let ret = 0
  for (let i = 0; i < 32; i++) {
    ret = ret * 2 + (n & 0x1)
    n /= 2
  }
  return ret
}
