package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRec(root.Left, root.Right)
}

func isSymmetricRec(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil || r2 == nil {
		return r1 == nil && r2 == nil
	}
	if r1.Val != r2.Val {
		return false
	}
	return isSymmetricRec(r1.Left, r2.Right) && isSymmetricRec(r1.Right, r2.Left)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	s := []*TreeNode{root.Left, root.Right}
	for len(s) > 0 {
		n1, n2, s := s[0], s[1], s[2:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil || n1.Val != n2.Val {
			return false
		}
		s = append(s, n1.Left, n2.Right, n1.Right, n2.Left)
	}
	return true
}
