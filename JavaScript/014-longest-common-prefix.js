let longestCommonPrefix = strs => {
  let ret = ''
  if (strs.length === 0) {
    return ret
  }
  for (let i = 1, fst = strs.shift(); i <= fst.length; i++) {
    let tmp = fst.substring(0, i)
    for (let wd of strs) {
      if (wd.length < i || wd.substring(0, i) !== tmp) {
        return ret
      }
    }
    ret = tmp
  }
  return ret
}
