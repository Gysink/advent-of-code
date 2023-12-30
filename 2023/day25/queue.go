package main

type Queue struct {
	queue []string
}

func NewQueue() *Queue {
	return &Queue{queue: make([]string, 0)}
}

func (q *Queue) push(item string) {
	q.queue = append(q.queue, item)
}

func (q *Queue) pop() string {
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *Queue) len() int {
	return len(q.queue)
}
