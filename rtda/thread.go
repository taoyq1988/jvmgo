package rtda

import "github.com/taoyq1988/jvmgo/rtda/heap"

const (
	defaultStackMaxSize = 1024
)

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/
type Thread struct {
	PC    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(defaultStackMaxSize),
	}
}

func (thread *Thread) TopFrame() *Frame {
	return thread.stack.top()
}

func (thread *Thread) PopFrame() *Frame {
	top := thread.stack.pop()
	for _, action := range top.OnPopActions {
		action(top)
	}
	return top
}

func (thread *Thread) InitClass(class *heap.Class) {
	initClass(thread, class)
}

func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

func (thread *Thread) NewFrame(method *heap.Method) *Frame {
	if method.IsNative() {
		return nil //todo
	} else {
		return newFrame(thread, method)
	}
}
