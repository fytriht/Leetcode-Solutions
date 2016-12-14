// https://leetcode.com/problems/island-perimeter/

/**
 * @param {number[][]} grid
 * @return {number}
 */
var islandPerimeter = function(grid) {

  function getPerimeter(y, x, grid) {
    var counts = 0
    if (x == 0 || !grid[y][x-1]) counts++
    if (!grid[y][x+1]) counts++
    if (y == 0 || !grid[y-1][x]) counts++
    if (!grid[y+1] || !grid[y+1][x]) counts++
    
    return counts
  }

  var i, j, yLen = grid.length, jLen = grid[0].length, counts = 0

  for (i = 0; i < yLen; i++) {
    for (j = 0; j < jLen; j++) {
      if (grid[i][j]) counts += getPerimeter(i, j, grid)
    }
  }
  return counts
};

// islandPerimeter([[0,1]])

