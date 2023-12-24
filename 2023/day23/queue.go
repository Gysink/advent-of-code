package main

type QueueItem struct {
	pos   Point
	steps int
	path  []Point
}

type Queue struct {
	queue []QueueItem
}

func NewQueue() *Queue {
	return &Queue{queue: make([]QueueItem, 0)}
}

func (q *Queue) push(item QueueItem) {
	q.queue = append(q.queue, item)
}

func (q *Queue) pop() QueueItem {
	item := q.queue[q.len()-1]
	q.queue = q.queue[:q.len()-1]
	return item
}

func (q *Queue) len() int {
	return len(q.queue)
}

func NewQueueItem(position Point, steps int, path []Point) QueueItem {
	return QueueItem{
		pos:   position,
		steps: steps,
		path:  path,
	}
}
