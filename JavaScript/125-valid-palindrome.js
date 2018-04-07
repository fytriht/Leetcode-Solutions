let isAlnum = x =>
  (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') || (x >= '0' && x <= '9')

let isPalindrome = s => {
  let i = 0
  let j = s.length - 1
  while (i < j) {
    while (i < j && !isAlnum(s[i])) i++
    while (i < j && !isAlnum(s[j])) j--
    if (s[i].toLowerCase() !== s[j].toLowerCase()) return false
    i++
    j--
  }
  return true
}
