// https://leetcode.com/problems/best-time-to-buy-and-sell-stock

/**
 * @param {number[]} prices
 * @return {number}
 *
 * Given prices = [7, 1, 5, 3, 6, 4], return 5
 */
var maxProfit = function(prices) {
  var max = 0, dayMax, i, j, len = prices.length
  function getMaxProfitByday(day) {
    dayMax = 0
    j = day
    while (++j < len) {
      if (prices[j] - prices[day] > dayMax) {
        dayMax = prices[j] - prices[day]
      }
    }
    return dayMax
  }
  for (i = 0; i < len; i++) {
    if (getMaxProfitByday(i) > max) {
      max = getMaxProfitByday(i)
    }
  }
  return max
};
// Obviously this is too SLOW

// console.log(maxProfit([7, 1, 5, 3, 6, 4]))