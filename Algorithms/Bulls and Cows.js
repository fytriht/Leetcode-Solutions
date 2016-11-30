// https://leetcode.com/problems/bulls-and-cows/

/**
 * @param {string} secret
 * @param {string} guess
 * @return {string}
 */
// var getHint = function(secret, guess) {
//   var hash = {}, bulls = 0, cows = 0
//   secret.split('').forEach((d, i) => {
//     if (d == guess[i]) {
//       bulls++
//     } else {
//       hash[d] ? hash[d]++ : (hash[d] = 1)
//     }
//   })
//   guess.split('').forEach((d, i) => {
//     if (hash[d] && d != secret[i]) {
//       cows++
//       hash[d]--
//     }
//   })
//   return `${bulls}A${cows}B`
// };

var getHint = function(secret, guess) {
  var hash = {}, bulls = 0, cows = 0, i = 0
  while (i < secret.length) {
    if (secret[i] == guess[i]) {
      bulls++
    } else {
      hash[secret[i]] || (hash[secret[i]] = 0)
      hash[secret[i]]++
    } 
    i++
  }
  while (i >= 0) {
    if (hash[guess[i]] && guess[i] != secret[i]) {
      cows++
      hash[guess[i]]--
    }
    i--
  }
  return bulls + 'A' + cows + 'B'
}

console.log(getHint('1807', '7810'))

// '1807', '7810' -> '1A3B'
// {
//   '1': 0,
//   '8': 1,
//   '0': 2,
//   '7': 3
// },
// {
//   '7': 0,
//   '8': 1,
//   '1': 2,
//   '0': 3
// }
// // '1123', '0111' -> '1A1B'
// {
//   '1': {
//     0: 0,
//     1: 1
//   },
//   '2': 2,
//   '3': 3
// },
// {
//   '0': 0,
//   '1': {
//     1: 1,
//     2: 2,
//     3: 3
//   }
// }


// const foo = {
//   bar : 1
// }
// delete foo.bar 
// console.log(Object.keys(foo).length)

