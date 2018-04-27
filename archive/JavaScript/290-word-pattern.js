let wordPattern = (pattern, str) => {
  str = str.split(' ')
  if (str.length != pattern.length) {
    return false
  }
  let m1 = {}
  let m2 = {}
  for (let i = 0; i < str.length; i++) {
    if (pattern[i] in m1) {
      if (m1[pattern[i]] != str[i]) {
        return false
      }
    } else {
      if (str[i] in m2) {
        return false
      }
      m2[str[i]] = pattern[i]
      m1[pattern[i]] = str[i]
    }
  }
  return true
}
