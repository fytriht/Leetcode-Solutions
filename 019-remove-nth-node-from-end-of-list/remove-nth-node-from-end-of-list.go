package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

//
// solution 1
//

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

//
// solution 2
//

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for ; n > 0; n-- {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
