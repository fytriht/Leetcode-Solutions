const reverse = x => {
  let ret = 0
  let sign = 1
  if (x < 0) {
    x = -x
    sign = -1
  }
  while (x !== 0) {
    ret = ret * 10 + x % 10
    x = Math.floor(x / 10)
  }
  return ret > Math.pow(2, 31) - 1 ? 0 : sign * ret
}
