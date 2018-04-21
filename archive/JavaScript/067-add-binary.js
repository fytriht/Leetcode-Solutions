let addBinary = (a, b) => {
  let ret = ''
  let carry = 0
  let la = a.length
  let lb = b.length
  for (let i = 0; i < Math.max(la, lb); i++) {
    let va = (vb = 0)
    let tmp = carry
    if (i < la) va = Number(a[la - i - 1])
    if (i < lb) vb = Number(b[lb - i - 1])
    tmp += va + vb
    carry = Math.floor(tmp / 2)
    ret = String(tmp % 2) + ret
  }
  return carry == 0 ? ret : String(carry) + ret
}
