package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	q1, q2 := []*TreeNode{p}, []*TreeNode{q}
	for len(q1) > 0 && len(q2) > 0 {
		n1, n2 := q1[len(q1)-1], q2[len(q2)-1]
		q1, q2 = q1[:len(q1)-1], q2[:len(q2)-1]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil || n1.Val != n2.Val {
			return false
		}
		q1 = append(q1, n1.Left, n1.Right)
		q2 = append(q2, n2.Left, n2.Right)
	}
	return true
}
