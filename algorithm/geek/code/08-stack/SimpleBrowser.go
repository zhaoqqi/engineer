package _8_stack

import "fmt"

type Browser struct {
	forwardStack  Stack
	backwardStack Stack
}

func NewBrowser() *Browser {
	return &Browser{
		NewLinkedListStack(),
		NewArrayStack(),
	}
}

func (this *Browser) CanForward() bool {
	if this.forwardStack.IsEmpty() {
		return false
	} else {
		return true
	}
}

func (this *Browser) CanBackward() bool {
	if this.backwardStack.IsEmpty() {
		return false
	} else {
		return true
	}
}

// Open指的的打开了新的页面，需要初始化浏览器的两个栈
func (this *Browser) Open(addr string) {
	fmt.Println("Open new address %+v\n", addr)
	this.forwardStack.Flush()
}

// Pushback指的是什么操作？
func (this *Browser) Pushback(addr string) {
	this.backwardStack.Flush()
}

func (this *Browser) Forward() {
	if this.forwardStack.IsEmpty() {
		return
	}
	top := this.forwardStack.Pop()
	this.backwardStack.Push(top)
	fmt.Println("forward to %+v", top)
}

func (this *Browser) Backward() {
	if this.backwardStack.IsEmpty() {
		return
	}
	top := this.backwardStack.Top()
	this.forwardStack.Push(top)
	fmt.Println("backward to %+v", top)
}
