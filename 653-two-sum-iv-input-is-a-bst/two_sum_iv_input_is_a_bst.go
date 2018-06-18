package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := map[int]bool{}

	var preorder func(*TreeNode) bool
	preorder = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if m[k-node.Val] {
			return true
		}
		m[node.Val] = true
		return preorder(node.Left) || preorder(node.Right)
	}

	return preorder(root)
}
