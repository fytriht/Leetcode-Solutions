package solution

type MyStack struct {
	q *Queue
}

func Constructor() MyStack {
	return MyStack{
		NewQueue(),
	}
}

func (s *MyStack) Push(x int) {
	s.q.Push(x)
	for i := s.q.Size() - 1; i > 0; i-- {
		s.q.Push(s.q.Pop())
	}
}

func (s *MyStack) Pop() int {
	return s.q.Pop()
}

func (s *MyStack) Top() int {
	return s.q.Peek()
}

func (s *MyStack) Empty() bool {
	return s.q.IsEmpty()
}

type MyStack2 struct {
	q   *Queue
	top int
}

func Constructor2() MyStack2 {
	return MyStack2{
		q: NewQueue(),
	}
}

func (s *MyStack2) Push(x int) {
	s.q.Push(x)
	s.top = x
}

func (s *MyStack2) Pop() int {
	for i := s.q.Size() - 1; i > 0; i-- {
		s.top = s.q.Pop()
		s.q.Push(s.top)
	}
	return s.q.Pop()
}

func (s *MyStack2) Top() int {
	return s.top
}

func (s *MyStack2) Empty() bool {
	return s.q.IsEmpty()
}
