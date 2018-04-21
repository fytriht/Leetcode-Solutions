package levelOrderBottom

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextQ := []*TreeNode{}
		level := []int{}
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			if n == nil {
				continue
			}
			level = append(level, n.Val)
			nextQ = append(nextQ, n.Left, n.Right)
		}
		q = nextQ
		if len(level) > 0 {
			ret = append([][]int{level}, ret...)
		}
	}
	return ret
}
