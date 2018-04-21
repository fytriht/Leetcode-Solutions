package solution

import (
	"testing"
)

func TestInvertTree(t *testing.T) {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 2}
	root := &TreeNode{Left: l, Right: r}

	invertTree(root)
	if root.Left.Val != r.Val || root.Right.Val != l.Val {
		t.Error("expected the tree to be inverted")
	}
}
