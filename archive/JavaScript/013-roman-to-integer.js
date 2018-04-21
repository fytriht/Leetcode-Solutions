let map = {
  I: 1,
  V: 5,
  X: 10,
  L: 50,
  C: 100,
  D: 500,
  M: 1000,
}

let romanToInt = s => {
  let ret = prev = 0
  for (let i = s.length - 1; i >= 0; i--)
    ret = prev > (prev = map[s[i]]) ? ret - prev : ret + prev
  return ret
}
