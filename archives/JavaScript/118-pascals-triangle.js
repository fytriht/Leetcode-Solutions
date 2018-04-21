let generate = numRows => {
  let ret = []
  if (numRows == 0) return ret
  ret.push([1])
  for (let i = 1; i < numRows; i++) {
    ret.push([1, ...ret[i - 1].map((v, j) => v + (ret[i - 1][j + 1] || 0))])
  }
  return ret
}
