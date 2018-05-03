package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	if n := root.Left; n != nil && n.Left == nil && n.Right == nil {
		return sum + n.Val
	}
	return sum
}
