package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	ret := []float64{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum, cnt := 0, len(q)
		for i := 0; i < cnt; i++ {
			node := q[0]
			sum += node.Val
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ret = append(ret, float64(sum)/float64(cnt))
	}
	return ret
}
