package solution

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	ret := strconv.Itoa(t.Val)
	if t.Left != nil || t.Right != nil {
		ret += "(" + tree2str(t.Left) + ")"
	}
	if t.Right != nil {
		ret += "(" + tree2str(t.Right) + ")"
	}
	return ret
}
