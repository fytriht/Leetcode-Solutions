// https://leetcode.com/problems/fizz-buzz/

/**
 * @param {string} s
 * @return {string}
 *
 * fizzBuzz(7) ==> [ '1', '2', 'Fizz', '4', 'Buzz', 'Fizz', '7' ]
 */
var fizzBuzz = function(n) {
  var arr = [], i = 0
  function trans(num) {
    if (num % 3 == 0 && num % 5 == 0) {
      return 'FizzBuzz'
    }
    return (num % 3 == 0) ? 'Fizz' : 'Buzz'
  }
  function isNeedTrans(num) {
    return num % 3 == 0 || num % 5 == 0
  }
  while(i < n) {
    arr[i] = isNeedTrans(i + 1) ? trans(i++ + 1) : String(i++ + 1)
    arr[--n] = isNeedTrans(n + 1) ? trans(n + 1) : String(n + 1)
  }
  return arr
};
