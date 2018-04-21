package solution

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.minStack) == 0 || x <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, x)
	}
}

func (s *MinStack) Pop() {
	if s.minStack[len(s.minStack)-1] == s.stack[len(s.stack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}
