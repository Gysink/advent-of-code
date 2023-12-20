package main

type QueueItem struct {
	origin Module
	target Module
	value  Pulse
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
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *Queue) len() int {
	return len(q.queue)
}

func NewQueueItem(orig Module, target Module, pulse Pulse) QueueItem {
	return QueueItem{
		origin: orig,
		target: target,
		value:  pulse,
	}
}
