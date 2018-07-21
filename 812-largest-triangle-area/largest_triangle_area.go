package solution

func largestTriangleArea(points [][]int) float64 {
	var ret float64
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				ret = max(ret, 0.5*abs(float64(
					points[i][0]*points[j][1]+
						points[j][0]*points[k][1]+
						points[k][0]*points[i][1]-
						points[j][0]*points[i][1]-
						points[k][0]*points[j][1]-
						points[i][0]*points[k][1],
				)))
			}
		}
	}
	return ret
}

func max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
