package rtda

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

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

func (thread *Thread) IsStackEmpty() bool {
	return thread.stack.isEmpty()
}

/**
Operation of frame
*/
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

func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

func (thread *Thread) NewFrame(method *heap.Method) *Frame {
	if method.IsNative() {
		return newNativeFrame(thread, method)
	} else {
		return newFrame(thread, method)
	}
}

func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.top()
}

func (thread *Thread) TopFrameN(n uint) *Frame {
	return thread.stack.topN(n)
}

/**
InitClass
*/
func (thread *Thread) InitClass(class *heap.Class) {
	initClass(thread, class)
}

/**
Invoke
*/
func (thread *Thread) InvokeMethod(method *heap.Method) {
	fmt.Printf("-- goto invoke method %s, class %s, is native %v\n", method.Name, method.Class.Name, method.IsNative())
	currentFrame := thread.CurrentFrame()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	if n := method.ParamSlotCount; n > 0 {
		parseArgs(currentFrame, newFrame, n)
	}

	if method.IsSynchronized() {
		var monitor *heap.Monitor
		if method.IsStatic() {
			classObj := method.Class.JClass
			monitor = classObj.Monitor
		} else {
			thisObj := newFrame.GetThis()
			monitor = thisObj.Monitor
		}

		monitor.Enter(thread)
		newFrame.AppendOnPopAction(func(*Frame) {
			monitor.Exit(thread)
		})
	}
}

func (thread *Thread) InvokeMethodWithShim(method *heap.Method, args []heap.Slot) {
	shimFrame := newShimFrame(thread, args)
	thread.PushFrame(shimFrame)
	thread.InvokeMethod(method)
}

func parseArgs(from *Frame, to *Frame, argSlotsCount uint) {
	args := from.PopTops(argSlotsCount)
	for i := uint(0); i < argSlotsCount; i++ {
		to.SetLocalVar(i, args[i])
		args[i] = heap.EmptySlot
	}
}

/**
Throw
*/
func (thread *Thread) ThrowNEP() {
	fmt.Println("null point execption")
}

func (thread *Thread) ThrowClassCastException(from, to *heap.Class) {
	//msg := fmt.Sprintf("%v cannot be cast to %v", from.NameJlsFormat(), to.NameJlsFormat())
	//thread.throwExceptionS("java/lang/ClassCastException", msg)
}

func (thread *Thread) ThrowNegativeArraySizeException() {

}

func (thread *Thread) ThrowArrayIndexOutOfBoundsException(index int32) {

}

func (thread *Thread) ThrowDivByZero() {

}
