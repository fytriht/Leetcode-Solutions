let climbStairs = n => {
  let [x, y] = [0, 1]
  for (let i = 0; i < n; i++) [x, y] = [y, x + y]
  return y
}
