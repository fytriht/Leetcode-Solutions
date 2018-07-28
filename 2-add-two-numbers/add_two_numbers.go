package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	head := ret
	carry := 0

	for l1 != nil || l2 != nil {
		s := carry
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{
			Val: s % 10,
		}
		head = head.Next
		carry = s / 10
	}

	if carry != 0 {
		head.Next = &ListNode{
			Val: carry,
		}
	}
	return ret.Next
}
