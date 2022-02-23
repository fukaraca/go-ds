package queues

import "sync"

type Queue struct {
	queue []interface{}
	lock  *sync.Mutex
	len   int
}

//creates a new queue structure
func New() *Queue {
	nev := Queue{}
	nev.lock = &sync.Mutex{}
	return &nev
}

//Enqueue inserts an element to the end of the queue
func (q *Queue) Enqueue(elem interface{}) {
	q.lock.Lock()
	q.queue = append(q.queue, elem)
	q.len++
	q.lock.Unlock()
}

//Dequeue removes an element from the start of the queue
func (q *Queue) Dequeue() {
	if q.len > 0 {
		q.lock.Lock()
		q.queue = q.queue[1:]
		q.len--
		q.lock.Unlock()
	}
}

//IsEmpty returns true if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

//Len returns length of the queue
func (q *Queue) Len() int {
	return q.len
}

//First returns the first element of the queue
func (q *Queue) First() interface{} {
	if q.len > 0 {
		return q.queue[0]
	}
	return nil
}

//Last returns the last element of the queue
func (q *Queue) Last() interface{} {
	if q.len > 0 {
		return q.queue[q.len-1]
	}
	return nil
}
