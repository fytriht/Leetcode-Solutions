let majorityElement = nums =>
  nums.reduce(
    ([cdd, cnt], n) => [
      (cnt == 0 && (cdd = n), cdd),
      cdd == n ? cnt + 1 : cnt - 1,
    ],
    [0, 0],
  )[0]
