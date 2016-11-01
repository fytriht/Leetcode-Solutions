/**
 * @param {string} s
 * @return {string}
 */
const fizzBuzz = function(n) {
  let arr = [], i = 0

  const trans = (num) =>
    num % 3 == 0 && num % 5 == 0
      ? 'fizzBuzz' : num % 3 == 0
      ? 'Fizz' : num % 5 == 0
      ? 'Buzz' : String(num) 
  
  while(i < n) {
    arr[i] = trans(i++ + 1)
    arr[--n] = trans(n + 1)
  }

  return arr
}
