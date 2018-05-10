package solution

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	tr := &TreeNode{
		10,
		&TreeNode{
			5,
			&TreeNode{
				3,
				&TreeNode{
					3,
					nil,
					nil,
				},
				&TreeNode{
					-2,
					nil,
					nil,
				},
			},
			&TreeNode{
				2,
				nil,
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			-3,
			nil,
			&TreeNode{
				11,
				nil,
				nil,
			},
		},
	}

	ina, inb, want := tr, 8, 3
	if got := pathSum(ina, inb); got != want {
		t.Errorf(
			"\n input: %v, %d \n got: %d \n want: %d",
			ina,
			inb,
			got,
			want,
		)
	}
}
