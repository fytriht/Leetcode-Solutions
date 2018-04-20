package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	s := []int{}
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	reversed := reverse(head)
	for head != nil {
		if head.Val != reversed.Val {
			return false
		}
		head, reversed = head.Next, reversed.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	i := head
	ret := &ListNode{Val: head.Val}
	for i.Next != nil {
		tmp := &ListNode{Val: i.Next.Val}
		tmp.Next = ret
		ret = tmp
		i = i.Next
	}
	return ret
}
