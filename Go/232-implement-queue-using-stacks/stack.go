package queue

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n int) {
	s.data = append(s.data, n)
}

func (s *Stack) Pop() int {
	n := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return n
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
