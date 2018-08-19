package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	head = tmp.Next
	tmp.Next = swapPairs(tmp.Next.Next)
	head.Next = tmp
	return head
}
