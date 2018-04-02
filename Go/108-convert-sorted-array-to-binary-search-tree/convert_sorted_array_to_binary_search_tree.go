package sortedArrayToBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return fn(&nums, 0, len(nums)-1)
}

func fn(nums *[]int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	m := start + (end-start)/2
	return &TreeNode{
		Val:   (*nums)[m],
		Left:  fn(nums, start, m-1),
		Right: fn(nums, m+1, end),
	}
}
