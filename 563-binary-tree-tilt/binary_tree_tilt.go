package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
// solution 1
//

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

//
// solution 2
//

func findTilt2(root *TreeNode) int {
	var sum int
	if root == nil {
		return sum
	}

	var postorder func(*TreeNode) int
	postorder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := postorder(node.Left)
		right := postorder(node.Right)
		sum += abs(left - right)
		return left + right + node.Val
	}

	postorder(root)
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
