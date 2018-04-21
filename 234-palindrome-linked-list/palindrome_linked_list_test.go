package solution

import (
	"testing"
)

type testcase struct {
	in   []int
	want bool
}

func TestIsPalindrome(t *testing.T) {
	cases := []testcase{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2, 3, 3, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1, 2}, false},
		{[]int{1, 2, 3, 2}, false},
		{[]int{1, 2, 2, 1, 2, 2, 1}, true},
		{[]int{1, 1, 2, 1}, false},
	}
	for _, c := range cases {
		if got := isPalindrome(build(c.in)); got != c.want {
			t.Errorf(
				"isPalindrome(%v): expected %t but got %t",
				c.in,
				c.want,
				got,
			)
		}
	}
}

func build(s []int) *ListNode {
	h := &ListNode{}
	i := h
	for _, n := range s {
		i.Next = &ListNode{Val: n}
		i = i.Next
	}
	return h.Next
}

func parse(h *ListNode) []int {
	s := []int{}
	for h != nil {
		s = append(s, h.Val)
		h = h.Next
	}
	return s
}
