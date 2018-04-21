package stack

import (
	"testing"
)

func TestMyStack(t *testing.T) {
	s := Constructor()
	if !s.Empty() {
		t.Errorf("expected to be empty")
	}
	s.Push(1)
	if s.Empty() {
		t.Errorf("expected push works")
	}

	s1 := Constructor()
	s1.Push(1)
	s1.Push(2)
	if s1.Top() != 2 {
		t.Errorf("expected to get top")
	}

	s2 := Constructor()
	s2.Push(1)
	s2.Pop()
	if !s2.Empty() {
		t.Errorf("expected pop works")
	}
}
