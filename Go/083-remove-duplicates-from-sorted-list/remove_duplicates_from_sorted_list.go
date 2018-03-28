package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	for curr := head; curr != nil; curr = curr.Next {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		}
	}
	return head
}
