package solution

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

func deleteDuplicates2(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head = deleteDuplicates2(head.Next)
		} else {
			head.Next = deleteDuplicates2(head.Next)
		}
	}
	return head
}
