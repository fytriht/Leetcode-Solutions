package solution

import (
	"testing"
)

func TestMyQueue(t *testing.T) {
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
	if s1.Peek() != 1 {
		t.Errorf("expected to get front element")
	}

	s2 := Constructor()
	s2.Push(1)
	s2.Pop()
	if !s2.Empty() {
		t.Errorf("expected pop works")
	}

	s3 := Constructor()
	s3.Push(1)
	s3.Push(2)
	s3.Push(3)
	if s3.Peek() != 1 {
		t.Error("expected to get front element")
	}

}
