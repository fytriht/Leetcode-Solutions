package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	s := []*TreeNode{root}
	for len(s) > 0 {
		n := s[len(s)-1]
		s = s[:len(s)-1]
		if n.Left == nil && n.Right == nil {
			if n.Val == sum {
				return true
			}
		}
		if n.Left != nil {
			n.Left.Val += n.Val
			s = append(s, n.Left)
		}
		if n.Right != nil {
			n.Right.Val += n.Val
			s = append(s, n.Right)
		}
	}
	return false
}
