package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var pre *TreeNode
	var rec func(*TreeNode)
	rec = func(node *TreeNode) {
		if node == nil {
			return
		}
		rec(node.Right)
		if pre != nil {
			node.Val += pre.Val
		}
		pre = node
		rec(node.Left)
	}
	rec(root)
	return root
}
