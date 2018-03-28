package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	for i := head; i != nil; i = i.Next {
		for i.Next != nil && i.Val == i.Next.Val {
			i.Next = i.Next.Next
		}
	}
	return head
}
