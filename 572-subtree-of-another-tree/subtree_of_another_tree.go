package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return hasSameStructure(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func hasSameStructure(x, y *TreeNode) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}
	return x.Val == y.Val && hasSameStructure(x.Left, y.Left) && hasSameStructure(x.Right, y.Right)
}
