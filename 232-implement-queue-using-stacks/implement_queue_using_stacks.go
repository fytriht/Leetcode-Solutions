package solution

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

type MyQueue2 struct {
	s1, s2 *Stack
}

func Constructor2() MyQueue2 {
	return MyQueue2{
		NewStack(),
		NewStack(),
	}
}

func (q *MyQueue2) Push(x int) {
	q.s1.Push(x)
}

func (q *MyQueue2) Pop() int {
	q.Peek()
	return q.s2.Pop()
}

func (q *MyQueue2) Peek() int {
	if q.s2.Size() == 0 {
		for !q.s1.IsEmpty() {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Peek()
}

func (q *MyQueue2) Empty() bool {
	return q.s1.IsEmpty() && q.s2.IsEmpty()
}
