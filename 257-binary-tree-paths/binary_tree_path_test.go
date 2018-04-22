package solution

import (
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	// [1,2,3,null,5]
	tr := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	want := []string{"1->2->5", "1->3"}
	if got := binaryTreePaths(tr); !isSame(got, want) {
		t.Errorf(
			"binaryTreePaths(%v): expected %v but got %v",
			tr,
			want,
			got,
		)
	}
}

func isSame(sa, sb []string) bool {
	if len(sa) != len(sb) {
		return false
	}
	for i, s := range sa {
		if s != sb[i] {
			return false
		}
	}
	return true
}
