package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	st := []*ListNode{}
	for n := head; n != nil; {
		st = append(st, n)
		n = n.Next
	}

	idx := len(st) - n - 1
	if idx < 0 {
		return head.Next
	}
	node := st[idx]
	node.Next = node.Next.Next
	return head
}
