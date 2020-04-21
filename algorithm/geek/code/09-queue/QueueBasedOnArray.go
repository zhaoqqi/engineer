package queue

import (
	"fmt"
)

// ArrayQueue queue data stuct implement by array
type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

// NewArrayQueue return a new ArrayQueue object
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

// EnQueue put a new number to the end of the queue
func (array *ArrayQueue) EnQueue(v interface{}) bool {
	if array.tail == array.capacity {
		return false
	}
	array.q[array.tail] = v
	array.tail++
	return true
}

// DeQueue delete a number from the start of the queue
func (array *ArrayQueue) DeQueue() interface{} {
	if array.head == array.tail {
		return nil
	}
	v := array.q[array.head]
	array.head++
	return v
}

// String
func (array *ArrayQueue) String() string {
	if array.head == array.tail {
		return "empty queue"
	}
	result := "head"
	for i := array.head; i < array.tail; i++ {
		result += fmt.Sprintf("%+v", array.q[i])
	}
	result += "<-tail"
	return result
}
