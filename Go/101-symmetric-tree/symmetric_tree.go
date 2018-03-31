package isSymmetric

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
