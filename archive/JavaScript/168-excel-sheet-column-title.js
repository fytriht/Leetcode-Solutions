let convertToTitle = n => {
  let ret = ''
  while (n != 0) {
    ret = String.fromCharCode((n - 1) % 26 + 'A'.charCodeAt()) + ret
    n = Math.floor((n - (n - 1) % 26) / 26)
  }
  return ret
}
