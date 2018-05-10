package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSumRec(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSumRec(root *TreeNode, sum int) int {
	var ret int
	sum -= root.Val
	if sum == 0 {
		ret++
	}
	if root.Left != nil {
		ret += pathSumRec(root.Left, sum)
	}
	if root.Right != nil {
		ret += pathSumRec(root.Right, sum)
	}
	return ret
}
