package queue

type Queue []int

func (q *Queue) Push(v int) Queue {
	*q = append(*q, v)
	return *q
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
