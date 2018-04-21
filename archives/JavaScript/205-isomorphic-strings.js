let isIsomorphic = (s, t) => {
  if (s.length != t.length) return false
  let s2t = {}
  let t2s = {}
  for (let i = 0; i < s.length; i++) {
    if (s[i] in s2t) {
      if (s2t[s[i]] !== t[i]) return false
    } else {
      if (t[i] in t2s) return false
      s2t[s[i]] = t[i]
      t2s[t[i]] = s[i]
    }
  }
  return true
}
