// https://leetcode.com/submissions/detail/85242069/

/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
  if (numRows == 1)  return s

  var i = 0, 
      result = '',
      step = 2 * (numRows - 1), // 8
      len = s.length - 1,
      row,
      j

  while (i < numRows) {
    row = ''
    j = 0

    while ((-i + step * (j + 1)) <= len || (i + step * j) <= len) {

      if ((i + step * j) <= len) {
        // console.log((i + step * j), '!!')
        row += s[i + step * j] // 0 
      }
      if ((i % (numRows - 1) != 0) && (-i + step * (j + 1)) <= len) {
        row += s[-i + step * (j + 1)]
      }
      j++
    }
    result += row
    i++
  }
  return result
};

// console.log(convert('PAYPALISHIRING', 2))

// 0P 1A 2Y 3P 4A 5L 6I 7S 8H 9I 10R 11I 12N 13G

// 1
// PAYPALISHIRING

// 2
// 0P 2Y 4A 6I 8H 10R 12N
// 1A 3P 5L 7S 9I 11I 13G
// PYAIHRNAPLSIIG

// 3
// 0P    4A    8H     12N
// 1A 3P 5L 7S 9I 11I 13G
// 2Y    6I    10R
// PAHNAPLSIIGYIR



        
// 4
// 0P               6I                   12N 
//   1A          5L    7S            11I     13G  
//      2Y    4A          8H     10R
//         3P                 9I
// PIN ALSIG YAHR PI

// 5
// 0P                     8H
//   1A                7S    9I
//      2Y          6I          10R
//         3P    5L                 11I     13G
//            4A                        12N
// PH ASI YIR PLIG AN

// 6
// 0P                            10R        
//    1A                      9I     11I        
//       2Y                8H            12N    
//          3P          7S                   13G
//             4A    6I                          
//                5L                                 
// PR AII YHN PSG AI L

