package reverse

import (
	"testing"
)

type testcase struct {
	in   []int
	want []int
}

func TestReverseList(t *testing.T) {
	cases := []testcase{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
	}
	for _, c := range cases {
		if got := parseList(reverseList(genList(c.in))); !isShallowEq(got, c.want) {
			t.Errorf(
				"reverseList(%v): expected %v but got %v",
				c.in,
				c.want,
				got,
			)
		}
	}
}

func genList(s []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, n := range s {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return head.Next
}

func parseList(l *ListNode) []int {
	s := []int{}
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return s
}

func isShallowEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}
