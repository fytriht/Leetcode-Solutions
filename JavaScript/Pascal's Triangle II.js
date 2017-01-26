// https://leetcode.com/problems/pascals-triangle-ii/

/**
 * @param {number} rowIndex
 * @return {number[]}
 */
var getRow = function(rowIndex) {
  var result = [1], i = 0, j
  
  return (function recur(arr) {
    var row = [1]
    for (j = 1; j < Math.round((arr.length + 1) / 2); j++) {
      row[j] = arr[j-1] + arr[j]
      row[arr.length - j] = row[j]
    }
    row.push(1)
    return i++ == rowIndex ? arr : recur(row)
  })(result)
};

/*
  TODO: tail call optimiztion
*/


console.log(getRow(0))

// [
//         [1 ],
//        [1 ,1 ],
//       [1 ,2 ,1 ],
//      [1 ,3 ,3 ,1 ],
//     [1 ,4 ,6 ,4 ,1 ],
//    [1 ,5 ,10,10,5 ,1 ]
// ]
