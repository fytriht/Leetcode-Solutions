package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(calcSum(root.Left)-calcSum(root.Right)) +
		findTilt(root.Left) +
		findTilt(root.Right)
}

func calcSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Val + calcSum(node.Left) + calcSum(node.Right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
