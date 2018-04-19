package queue

type MyQueue struct {
	s *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		NewStack(),
	}
}

func (q *MyQueue) Push(x int) {
	tmp := []int{}
	for !q.s.IsEmpty() {
		tmp = append(tmp, q.s.Pop())
	}
	q.s.Push(x)
	for i := len(tmp) - 1; i >= 0; i-- {
		q.s.Push(tmp[i])
	}
}

func (q *MyQueue) Pop() int {
	return q.s.Pop()
}

func (q *MyQueue) Peek() int {
	return q.s.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.s.IsEmpty()
}
