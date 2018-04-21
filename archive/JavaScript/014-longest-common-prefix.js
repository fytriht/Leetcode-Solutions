let longestCommonPrefix = strs => {
  if (strs.length === 0) {
    return ''
  }
  let fst = strs.shift()
  for (let i = 1 ; i <= fst.length; i++) {
    for (let wd of strs) {
      if (wd.length < i || wd.substring(0, i) !== fst.substring(0, i)) {
        return fst.substring(0, i -1 )
      }
    }
  }
  return fst
}
