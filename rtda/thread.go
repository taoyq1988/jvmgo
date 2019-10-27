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
	JThread         *heap.Object // java.lang.Thread
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(defaultStackMaxSize),
	}
}

func (thread *Thread) IsStackEmpty() bool {
	return thread.stack.isEmpty()
}

func (thread *Thread) StackDepth() uint {
	return thread.stack.size
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
	if method.Name == "loadLibrary0" && method.Class.Name == "java/lang/Runtime" {
		fmt.Println("skip loadlibrary0")
		return
	}
	if method.Name == "loadLibrary0" && method.Class.Name == "java/lang/Runtime" {
		fmt.Println("skip loadlibrary0")
		return
	}

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
func (thread *Thread) throwException(className, initDesc string, initArgs ...heap.Slot) {
	class := heap.BootLoader().LoadClass(className)
	exObj := class.NewObj()
	athrowFrame := newAthrowFrame(thread, exObj, initArgs)
	thread.PushFrame(athrowFrame)
	// init exObj
	constructor := class.GetConstructor(initDesc)
	thread.InvokeMethod(constructor)
}

func (thread *Thread) throwExceptionV(className string) {
	thread.throwException(className, "()V")
}
func (thread *Thread) throwExceptionS(className, msg string) {
	msgObj := heap.JSFromGoStr(msg)
	thread.throwException(className, "(Ljava/lang/String;)V", heap.NewRefSlot(msgObj))
}

func (thread *Thread) HandleUncaughtException(ex *heap.Object) {
	thread.stack.clear()
	sysClass := heap.BootLoader().LoadClass("java/lang/System")
	sysErr := sysClass.GetStaticValue("out", "Ljava/io/PrintStream;").Ref
	printStackTrace := ex.Class.GetInstanceMethod("printStackTrace", "(Ljava/io/PrintStream;)V")

	newFrame := thread.NewFrame(printStackTrace)
	newFrame.SetRefVar(0, ex)
	newFrame.SetRefVar(1, sysErr)
	thread.PushFrame(newFrame)
}

func (thread *Thread) ThrowNPE() {
	thread.throwExceptionV("java/lang/NullPointerException")
}

func (thread *Thread) ThrowNegativeArraySizeException() {
	thread.throwExceptionV("java/lang/NegativeArraySizeException")
}

func (thread *Thread) ThrowClassCastException(from, to *heap.Class) {
	msg := fmt.Sprintf("%v cannot be cast to %v", from.NameJlsFormat(), to.NameJlsFormat())
	thread.throwExceptionS("java/lang/ClassCastException", msg)
}

func (thread *Thread) ThrowArrayIndexOutOfBoundsException(index int32) {
	msg := fmt.Sprintf("%v", index)
	thread.throwExceptionS("java/lang/ArrayIndexOutOfBoundsException", msg)
}

func (thread *Thread) ThrowDivByZero() {
	thread.throwExceptionS("java/lang/ArithmeticException", "/ by zero")
}
