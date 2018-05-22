package solution

func islandPerimeter(grid [][]int) int {
	var ret int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i == 0 || grid[i-1][j] == 0 {
				ret++
			}
			if j == len(grid[i])-1 || grid[i][j+1] == 0 {
				ret++
			}
			if i == len(grid)-1 || grid[i+1][j] == 0 {
				ret++
			}
			if j == 0 || grid[i][j-1] == 0 {
				ret++
			}
		}
	}
	return ret
}
