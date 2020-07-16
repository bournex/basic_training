package queue

import "errors"

// FIFO队列实现

var (
	QUEUE_FULL  = errors.New("queue is full")
	QUEUE_EMPTY = errors.New("queue is empty")
)

type Queue interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	Length() int
}

func MakeQueue(maxlen int) Queue {
	q := new(queue)
	q.head = nil
	q.limit = maxlen
	q.tail = nil
	q.qlen = 0
	return q
}

type queueNode struct {
	next  *queueNode
	value interface{}
}

type queue struct {
	limit int
	qlen  int
	head  *queueNode
	tail  *queueNode
}

func (q *queue) Enqueue(v interface{}) error {
	if q.qlen == q.limit {
		return QUEUE_FULL
	}
	n := &queueNode{value: v}
	if q.tail != nil {
		q.tail.next = n
	}
	q.tail = n
	if q.head == nil {
		q.head = q.tail
	}
	q.qlen++
	return nil
}

func (q *queue) Dequeue() (interface{}, error) {
	var tmp *queueNode

	if q.head == nil && q.tail == nil {
		return nil, QUEUE_EMPTY
	}
	tmp = q.head
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.qlen--
	return tmp.value, nil
}

func (q *queue) Length() int {
	return q.qlen
}
