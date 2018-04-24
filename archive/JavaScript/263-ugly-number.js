let isUgly = num =>
  num > 0 &&
  [2, 3, 5].reduce((acc, n) => {
    while (acc % n == 0) acc /= n
    return acc
  }, num) == 1
