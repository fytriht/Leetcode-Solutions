package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	s := []*TreeNode{root}
	for len(s) > 0 {
		node := s[0]
		s = s[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if node.Right != nil {
			s = append(s, node.Right)
		}
	}
	return root
}

func invertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	s := []*TreeNode{root}
	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if node.Right != nil {
			s = append(s, node.Right)
		}
	}
	return root
}
