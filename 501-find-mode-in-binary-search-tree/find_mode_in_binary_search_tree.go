package solution

// TODO: test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var ret []int
	var pre *TreeNode
	max, cnt := 0, 1

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if pre != nil {
			if node.Val == pre.Val {
				cnt++
			} else {
				cnt = 1
			}
		}
		if cnt >= max {
			if cnt > max {
				ret = ret[0:0]
				max = cnt
			}
			ret = append(ret, node.Val)
		}
		pre = node
		inorder(node.Right)
	}

	inorder(root)
	return ret
}
