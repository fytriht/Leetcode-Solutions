package solution

var dirs = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func maxAreaOfIsland(grid [][]int) int {
	var ret int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 1 {
				continue
			}
			var cnt int
			q := [][]int{{i, j}}
			grid[i][j] = 0
			for len(q) > 0 {
				x, y := q[0][0], q[0][1]
				q = q[1:]
				cnt++
				for _, dir := range dirs {
					i, j := x+dir[0], y+dir[1]
					if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
						continue
					}
					q = append(q, []int{i, j})
					grid[i][j] = 0
				}
			}
			ret = max(ret, cnt)
		}
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
