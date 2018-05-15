let minMoves = nums => {
  let sum = 0
  let min = Number.POSITIVE_INFINITY
  nums.forEach(n => {
    if (n < min) min = n
    sum += n
  })
  return sum - nums.length * min
}
