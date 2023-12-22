package main

type Queue struct {
	queue []int
}

func NewQueue() *Queue {
	return &Queue{queue: make([]int, 0)}
}

func (q *Queue) push(item int) {
	q.queue = append(q.queue, item)
}

func (q *Queue) pop() int {
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *Queue) len() int {
	return len(q.queue)
}
