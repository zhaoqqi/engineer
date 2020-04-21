package queue

import (
	"fmt"
)

// ListNode node of linked list
type ListNode struct {
	val  interface{}
	next *ListNode
}

// LinkedListQueue based on linked list
type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

// NewLinkedlistQueue create a new LinkedListQueue object
func NewLinkedlistQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// EnQueue put a new number to the end of the queue
func (q *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{val: v, next: nil}
	if nil == q.tail {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

// DeQueue delete a number from the start of the queue
func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}
	node := q.head
	q.head = q.head.next
	q.length--
	return node.val
}

// String
func (q *LinkedListQueue) String() string {
	if q.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
