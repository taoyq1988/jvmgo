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
	if top.OnPopAction != nil {
		top.OnPopAction()
	}
	return top
}

func (thread *Thread) InitClass(class *heap.Class) {
	//todo
}


//todo remove
func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}
//todo remove
func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrameTmp(self, maxLocals, maxStack)
}
