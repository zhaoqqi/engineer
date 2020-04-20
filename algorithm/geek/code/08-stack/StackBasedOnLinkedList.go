package _8_stack

import "fmt"

/*
基于链表实现到栈
*/

type node struct {
	next *node
	val interface{}
}

type LinkedListStack struct {
	// 栈顶节点
	topNode *node
}

func NewLinkedListStack() *LindedListStack {
	return &LinkedListStack(nil)
}

func (this *LinkedListStack) IsEmpty() bool {
	return nil == this.topNode 
}

func (this *LinkedListStack) push(v *node) {
	this.topNode = &node(next:this.topNode, val:v)
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	} else {
		popVal := this.topNode.val
		this.topNode = this.topNode.next
		return popVal
	}
}

func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	} else {
		return this.topNode.val
	}
}

func (this *LinkedListStack) Flush() {
	this.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for nil != this.topNode {
			fmt.Println(this.topNode.val)
		}
	}
}
