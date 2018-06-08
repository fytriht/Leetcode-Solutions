package solution

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode
	ret := math.MaxInt32

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if pre != nil {
			ret = min(ret, node.Val-pre.Val)
		}
		pre = node
		inorder(node.Right)
	}

	inorder(root)
	return ret
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
