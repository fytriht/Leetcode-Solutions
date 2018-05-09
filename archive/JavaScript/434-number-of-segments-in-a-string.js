let countSegments = s => {
  let ret = 0
  for (let i = 1; i <= s.length; i++) {
    if ((i == s.length || s[i] == ' ') && s[i - 1] != ' ') {
      ret++
    }
  }
  return ret
}
