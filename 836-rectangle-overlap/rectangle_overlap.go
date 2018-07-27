package solution

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return isIntersect(rec1[0], rec1[2], rec2[0], rec2[2]) && isIntersect(rec1[1], rec1[3], rec2[1], rec2[3])
}

func isIntersect(cl1, cr1, cl2, cr2 int) bool {
	return max(cl1, cl2) < min(cr1, cr2)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
