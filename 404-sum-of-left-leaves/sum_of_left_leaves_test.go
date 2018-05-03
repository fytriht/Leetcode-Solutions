package solution

import (
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
	tr := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}
	want := 24
	if got := sumOfLeftLeaves(tr); got != want {
		t.Errorf("expected: %d got: %d", want, got)
	}
}
