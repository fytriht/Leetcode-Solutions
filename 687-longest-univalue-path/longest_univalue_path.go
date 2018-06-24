package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(
		longestPath(root.Left, root)+longestPath(root.Right, root),
		max(
			longestUnivaluePath(root.Left),
			longestUnivaluePath(root.Right),
		),
	)
}

func longestPath(root, pre *TreeNode) int {
	if root == nil || root.Val != pre.Val {
		return 0
	}
	return 1 + max(longestPath(root.Left, root), longestPath(root.Right, root))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
