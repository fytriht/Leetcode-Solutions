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
	ret := []string{}
	binaryTreePathsRec(root, &ret, []string{})
	return ret
}

func binaryTreePathsRec(root *TreeNode, ret *[]string, path []string) {
	if root == nil {
		return
	}
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, strings.Join(path, "->"))
	}
	binaryTreePathsRec(root.Left, ret, path)
	binaryTreePathsRec(root.Right, ret, path)
}
