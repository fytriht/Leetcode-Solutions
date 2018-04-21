let maxProfit = prices =>
  prices.reduce(
    ([max, begin], _, end) => (
      (pf = prices[end] - prices[begin]),
      pf <= 0 ? (begin = end) : pf > max && (max = pf),
      [max, begin]
    ),
    [0, 0],
  )[0]
