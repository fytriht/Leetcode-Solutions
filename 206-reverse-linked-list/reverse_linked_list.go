package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	r := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return r
}

func reverseList2(head *ListNode) *ListNode {
	var h, tmp *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = h
		h = head
		head = tmp
	}
	return h
}
