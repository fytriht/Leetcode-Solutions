package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(
		getHeight(root.Left)+getHeight(root.Right),
		max(
			diameterOfBinaryTree(root.Left),
			diameterOfBinaryTree(root.Right),
		),
	)
}

var m = make(map[*TreeNode]int)

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if v, ok := m[node]; ok {
		return v
	}
	m[node] = 1 + max(getHeight(node.Left), getHeight(node.Right))
	return m[node]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
