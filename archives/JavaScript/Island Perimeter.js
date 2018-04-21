// https://leetcode.com/problems/island-perimeter/

/**
 * @param {number[][]} grid
 * @return {number}
 */
var islandPerimeter = function(grid) {

  function getPerimeter(y, x, grid) {
    var counts = 0
    
    if (x == 0 || !grid[y][x-1]) counts++
    if (y == 0 || !grid[y-1][x]) counts++
    if (!grid[y+1] || !grid[y+1][x]) counts++
    if (!grid[y][x+1]) counts++
    
    return counts
  }

  var 
    yLen = grid.length, 
    xLen = grid[0].length, 
    counts = 0,
    i,
    j

  for (i = 0; i < yLen; i++) 
    for (j = 0; j < xLen; j++) 
      if (grid[i][j]) 
        counts += getPerimeter(i, j, grid)
    
  return counts
};

// console.log(islandPerimeter([[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]))

