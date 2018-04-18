package stack

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
