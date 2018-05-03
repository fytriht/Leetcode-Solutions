package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
// solution 1
//

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	if n := root.Left; n != nil && n.Left == nil && n.Right == nil {
		return sum + n.Val
	}
	return sum
}

//
// solution 2
//

func sumOfLeftLeaves2(root *TreeNode) int {
	return sumOfLeftLeavesRec(root, false)
}

func sumOfLeftLeavesRec(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if isLeft {
			return root.Val
		}
		return 0
	}
	return sumOfLeftLeavesRec(root.Left, true) + sumOfLeftLeavesRec(root.Right, false)
}
