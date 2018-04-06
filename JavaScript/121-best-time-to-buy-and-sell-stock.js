let maxProfit = prices => {
  let begin = 0
  let end = 0
  let max = 0
  while (end < prices.length) {
    let profit = prices[end] - prices[begin]
    if (profit <= 0) {
      begin = end
    } else if (profit > max) {
      max = profit
    }
    end++
  }
  return max
}
