package solution

func numberOfBoomerangs(points [][]int) int {
	var ret int
	for i := 0; i < len(points); i++ {
		m := map[int]int{}
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			m[dx*dx+dy*dy]++
		}
		for _, v := range m {
			if v > 1 {
				ret += v * (v - 1)
			}
		}
	}
	return ret
}
