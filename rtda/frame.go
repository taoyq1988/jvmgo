package rtda

import (
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

type Frame struct {
	lower *Frame
	LocalVars
	OperandStack
	Thread    *Thread
	Method    *heap.Method
	maxLocals uint
	maxStack  uint
	NextPC    int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		Thread:       thread,
		Method:       method,
		maxLocals:    method.MaxLocals,
		maxStack:     method.MaxStack,
		LocalVars:    newLocalVars(method.MaxLocals),
		OperandStack: newOperandStack(method.MaxStack),
	}
}

func (frame *Frame) Load(idx uint, isL bool) {
	s := frame.GetLocalVar(idx)
	frame.Push(s)
	if isL {
		frame.PushNull()
	}
}

func (frame *Frame) Store(idx uint, isL bool) {
	if isL {
		frame.Pop()
	}
	s := frame.Pop()
	frame.SetLocalVar(idx, s)
}

func (frame *Frame) GetConstantPool() heap.ConstantPool {
	return frame.method.Class.ConstantPool
}
