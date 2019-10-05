package rtda

import (
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

type Frame struct {
	lower *Frame
	LocalVars
	OperandStack
	thread *Thread
	method *heap.Method
	maxLocals uint
	maxStack uint
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:thread,
		method:method,
		maxLocals:method.MaxLocals,
		maxStack:method.MaxStack,
		LocalVars: newLocalVars(method.MaxLocals),
		OperandStack: newOperandStack(method.MaxStack),
	}
}

func (frame *Frame) GetConstantPool() heap.ConstantPool {
	return frame.method.Class.ConstantPool
}
