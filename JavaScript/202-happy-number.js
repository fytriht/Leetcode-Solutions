let isHappy = n => {
  let m = {}
  while (true) {
    n = getNext(n)
    if (m[n]) return false
    if (n == 1) return true
    m[n] = true
  }
}

let getNext = n => {
  let s = 0
  for (let d of String(n)) s += d * d
  return s
}
