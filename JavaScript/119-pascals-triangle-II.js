let getRow = rowIndex => {
  let ret = [1]
  for (let i = 0; i < rowIndex; i++) {
    for (let j = i + 1; j > 0; j--) {
      ret[j] = (ret[j] || 0) + ret[j - 1]
    }
  }
  return ret
}
