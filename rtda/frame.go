package rtda

import (
	. "github.com/taoyq1988/jvmgo/instructions/options"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"strings"
)

type OnPopAction func(popped *Frame)

type Frame struct {
	lower *Frame
	LocalVars
	OperandStack
	Thread       *Thread
	Method       *heap.Method
	maxLocals    uint
	maxStack     uint
	NextPC       int
	OnPopActions []OnPopAction
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
	return frame.Method.Class.ConstantPool
}

func (frame *Frame) RevertNextPC() {
	frame.NextPC = frame.Thread.PC
}

func (frame *Frame) AppendOnPopAction(action OnPopAction) {
	frame.OnPopActions = append(frame.OnPopActions, action)
}

func (frame *Frame) GetClass() *heap.Class {
	return frame.Method.Class
}

func (frame *Frame) GetClassLoader() *heap.ClassLoader {
	return heap.BootLoader() //todo
}

func newNativeFrame(thread *Thread, method *heap.Method) *Frame {
	frame := &Frame{
		Thread:       thread,
		Method:       method,
		LocalVars:    newLocalVars(method.ParamSlotCount),
		OperandStack: newOperandStack(4),
	}
	if method.Code == nil {
		method.Code = getHackCode(method.Descriptor)
	}
	return frame
}

var (
	_invokeNativeIReturn = []byte{OpInvokeNative, OpIReturn}
	_invokeNativeLReturn = []byte{OpInvokeNative, OpLReturn}
	_invokeNativeFReturn = []byte{OpInvokeNative, OpFReturn}
	_invokeNativeDReturn = []byte{OpInvokeNative, OpDReturn}
	_invokeNativeAReturn = []byte{OpInvokeNative, OpAReturn}
	_invokeNativeReturn  = []byte{OpInvokeNative, OpReturn}
)

func getHackCode(methodDescriptor string) []byte {
	rParenIdx := strings.IndexByte(methodDescriptor, ')')
	switch methodDescriptor[rParenIdx+1] {
	case 'V':
		return _invokeNativeReturn
	case 'L', '[':
		return _invokeNativeAReturn
	case 'D':
		return _invokeNativeDReturn
	case 'F':
		return _invokeNativeFReturn
	case 'J':
		return _invokeNativeLReturn
	default:
		return _invokeNativeIReturn
	}
}
