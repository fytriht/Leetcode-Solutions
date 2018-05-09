let addStrings = (num1, num2) => {
  let ret = ''
  let i = num1.length - 1
  let j = num2.length - 1
  let carry = 0
  while (i >= 0 || j >= 0 || carry != 0) {
    if (i >= 0) {
      carry += Number(num1[i])
      i--
    }
    if (j >= 0) {
      carry += Number(num2[j])
      j--
    }
    ret = String(carry % 10) + ret
    carry = Math.floor(carry / 10)
  }
  return ret
}
