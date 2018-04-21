package stack

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(n int) {
	q.data = append(q.data, n)
}

func (q *Queue) Peek() int {
	return q.data[0]
}

func (q *Queue) Pop() int {
	fst := q.Peek()
	q.data = q.data[1:]
	return fst
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
