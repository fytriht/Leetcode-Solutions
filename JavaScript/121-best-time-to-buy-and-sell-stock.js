let maxProfit = prices =>
  prices.reduce(
    ([max, begin], _, end) => (
      (pf = prices[end] - prices[begin]),
      pf <= 0 ? [max, end] : pf > max ? [pf, begin] : [max, begin]
    ),
    [0, 0],
  )[0]
