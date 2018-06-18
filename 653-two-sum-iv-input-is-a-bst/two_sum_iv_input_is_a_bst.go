package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := map[int]bool{}
	found := false

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if m[k-node.Val] {
			found = true
			return
		}
		m[node.Val] = true
		inorder(node.Right)
	}

	inorder(root)
	return found
}
