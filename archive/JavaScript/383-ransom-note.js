let canConstruct = (ransomNote, magazine) => {
  let m = {}
  for (let b of magazine) {
    if (!m.hasOwnProperty(b)) {
      m[b] = 0
    }
    m[b]++
  }
  for (let b of ransomNote) {
    if (!m.hasOwnProperty(b)) {
      return false
    }
    m[b]--
    if (m[b] < 0) {
      return false
    }
  }
  return true
}
