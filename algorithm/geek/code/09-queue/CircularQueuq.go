package queue

import (
	"fmt"
)

// CircularQueue define a circular queue
type CircularQueue struct {
	q        []interface{}
	head     int
	tail     int
	capacity int
}

// NewCircularQueue create a new circular queue
func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), 0, 0, 0}
}

// IsEmpty check if the queue is empty
// 队空条件：head==tail为true
func (q *CircularQueue) IsEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

// IsFull check if the queue is full
// 队满条件：(tail+1)%capacity==head为true
func (q *CircularQueue) IsFull() bool {
	if q.head == (q.tail+1)%q.capacity {
		return true
	}
	return false
}

// EnQueue put a new one into the queue
func (q *CircularQueue) EnQueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.q[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

// DeQueue take over an element from the queue's head
func (q *CircularQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		return false
	}
	val := q.q[q.head]
	q.head = (q.head + 1) % q.capacity
	return val
}

// String make the queue a string
func (q *CircularQueue) String() string {
	if q.head == q.tail {
		return "empty queue"
	}
	result := "head<-"
	i := q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.q[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
