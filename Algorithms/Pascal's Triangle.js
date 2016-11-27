// // https://leetcode.com/problems/pascals-triangle/

// /**
//  * @param {number} numRows
//  * @return {number[][]}
//  */
var generate = function(numRows) {
  if (!numRows) {
    return []
  }
  var result = [[1]], i = 0, j, temp

  function genRow(arr) {
    var row = [1]
    for (j = 1; j < Math.round((arr.length + 1) / 2) ; j++) {
      temp = arr[j-1] + arr[j]
      row[j] = temp
      row[arr.length - j] = temp
    }
    row.push(1)
    return row
  }
  while (i < numRows - 1) {
    result.push(genRow(result[i]))
    i++
  }
  return result
};
console.log(generate(6))

// // 6

// [
//         [1 ],
//        [1 ,1 ],
//       [1 ,2 ,1 ],
//      [1 ,3 ,3 ,1 ],
//     [1 ,4 ,6 ,4 ,1 ],
//    [1 ,5 ,10,10,5 ,1 ]
// ]



// function genRow(arr) {
//   var row = [1]
//   var temp
//   for (i = 1; i < Math.round((arr.length +1) / 2) ; i++) {
//     temp = arr[i-1] + arr[i]
//     row[i] = temp
//     row[arr.length - i] = temp
//   }
//   row.push(1)
//   return row
// }

// console.log(genRow([1,3,3,1]))