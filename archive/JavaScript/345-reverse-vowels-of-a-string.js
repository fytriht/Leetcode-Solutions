let isVowels = b =>
  ({
    a: true,
    e: true,
    i: true,
    o: true,
    u: true,
  }[b.toLowerCase()])

let reverseVowels = s => {
  let ret = s.split('')
  let i = 0
  let j = s.length - 1
  while (true) {
    while (i < j && !isVowels(s[i])) i++
    while (i < j && !isVowels(s[j])) j--
    if (i >= j) break
    ;[[ret[i], ret[j]]] = [[ret[j], ret[i]]]
    i++
    j--
  }
  return ret.join('')
}
