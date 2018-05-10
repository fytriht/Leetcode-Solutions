package solution

import (
// "fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
// solution 1
//

func pathSum(root *TreeNode, sum int) int {
	m := map[int]int{0: 1}

	var rec func(root *TreeNode, target int, acc int) int
	rec = func(root *TreeNode, target int, acc int) int {
		if root == nil {
			return 0
		}
		acc += root.Val
		ret := m[acc-target]
		m[acc]++
		ret += rec(root.Left, target, acc) + rec(root.Right, target, acc)
		m[acc]--
		return ret
	}

	return rec(root, sum, 0)
}

//
// solution 2
//

func pathSum2(root *TreeNode, sum int) int {
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
