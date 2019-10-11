package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type ISub struct {
	base.NoOperandsInstruction
}

func (_ *ISub) Execute(frame *rtda.Frame) {
	val1 := frame.PopInt()
	val2 := frame.PopInt()
	frame.PushInt(val2 - val1)
}

type LSub struct {
	base.NoOperandsInstruction
}

func (_ *LSub) Execute(frame *rtda.Frame) {
	val1 := frame.PopLong()
	val2 := frame.PopLong()
	frame.PushLong(val2 - val1)
}

type FSub struct {
	base.NoOperandsInstruction
}

func (_ *FSub) Execute(frame *rtda.Frame) {
	val1 := frame.PopFloat()
	val2 := frame.PopFloat()
	frame.PushFloat(val2 - val1)
}

type DSub struct {
	base.NoOperandsInstruction
}

func (_ *DSub) Execute(frame *rtda.Frame) {
	val1 := frame.PopDouble()
	val2 := frame.PopDouble()
	frame.PushDouble(val2 - val1)
}
