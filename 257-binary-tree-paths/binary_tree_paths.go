package solution

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var ret []string

	var df func(*TreeNode, []string)
	df = func(root *TreeNode, path []string) {
		if root == nil {
			return
		}
		path = append(path, strconv.Itoa(root.Val))
		if root.Left == nil && root.Right == nil {
			ret = append(ret, strings.Join(path, "->"))
		}
		df(root.Left, path)
		df(root.Right, path)
	}

	df(root, []string{})
	return ret
}
