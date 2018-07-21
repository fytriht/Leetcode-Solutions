package solution

var maxWidth = 100

func numberOfLines(widths []int, S string) []int {
	var lines, lineWidth int
	for _, r := range S {
		if w := widths[r-'a']; lineWidth+w <= maxWidth {
			lineWidth += w
		} else {
			lines++
			lineWidth = w
		}
	}
	return []int{lines + 1, lineWidth}
}
