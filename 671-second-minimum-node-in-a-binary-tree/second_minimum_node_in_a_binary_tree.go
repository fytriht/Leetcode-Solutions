package solution

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	sec := math.MaxInt32

	for q := []*TreeNode{root}; len(q) > 0; {
		node := q[0]
		q = q[1:]
		if node.Val != root.Val {
			sec = min(sec, node.Val)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	if sec == root.Val || sec == math.MaxInt32 {
		return -1
	}
	return sec
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
