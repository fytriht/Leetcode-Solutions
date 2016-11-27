// // https://leetcode.com/problems/pascals-triangle/

// /**
//  * @param {number} numRows
//  * @return {number[][]}
//  */
var generate = function(numRows) {
  if (!numRows) {
    return []
  }
  var result = [[1]], i = 0
  function genRow(arr) {
    var row = [1]
    for (i = 0; i < arr.length - 1; i++) {
      row.push(arr[i] + arr[i+1])
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
  //   row = [1]
  //   for (i = 0; i < arr.length - 1; i++) {
  //     row.push(arr[i] + arr[i+1])
  //   }
  //   row.push(1)
  //   return row
  // }
  // console.log(genRow([1]))