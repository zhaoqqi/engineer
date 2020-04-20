package _8_stack

type Stack interface {
	Push(v interface{})
	Pop()
	IsEmpty() bool
	Top() interface{}
	Flush()
}