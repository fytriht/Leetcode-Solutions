package levelOrderBottom

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	ret := [][]int{}
	for len(q) > 0 {
		nextQ := []*TreeNode{}
		level := []int{}
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			level = append(level, n.Val)
			if n.Left != nil {
				nextQ = append(nextQ, n.Left)
			}
			if n.Right != nil {
				nextQ = append(nextQ, n.Right)
			}
		}
		q = nextQ
		ret = append([][]int{level}, ret...)
	}
	return ret
}
