package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

//
// solution 1
//

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

//
// solution 2
//

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

//
// solution 3
//

func isPalindrome3(head *ListNode) bool {
	if head == nil {
		return true
	}
	mid := findMid(head)
	revMid := reverseInPlace(mid.Next)
	mid.Next = nil
	for revMid != nil {
		if head.Val != revMid.Val {
			return false
		}
		head, revMid = head.Next, revMid.Next
	}
	return true
}

// assumes non-nil head
func findMid(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseInPlace(head *ListNode) *ListNode {
	var ret *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = ret
		ret = head
		head = tmp
	}
	return ret
}

//
// solution 4
//

func isPalindrome4(head *ListNode) bool {
	left := head
	var rec func(*ListNode) bool
	rec = func(right *ListNode) bool {
		switch {
		case right == nil:
			return true
		case !rec(right.Next):
			return false
		case left.Val != right.Val:
			return false
		}
		left = left.Next
		return true
	}
	return rec(head)
}
