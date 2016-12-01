// https://leetcode.com/problems/pascals-triangle-ii/

/**
 * @param {number} rowIndex
 * @return {number[]}
 */
var getRow = function(rowIndex) {
  var result = [1], i = 0, j, temp
  function recur(arr) {
    if (i++ == rowIndex) {
      return arr
    }
    var row = [1]
    for (j = 1; j < Math.round((arr.length + 1) / 2); j++) {
      temp = arr[j-1] + arr[j]
      row[j] = temp
      row[arr.length - j] = temp
    }
    row.push(1)
    return recur(row)
  }
  return recur(result)
};

/*
  TODO: tail call optimiztion
*/


console.log(getRow(3))

// [
//         [1 ],
//        [1 ,1 ],
//       [1 ,2 ,1 ],
//      [1 ,3 ,3 ,1 ],
//     [1 ,4 ,6 ,4 ,1 ],
//    [1 ,5 ,10,10,5 ,1 ]
// ]
