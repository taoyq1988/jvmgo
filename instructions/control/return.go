package control

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IReturn struct {
	base.NoOperandsInstruction
}

func (_ *IReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.PopInt()
	invokeFrame.PushInt(val)
}

type LReturn struct {
	base.NoOperandsInstruction
}

func (_ *LReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.PopLong()
	invokeFrame.PushLong(val)
}

type FReturn struct {
	base.NoOperandsInstruction
}

func (_ *FReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.PopFloat()
	invokeFrame.PushFloat(val)
}

type DReturn struct {
	base.NoOperandsInstruction
}

func (_ *DReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.PopDouble()
	invokeFrame.PushDouble(val)
}

type AReturn struct {
	base.NoOperandsInstruction
}

func (_ *AReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.PopRef()
	invokeFrame.PushRef(val)
}

type Return struct {
	base.NoOperandsInstruction
}

func (_ *Return) Execute(frame *rtda.Frame) {
	frame.Thread.PopFrame()
}
